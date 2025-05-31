package helper

// KeepNegatives 从一个数值流中筛选出负数并保留原始值，同时将正数替换为0
//
// Example:
//
//	c := helper.SliceToChan([]int{-10, 20, 4, -5})
//	negatives := helper.KeepPositives(c)
//	fmt.Println(helper.ChanToSlice(negatives)) // [-10, 0, 0, -5]
func KeepNegatives[T Number](c <-chan T) <-chan T {
	return Apply(c, func(n T) T {
		if n < 0 {
			return n
		}

		return 0
	})
}
