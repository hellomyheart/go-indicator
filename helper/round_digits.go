package helper

// RoundDigits 获取一个数字的chan，并将其四舍五入到小数点后d位。
//
// Example:
//
//	c := helper.SliceToChan([]float64{10.1234, 5.678, 6.78, 8.91011})
//	rounded := helper.RoundDigits(c, 2)
//	fmt.Println(helper.ChanToSlice(rounded)) // [10.12, 5.68, 6.78, 8.91]
func RoundDigits[T Number](c <-chan T, d int) <-chan T {
	return Apply(c, func(n T) T {
		return RoundDigit(n, d)
	})
}
