package helper

// Head 从通道（channel）中提取前N个数值类型元素的功能
//
// Example:
//
//	c := helper.SliceToChan([]int{2, 4, 6, 8})
//	actual := helper.Head(c, 2)
//	fmt.Println(helper.ChanToSlice(actual)) // [2, 4]
func Head[T Number](c <-chan T, count int) <-chan T {
	result := make(chan T, cap(c))

	go func() {
		defer close(result)

		for i := 0; i < count; i++ {
			n, ok := <-c
			if !ok {
				break
			}

			result <- n
		}
	}()

	return result
}
