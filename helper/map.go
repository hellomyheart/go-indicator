package helper

// Map将给定的转换函数应用于输入通道中的每个元素，并返回一个包含转换后的值的新通道。转换函数接受任何类型值作为输入，并返回任何类型值作为输出。
//
// Example:
//
//	timesTwo := helper.Map(c, func(n int) int {
//		return n * 2
//	})
func Map[F, T any](c <-chan F, f func(F) T) <-chan T {
	mc := make(chan T)

	go func() {
		defer close(mc)

		for n := range c {
			mc <- f(n)
		}
	}()

	return mc
}
