package helper

// Divide 这段代码实现了两个通道中数值的逐元素除法运算。具体来说：
//
// Example:
//
//	ac := helper.SliceToChan([]int{2, 4, 6, 8, 10})
//	bc := helper.SliceToChan([]int{2, 1, 3, 2, 5})
//
//	division := helper.Divide(ac, bc)
//
//	fmt.Println(helper.ChanToSlice(division)) // [1, 4, 2, 4, 2]
func Divide[T Number](ac, bc <-chan T) <-chan T {
	return Operate(ac, bc, func(a, b T) T {
		return a / b
	})
}
