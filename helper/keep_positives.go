package helper

// KeepPositives 从一个数值流中筛选出正数并保留原始值，同时将负数替换为0
//
// Example:
//
//	c := helper.SliceToChan([]int{-10, 20, 4, -5})
//	positives := helper.KeepPositives(c)
//	fmt.Println(helper.ChanToSlice(positives)) // [0, 20, 4, 0]
func KeepPositives[T Number](c <-chan T) <-chan T {
	return Apply(c, func(n T) T {
		if n > 0 {
			return n
		}

		return 0
	})
}
