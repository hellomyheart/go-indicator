package helper

// IncrementBy 将输入通道中的每个元素按指定的增量值递增，并返回一个包含增量值的新通道。
//
// Example:
//
//	input := []int{1, 2, 3, 4}
//	actual := helper.IncrementBy(helper.SliceToChan(input), 1)
//	fmt.Println(helper.ChanToSlice(actual)) // [2, 3, 4, 5]
func IncrementBy[T Number](c <-chan T, i T) <-chan T {
	return Apply(c, func(n T) T {
		return n + i
	})
}
