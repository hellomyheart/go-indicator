package helper

// ChangeRatio 计算数据流中每个元素与其N个位置前的元素之间的**比率变化（百分比变化）**功能。
//
// Example:
//
//	c := helper.ChanToSlice([]float64{1, 2, 5, 5, 8, 2, 1, 1, 3, 4})
//	actual := helper.ChangeRatio(c, 2))
//	fmt.Println(helper.ChanToSlice(actual)) // [400, 150, 60, -60, -87.5, -50, 200, 300]
func ChangeRatio[T Number](c <-chan T, before int) <-chan T {
	cs := Duplicate(c, 2)
	cs[1] = Buffered(cs[1], before)
	return Divide(Change(cs[0], before), cs[1])
}
