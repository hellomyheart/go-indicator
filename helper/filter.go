package helper

// Filter根据提供的谓词函数过滤输入通道中的项。谓词函数接受一个float64值作为输入，并返回一个布尔值，该值指示是否应将该值包含在输出通道中。
//
// 例如:
//
//	even := helper.Filter(c, func(n int) bool {
//	  return n%2 == 0
//	})
func Filter[T any](c <-chan T, p func(T) bool) <-chan T {
	fc := make(chan T)

	go func() {
		for n := range c {
			if p(n) {
				fc <- n
			}
		}

		close(fc)
	}()

	return fc
}
