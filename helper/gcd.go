package helper

// 计算多个整数的最大公约数

// 算法
// 欧几里得算法（辗转相除法）：

// 通过循环取余运算 gcd, value = value, gcd % value 直到 value == 0，此时 gcd 即为两个数的最大公约数。
// 对多个数依次计算累积 GCD（如 GCD(GCD(a,b), c)）。
func Gcd(values ...int) int {

	// 赋值最大公约数为第一个元素
	gcd := values[0]

	// 遍历数字
	for i := 1; i < len(values); i++ {
		value := values[i]

		for value > 0 {
			gcd, value = value, gcd%value
		}

		if gcd == 1 {
			break
		}
	}

	return gcd
}
