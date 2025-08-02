package volatility

import (
	"math"

	"github.com/hellomyheart/go-indicator/helper"
)

const (
	// DefaultMovingStdPeriod 是移动标准偏差的默认时间段。
	DefaultMovingStdPeriod = 1
)

// MovingStd表示用于计算指定时间段内移动标准偏差的配置参数。
//
//	Std = Sqrt(1/Period * Sum(Pow(value - sma, 2)))
type MovingStd[T helper.Number] struct {
	// Time period.
	Period int
}

// NewMovingStd 函数 使用默认参数初始化一个新的Moving Standard Deviation实例。
func NewMovingStd[T helper.Number]() *MovingStd[T] {
	return NewMovingStdWithPeriod[T](DefaultMovingStdPeriod)
}

// NewMovingStdWithPeriod function initializes a new Moving Standard Deviation instance with the given period.
func NewMovingStdWithPeriod[T helper.Number](period int) *MovingStd[T] {
	return &MovingStd[T]{
		Period: period,
	}
}

// 计算函数接受一个数字通道，并计算指定期间的移动标准偏差。
func (m *MovingStd[T]) Compute(c <-chan T) <-chan T {
	// 创建一个结果通道
	// 缓存值为输入通道的容量
	result := make(chan T, cap(c))

	//	Std = Sqrt(1/Period * Sum(Pow(value - sma, 2)))
	go func() {
		// 推迟最后关闭通道
		defer close(result)

		// 创建一个环形缓存
		// 长度为周期值
		ring := helper.NewRing[T](m.Period)
		// 初始化sum 为0
		sum := T(0)

		// 消费输入通道
		for n := range c {
			// 循环求和 删掉周期前的值
			// 并且环形缓存放入当前值
			sum -= ring.Put(n)
			// 加上当前值
			// 其实就是移动和
			sum += n

			// 如果缓存区已经
			// 其实就是判断是否达到周期
			if ring.IsFull() {
				// sma 计算 移动平均 / 周期数
				sma := sum / T(m.Period)
				// 求和值2
				sum2 := T(0)

				// 遍历这一个周期的循环缓存
				for i := 0; i < m.Period; i++ {
					// sum2 的值为： 周期内每一个值 - 当前sma 的差的平方
					sum2 += T(math.Pow(float64(ring.At(i)-sma), 2))
				}

				// 将这个差的平方 除以周期数 开平方根
				// 结果就是标准差
				std := T(math.Sqrt(float64(sum2 / T(m.Period))))
				result <- std
			}
		}
	}()

	return result
}

// IdlePeriod 是不产生任何结果的初始阶段。
// 周期数 - 1
func (m *MovingStd[T]) IdlePeriod() int {
	return m.Period - 1
}
