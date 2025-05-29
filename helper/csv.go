package helper

import (
	"encoding/csv"
	"errors"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"reflect"
	"sync"
)

const (
	// csv的标题
	CsvHeaderTag = "header"

	// 分隔符
	CsvFormatTag = "format"

	// 默认日期时间格式
	DefaultDateTimeFormat = "2006-01-02 15:04:05"
)

// csv的列结构
type csvColumn struct {
	Header      string
	ColumnIndex int
	FieldIndex  int
	Format      string
}

// csv的配置
type Csv[T any] struct {
	// 是否有标题
	hasHeader bool

	// csv的列
	columns []csvColumn

	// 日志实例
	Logger *slog.Logger
}

// csv的构造函数
// 返回结果是指针
func NewCsv[T any](hasHeader bool) (*Csv[T], error) {
	c := &Csv[T]{
		hasHeader: hasHeader,
		Logger:    slog.Default(),
	}

	// csv的行一定要是结构体
	structType := reflect.TypeOf((*T)(nil)).Elem()
	if structType.Kind() != reflect.Struct {
		return nil, errors.New("type not a struct")
	}

	// 创建一个映射，将csv的列与响应结构体字段链接
	c.columns = make([]csvColumn, structType.NumField())
	// 遍历字段
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)

		// 设置列标题
		// 获取结构体字段的标签，取不到则使用字段名
		header, ok := field.Tag.Lookup(CsvHeaderTag)
		if !ok {
			header = field.Name
		}

		// 获取format标签，取不到使用默认
		format, ok := field.Tag.Lookup(CsvFormatTag)
		if !ok {
			format = DefaultDateTimeFormat
		}

		// 设置列信息
		c.columns[i] = csvColumn{
			Header:      header,
			ColumnIndex: i,
			FieldIndex:  i,
			Format:      format,
		}
	}

	return c, nil
}

// ReadFromReader从提供的读取器中解析
// CSV数据，将数据映射到相应的结构字段，并
// 通过通道传递结果。
func (c *Csv[T]) ReadFromReader(reader io.Reader) <-chan *T {
	// 行 chan
	rows := make(chan *T)

	go func() {
		defer close(rows)

		// 使用go内置csv库解析
		csvReader := csv.NewReader(reader)

		// 如果CSV有标题，请将列索引对齐以匹配列标
		// 题的顺序。
		if c.hasHeader {
			err := c.updateColumnIndexes(csvReader)
			if err != nil {
				c.Logger.Error("Unable to update the column indexes.", "error", err)
				return
			}
		}

		//开始读csv
		for {
			record, err := csvReader.Read()
			if err == io.EOF {
				break
			}

			if err != nil {
				c.Logger.Error("Unable to read row.", "error", err)
				break
			}

			// 创建一个结构体实例，并填充数据
			row := new(T)
			// 获取值对象
			rowValue := reflect.ValueOf(row).Elem()

			for _, column := range c.columns {
				// 在标题头里没找到
				if column.ColumnIndex == -1 {
					continue
				}

				// 设置值
				err := setReflectValue(rowValue.Field(column.FieldIndex),
					record[column.ColumnIndex], column.Format)
				if err != nil {
					c.Logger.Error("Unable to set value.", "error", err)
					return
				}
			}

			// 发送行数据
			rows <- row
		}
	}()

	// 返回chan
	return rows
}

// ReadFromFile从提供的文件名中解析CSV数据；
// 将数据映射到相应的struct字段，并交付
// 结果行通过通道。
func (c *Csv[T]) ReadFromFile(fileName string) (<-chan *T, error) {
	// 打开文件
	file, err := os.Open(filepath.Clean(fileName))
	if err != nil {
		return nil, err
	}

	// 创建 WaitGroup
	wg := &sync.WaitGroup{}
	// 等待函数，等待完成
	// wg是为了解决，一个函数返回chan, 同时还需要在chan关闭后做一些别的操作
	// 本来是可以使用defer的，在chan发送完之后做一些操作
	// 但是如果chan是在别的函数生成的，或者是当前函数的参数，无法在发送消息的时候再设置defer
	// 	只能再写一个chan, 将chan发送给当前chan,然后使用defer
	// 同时使用wg,可以把这个之后操作放在代码后面
	rows := Waitable(wg, c.ReadFromReader(file))

	go func() {
		// 等待为0
		wg.Wait()
		// 关闭文件
		err := file.Close()
		if err != nil {
			c.Logger.Error("Unable to close file.", "error", err)
		}
	}()

	return rows, nil
}

// AppendToFile将提供的数据行追加到指定文件的末尾，如果文件不存在则创建该文件。
// 在追加模式下，该函数假定现有文件的列顺序与给定行结构的字段顺序匹配，以确保数据结构的一致性。
func (c *Csv[T]) AppendToFile(fileName string, rows <-chan *T) error {
	// 打开文件
	file, err := os.OpenFile(filepath.Clean(fileName), os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}

	//  写入数据
	err = c.writeToWriter(file, false, rows)
	if err != nil {
		return err
	}
	// 关闭文件
	return file.Close()
}

// WriteToFile创建一个具有给定名称的新文件，并将提供的数据行写入该文件，覆盖任何现有内容。
func (c *Csv[T]) WriteToFile(fileName string, rows <-chan *T) error {
	// 打开文件
	file, err := os.OpenFile(filepath.Clean(fileName), os.O_CREATE|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}

	// 写入数据
	err = c.writeToWriter(file, true, rows)
	if err != nil {
		return err
	}

	// 关闭文件
	return file.Close()
}

// 将列索引对齐以匹配列标题的顺序。
func (c *Csv[T]) updateColumnIndexes(csvReader *csv.Reader) error {
	// 读一行
	headers, err := csvReader.Read()
	if err != nil {
		return err
	}

	// map key: 标题列名称， value: 标题列索引
	headerMap := make(map[string]int)
	for i, header := range headers {
		headerMap[header] = i
	}

	// 遍历csv对象列
	for i := range c.columns {
		// 尝试根据名称匹配列索引
		index, ok := headerMap[c.columns[i].Header]
		if !ok {
			// 匹配不到则index为-1
			index = -1
		}
		// 设置列对象索引
		c.columns[i].ColumnIndex = index
	}

	return nil
}

// writeToWriter将提供的数据行写入指定的写入器，并可选择包含或排除标题，以提高数据表示的灵活性。
func (c *Csv[T]) writeToWriter(writer io.Writer, writeHeader bool, rows <-chan *T) error {
	// 创建csv写入器
	csvWriter := csv.NewWriter(writer)

	// 如果写入列头
	if writeHeader {
		err := c.writeHeaderToCsvWriter(csvWriter)
		if err != nil {
			return err
		}
	}

	// 创建写入的对象
	record := make([]string, len(c.columns))

	for row := range rows {
		// 获取行值
		rowValue := reflect.ValueOf(row).Elem()

		// 遍历列对象
		for i, column := range c.columns {
			// 获取行[i] 每一列的值
			// 使用反射
			stringValue, err := getReflectValue(rowValue.Field(column.FieldIndex), column.Format)
			if err != nil {
				return err
			}
			// 行值赋值
			record[i] = stringValue
		}

		// 写入行
		err := csvWriter.Write(record)
		if err != nil {
			return err
		}
	}

	// 刷新缓冲
	csvWriter.Flush()

	// 返回写入过程中的错误
	return csvWriter.Error()
}

// writeHeaderToCsvWriter将CSV数据的列头写入指定的CSV写入器。
func (c *Csv[T]) writeHeaderToCsvWriter(csvWriter *csv.Writer) error {
	// 创建字符串切片
	header := make([]string, len(c.columns))

	for i, column := range c.columns {
		// 遍历标题头，赋值给切片
		header[i] = column.Header
	}
	// 写入标题头
	return csvWriter.Write(header)
}

// ReadFromCsvFile创建一个CSV实例，从提供的文件名解析CSV数据，将数据映射到相应的结构字段，并通过通道传递。
func ReadFromCsvFile[T any](fileName string, hasHeader bool) (<-chan *T, error) {
	c, err := NewCsv[T](hasHeader)
	if err != nil {
		return nil, err
	}

	return c.ReadFromFile(fileName)
}

// AppendOrWriteToCsvFile将提供的数据行写入指定的文件，如果现有文件存在，则追加到现有文件，如果不存在，则创建一个新文件。在追加模式下，该函数假定现有文件的列顺序与给定行结构的字段顺序匹配，以确保数据结构的一致性。
func AppendOrWriteToCsvFile[T any](fileName string, hasHeader bool, rows <-chan *T) error {
	c, err := NewCsv[T](hasHeader)
	if err != nil {
		return err
	}

	// 获取文件元数据
	stat, err := os.Stat(filepath.Clean(fileName))
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			// 如果报错并且错误不是文件不存在，返回错误
			return err
		}
	} else if stat.Size() > 0 {
		// 如果文件大小大于 0，追加文件
		return c.AppendToFile(fileName, rows)
	}
	// 写文件
	return c.WriteToFile(fileName, rows)
}
