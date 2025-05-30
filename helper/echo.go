package helper

// Echo接受一个数字通道，在末尾重复指定的数字计数。
// 例如：
//
//	input := helper.SliceToChan([]int{2, 4, 6, 8})
//	output := helper.Echo(input, 2, 4))
//	fmt.Println(helper.ChanToSlice(output)) // [2, 4, 6, 8, 6, 8, 6, 8, 6, 8, 6, 8]
func Echo[T any](input <-chan T, last, count int) <-chan T {
	output := make(chan T)
	//使用 NewRing[T](last) 创建环形缓冲区，仅保留最近的 last 个元素
	memory := NewRing[T](last)

	go func() {
		defer close(output)

		// 先输出原有chan
		for n := range input {
			// 放入缓存区
			memory.Put(n)
			output <- n
		}

		// 输出环形缓冲区
		// 第一层循环 循环次数
		for i := 0; i < count; i++ {
			// 第二层循环 每次最后个数
			for j := 0; j < last; j++ {
				output <- memory.At(j)
			}
		}
	}()

	return output
}
