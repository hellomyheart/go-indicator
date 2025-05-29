package helper

import "math"

// 对通道中的每个数值计算绝对值：
// 将数据转为 float64 类型，然后计算绝对值，再转为原始类型。
func Abs[T Number](c <-chan T) <-chan T {
	return Apply(c, func(n T) T {
		return T(math.Abs(float64(n)))
	})
}
