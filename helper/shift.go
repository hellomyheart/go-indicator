package helper

// Shift 获取一个数字通道，按指定的计数向右移动它们，并用提供的填充值填充任何缺失的值。
//
// Example:
//
//	input := helper.SliceToChan([]int{2, 4, 6, 8})
//	output := helper.Shift(input, 4, 0)
//	fmt.Println(helper.ChanToSlice(output)) // [0, 0, 0, 0, 2, 4, 6, 8]
func Shift[T any](c <-chan T, count int, fill T) <-chan T {
	result := make(chan T, cap(c)+count)

	go func() {
		for i := 0; i < count; i++ {
			result <- fill
		}

		Pipe(c, result)
	}()

	return result
}
