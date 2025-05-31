package helper

// DecrementBy 通道数据减法运算转换器，其核心作用是将输入通道中的每个数值元素减去指定值后输出到新通道。
//
// Example:
//
//	input := helper.SliceToChan([]int{1, 2, 3, 4})
//	substractOne := helper.DecrementBy(input, 1)
//	fmt.Println(helper.ChanToSlice(substractOne)) // [0, 1, 2, 3]
func DecrementBy[T Number](c <-chan T, d T) <-chan T {
	return Apply(c, func(n T) T {
		return n - d
	})
}
