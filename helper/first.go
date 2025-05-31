package helper

// First 从通道（channel）中提取前N个元素的功能
func First[T any](c <-chan T, count int) <-chan T {
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
		// 排空通道
		Drain(c)
	}()

	return result
}
