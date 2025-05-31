package helper

// Last 从输入通道中提取最后 N 个元素的功能
func Last[T any](c <-chan T, count int) <-chan T {
	result := make(chan T, cap(c))

	go func() {
		defer close(result)

		ring := NewRing[T](count)

		for n := range c {
			ring.Put(n)
		}

		for !ring.IsEmpty() {
			n, _ := ring.Get()
			result <- n
		}
	}()

	return result
}
