package helper

import (
	"math"
	"time"
)

// DaysBetween 计算给定两个时间之间的天数。
// 自然计算，类似前包后不包的方式(同天是0)
func DaysBetween(from, to time.Time) int {
	return int(math.Floor(to.Sub(from).Hours() / 24))
}
