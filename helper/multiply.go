package helper

// Multiply 对两个通道中的数值进行逐元素乘法运算，返回包含计算结果的新通道。
//
// Example:
//
//	ac := helper.SliceToChan([]int{1, 4, 2, 4, 2})
//	bc := helper.SliceToChan([]int{2, 1, 3, 2, 5})
//
//	multiplication := helper.Multiply(ac, bc)
//
//	fmt.Println(helper.ChanToSlice(multiplication)) // [2, 4, 6, 8, 10]
func Multiply[T Number](ac, bc <-chan T) <-chan T {
	return Operate(ac, bc, func(a, b T) T {
		return a * b
	})
}
