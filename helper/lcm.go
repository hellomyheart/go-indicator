package helper

// Lcm 计算给定数字的最小公倍数。
func Lcm(values ...int) int {

	// 赋值最小公倍数为第一个元素
	lcm := values[0]

	for i := 1; i < len(values); i++ {
		lcm = (values[i] * lcm) / Gcd(values[i], lcm)
	}

	return lcm
}
