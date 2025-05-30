package helper

import "math"

// RoundDigit 将给定的float64数舍入到小数点后d位。
//
// Example:
//
//	n := helper.RoundDigit(10.1234, 2)
//	fmt.Println(n) // 10.12
func RoundDigit[T Number](n T, d int) T {
	// 计算缩放因子
	m := math.Pow(10, float64(d))
	// 缩放并四舍五入
	// 缩小回原数值范围
	return T(math.Round(float64(n)*m) / m)
}
