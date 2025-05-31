package helper

import "math"

// Pow 对通道中每个元素进行指数运算的功能
//
// Example:
//
//	c := helper.SliceToChan([]int{2, 3, 5, 10})
//	squared := helper.Pow(c, 2)
//	fmt.Println(helper.ChanToSlice(squared)) // [4, 9, 25, 100]
func Pow[T Number](c <-chan T, y T) <-chan T {
	return Apply(c, func(n T) T {
		return T(math.Pow(float64(n), float64(y)))
	})
}
