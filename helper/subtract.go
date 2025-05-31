package helper

// Subtract Subtract 函数的作用是 对两个通道中的数值进行逐元素减法运算，返回包含计算结果的新通道。
//
// Example:
//
//	ac := helper.SliceToChan([]int{2, 4, 6, 8, 10})
//	bc := helper.SliceToChan([]int{1, 2, 3, 4, 5})
//	actual := helper.Subtract(ac, bc)
//	fmt.Println(helper.ChanToSlice(actual)) // [1, 2, 3, 4, 5]
func Subtract[T Number](ac, bc <-chan T) <-chan T {
	return Operate(ac, bc, func(a, b T) T {
		return a - b
	})
}
