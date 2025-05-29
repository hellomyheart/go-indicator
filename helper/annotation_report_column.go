package helper

import "fmt"

// annotationReportColumn是注释报告列结构。
type annotationReportColumn struct {
	ReportColumn               // 接口组合 接口组合：通过匿名嵌入实现接口方法继承
	values       <-chan string // 数据源通道 通道驱动：使用只读通道实现实时数据流处理
}

// NewAnnotationReportColumn返回一个报告注释列的新实例。
// 工厂模式：封装实例创建逻辑
// 依赖注入：通过参数传递数据通道
func NewAnnotationReportColumn(values <-chan string) ReportColumn {
	return &annotationReportColumn{
		values: values,
	}
}

//Name返回报表列的名称。
func (*annotationReportColumn) Name() string {
	return ""
}

// Type返回string作为数据类型。
func (*annotationReportColumn) Type() string {
	return "string"
}

// Role返回报表列的角色。
func (*annotationReportColumn) Role() string {
	return "annotation"
}

// Value返回报表列的下一个数据值。
// 空值处理：返回"null"兼容JSON格式
// 字符串安全：使用%q自动处理特殊字符
func (c *annotationReportColumn) Value() string {
	value := <-c.values

	if value != "" {
		return fmt.Sprintf("%q", value)
	}

	return "null"
}
