package helper

import (
	// Go embed report template.
	_ "embed"
	"io"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

// 使用 Go 1.16+ 的 embed 特性
// 将 HTML 模板文件编译进二进制
// 保证报告模板与代码强关联
//
//go:embed "report.tmpl"
var reportTmpl string

const (
	//使用 Go 的时间格式化语法（以参考时间 2006-01-02 为例）
	DefaultReportDateFormat = "2006-01-02"
)

// 核心接口：ReportColumn
//
// ReportColumn定义所有报表数据列必须实现的接口。
//该接口确保可以使用不同类型的数据列
//在报表生成过程中保持一致。

// 设计模式
// 策略模式：允许不同列类型实现统一接口
// 迭代器模式：通过 Value() 提供流式数据访问
// 装饰器模式：可组合基础列功能扩展新类型
type ReportColumn interface {
	// 报告列的名称
	Name() string

	// 报告列的数据类型
	Type() string

	// 报告列的角色
	Role() string

	// value 返回报告列的下一个值
	Value() string
}

// 核心结构体：Report
// 报告生成一个HTML文件，其中包含一个交互式图表
// 可视化地表示所提供的数据和注释。
//
// 生成的HTML文件可以在web浏览器中打开进行浏览
// 数据可视化，与图表元素交互，并查看
// 关联的注解。

type Report struct {
	Title       string           //报告标题
	Date        <-chan time.Time //时间轴数据
	Columns     []ReportColumn   // 数据列集合
	Views       [][]int          // 图表视图配置
	DateFormat  string           // 日期格式
	GeneratedOn string           // 生成时间戳
}

// NewReport以一个通道的时间作为时间轴，并返回一个新的
// 报告结构的实例。此实例稍后可用于
// 添加数据和注释，然后生成报告。

func NewReport(title string, date <-chan time.Time) *Report {
	return &Report{
		Title:   title,
		Date:    date,
		Columns: []ReportColumn{},
		Views: [][]int{
			{},
		},
		DateFormat:  DefaultReportDateFormat,
		GeneratedOn: time.Now().String(),
	}
}

// AddChart向报表添加一个新图表并返回其惟一值
//标识符。此标识符稍后可用于引用
//添加列。

// 添加新图表区域
// 返回图表ID（用于后续操作）

func (r *Report) AddChart() int {
	r.Views = append(r.Views, []int{})
	return len(r.Views) - 1
}

// AddColumn添加一个新的数据列到指定的图表。如果没有
//指定的，它将被添加到主图表。

// 支持多图表绑定
// 列ID自动递增
// 支持灵活的列-图表映射
func (r *Report) AddColumn(column ReportColumn, charts ...int) {
	r.Columns = append(r.Columns, column)
	columnID := len(r.Columns)

	if len(charts) == 0 {
		charts = append(charts, 0)
	}

	for _, chartID := range charts {
		r.Views[chartID] = append(r.Views[chartID], columnID)
	}
}

// WriteToWriter将报表内容写入提供的io.Writer。
// 这允许将报告发送到不同的目的地，例如
// 作为文件，网络套接字，甚至是标准输出。
func (r *Report) WriteToWriter(writer io.Writer) error {
	tmpl, err := template.New("report").Parse(reportTmpl)
	if err != nil {
		return err
	}

	return tmpl.Execute(writer, r)
}

// WriteToFile将生成的报表内容写入一个文件
// 指定的名称。这允许用户方便地保存
// 报告供以后查看或分析。
func (r *Report) WriteToFile(fileName string) error {
	file, err := os.Create(filepath.Clean(fileName))
	if err != nil {
		return err
	}

	err = r.WriteToWriter(file)
	if err != nil {
		return err
	}

	return file.Close()
}
