package helper

import "math"

// Sqrt 对通道中每个元素计算平方根的功能
//
// Example:
//
//	c := helper.SliceToChan([]int{9, 81, 16, 100})
//	sqrt := helper.Sqrt(c)
//	fmt.Println(helper.ChanToSlice(sqrt)) // [3, 9, 4, 10]
func Sqrt[T Number](c <-chan T) <-chan T {
	return Apply(c, func(n T) T {
		return T(math.Sqrt(float64(n)))
	})
}
