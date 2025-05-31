package helper

// DivideBy 对通道中每个元素进行固定值除法运算的功能。
//
// Example:
//
//	half := helper.DivideBy(helper.SliceToChan([]int{2, 4, 6, 8}), 2)
//	fmt.Println(helper.ChanToSlice(half)) // [1, 2, 3, 4]
func DivideBy[T Number](c <-chan T, d T) <-chan T {
	return Apply(c, func(n T) T {
		return n / d
	})
}
