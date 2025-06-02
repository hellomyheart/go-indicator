package trend

import "github.com/hellomyheart/go-indicator/helper"

// MovingSum 表示在指定期间内计算移动和的配置参数。 结构体
//
// Example:
//
//	sum := trend.NewMovingSum[float64]()
//	sum.Period = 20
type MovingSum[T helper.Number] struct {
	// Time period.
	Period int
}

// NewMovingSum 函数使用默认参数初始化一个新的移动求和实例。
// 默认周期为1
func NewMovingSum[T helper.Number]() *MovingSum[T] {
	return NewMovingSumWithPeriod[T](1)
}

// NewMovingSumWithPeriod 函数用给定的周期初始化一个新的移动求和实例。
func NewMovingSumWithPeriod[T helper.Number](period int) *MovingSum[T] {
	return &MovingSum[T]{
		Period: period,
	}
}

// Compute 函数接受一个数字通道，并计算指定时间段内的移动和。
func (m *MovingSum[T]) Compute(c <-chan T) <-chan T {
	// 将一个chan复制为两个chan
	cs := helper.Duplicate(c, 2)
	// 第二个chan进行数据右移，右移一个周期，并且补零
	cs[1] = helper.Shift(cs[1], m.Period, 0)

	// 定义移动和 0
	sum := T(0)

	// 移动和计算公式 sum +  chan1 i - chan2i (等价于 chan1 - （周期1）)
	// 不满一个周期的时候 移动和都是 chan1[i]
	// 第一个周期的最后一个移动和： 等价于，

	// 示例： 周期4， 数据是： 1 2 3 4 5 6 7 8
	// chan1 1  2  3  4  5  6  7  8
	// chan2 0  0  0  0  1  2  3  4  5  6  7  8
	// 移动和 1 3  6  10 14 18 22 26
	// 跳过前三个     10 14 18 22  26
	sums := helper.Operate(cs[0], cs[1], func(c, b T) T {
		sum = sum + c - b
		return sum
	})

	// 跳过周期-1个数据，因为前周期-1个数据的数据个数没有周期个
	return helper.Skip(sums, m.Period-1)
}

// IdlePeriod 是移动求和不会产生任何结果的初始阶段。是周期 -1
func (m *MovingSum[T]) IdlePeriod() int {
	return m.Period - 1
}
