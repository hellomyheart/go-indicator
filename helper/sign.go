package helper

// Sign 接受一个float64值的通道，并返回它们的符号：-1表示负，0表示零，1表示正。
//
// Example:
//
//	c := helper.SliceToChan([]int{-10, 20, -4, 0})
//	sign := helper.Sign(c)
//	fmt.Println(helper.ChanToSlice(sign)) // [-1, 1, -1, 0]
func Sign[T Number](c <-chan T) <-chan T {
	return Apply(c, func(n T) T {
		if n > 0 {
			return 1
		} else if n < 0 {
			return -1
		}

		return 0
	})
}
