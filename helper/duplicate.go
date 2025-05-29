package helper

// chan复制器，将一个chan复制为多个chan,实现广播
//
// 比如:
//
//	expected := helper.SliceToChan([]float64{-10, 20, -4, -5})
//	outputs := helper.Duplicates[float64](helper.SliceToChan(expected), 2)
//
//	fmt.Println(helper.ChanToSlice(outputs[0])) // [-10, 20, -4, -5]
//	fmt.Println(helper.ChanToSlice(outputs[1])) // [-10, 20, -4, -5]
func Duplicate[T any](input <-chan T, count int) []<-chan T {
	// 因为类型转换问题，所以有两个变量
	// 双向通道（chan T）可以隐式转换为只读通道（<-chan T），
	// 但切片类型无法直接转换。例如：[]chan T 不能直接转换为 []<-chan T。
	// TODO: 看看有没有优化的方式
	outputs := make([]chan T, count)
	result := make([]<-chan T, count)

	for i := range outputs {
		outputs[i] = make(chan T, cap(input))
		result[i] = outputs[i]
	}

	go func() {
		for _, output := range outputs {
			defer close(output)
		}

		for n := range input {
			for _, output := range outputs {
				output <- n
			}
		}
	}()

	return result
}
