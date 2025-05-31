package helper

// MultiplyBy 将通道中的每个数值乘以指定倍数，返回包含计算结果的新通道。
//
// Example:
//
//	c := helper.SliceToChan([]int{1, 2, 3, 4})
//	twoTimes := helper.MultiplyBy(c, 2)
//	fmt.Println(helper.ChanToSlice(twoTimes)) // [2, 4, 6, 8]
func MultiplyBy[T Number](c <-chan T, m T) <-chan T {
	return Apply(c, func(n T) T {
		return n * m
	})
}
