package helper

// Skip 跳过通道前count个元素
//
// Example:
//
//	c := helper.SliceToChan([]int{2, 4, 6, 8})
//	actual := helper.Skip(c, 2)
//	fmt.Println(helper.ChanToSlice(actual)) // [6, 8]
func Skip[T any](c <-chan T, count int) <-chan T {
	result := make(chan T, cap(c))

	go func() {
		for i := 0; i < count; i++ {
			_, ok := <-c
			if !ok {
				break
			}
		}

		Pipe(c, result)
	}()

	return result
}
