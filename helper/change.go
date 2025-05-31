package helper

// Change 计算数据流中每个元素与其N个位置前的元素之间的差值功能。
// 返回chan的长度是 len(input) - before
//
// Example:
//
//	input := []int{1, 2, 5, 5, 8, 2, 1, 1, 3, 4}
//	output := helper.Change(helper.SliceToChan(input), 2)
//	fmt.Println(helper.ChanToSlice(output)) // [4, 3, 3, -3, -7, -1, 2, 3]
func Change[T Number](c <-chan T, before int) <-chan T {
	cs := Duplicate(c, 2)
	cs[0] = Buffered(cs[0], before)
	cs[1] = Skip(cs[1], before)

	return Subtract(cs[1], cs[0])
}
