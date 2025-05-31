package helper

// Since 计算自数字通道中最后一次更改值以来的周期数。
//  统计输入通道中元素连续出现次数的功能
func Since[T comparable, R Number](c <-chan T) <-chan R {
	first := true

	var last T
	var count R

	return Map(c, func(n T) R {
		if first || last != n { // 元素首次出现或值发生变化
			first = false
			last = n
			count = 0 // 重置计数器
		} else {
			count++ // 计数器加一
		}

		return count //重复出现的次数 = 出现的次数 - 1
	})
}
