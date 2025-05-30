package helper

// MapWithPrevious对输入通道中的每个元素应用转换函数，用转换后的值创建一个新的通道。
// 它维护先前结果的“内存”，允许转换函数同时考虑当前元素和先前转换的结果。这支持依赖于元素之间的累积状态或顺序依赖的函数。
//
// 例如:
//
//	sum := helper.MapWithPrevious(c, func(p, c int) int {
//		return p + c
//	}, 0)
func MapWithPrevious[F, T any](c <-chan F, f func(T, F) T, previous T) <-chan T {
	mc := make(chan T)

	go func() {
		defer close(mc)

		for n := range c {
			previous = f(previous, n)
			mc <- previous
		}
	}()

	return mc
}
