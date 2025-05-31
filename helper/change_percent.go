package helper

// ChangePercent 百分比变化的功能，用于比较通道中当前值与指定位置前的值之间的百分比变化
//
// Example:
//
//	c := helper.ChanToSlice([]float64{1, 2, 5, 5, 8, 2, 1, 1, 3, 4})
//	actual := helper.ChangePercent(c, 2))
//	fmt.Println(helper.ChanToSlice(actual)) // [400, 150, 60, -60, -87.5, -50, 200, 300]
func ChangePercent[T Number](c <-chan T, before int) <-chan T {
	return MultiplyBy(ChangeRatio(c, before), 100)
}
