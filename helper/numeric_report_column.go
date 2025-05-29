package helper

import "fmt"

// numericReportColumn是数字报表列结构体
type numericReportColumn[T Number] struct {
	ReportColumn          // 接口组合 接口组合：通过匿名嵌入实现接口方法继承
	name         string   //列名称
	values       <-chan T //数值流通道
}

// NewNumericReportColumn返回报告的数字数据列的新实例。
func NewNumericReportColumn[T Number](name string, values <-chan T) ReportColumn {
	return &numericReportColumn[T]{
		name:   name,
		values: values,
	}
}

// 返回列名
func (c *numericReportColumn[T]) Name() string {
	return c.name
}

// 返回列类型
func (*numericReportColumn[T]) Type() string {
	return "number"
}

// 返回列角色
func (*numericReportColumn[T]) Role() string {
	return "data"
}

// 获取列的下一个数据值
func (c *numericReportColumn[T]) Value() string {
	return fmt.Sprintf("%v", <-c.values)
}
