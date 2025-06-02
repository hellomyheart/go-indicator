package trend

import (
	"math"

	"github.com/hellomyheart/go-indicator/helper"
)

// tr 表示计算TR （True Range）。它是一种技术分析指标，通过分解该时期的整个股票价格范围来衡量市场波动性。
//
//	TR = Max((High - Low), (High - Previous Closing), (Previous Closing - Low))
//
// Example:
//
//	tr := volatility.Newtr()
//	tr.Compute(highs, lows, closings)
type Tr[T helper.Number] struct {
}

// NewTr 函数使用默认参数初始化新的TR实例。
func NewTr[T helper.Number]() *Tr[T] {
	return &Tr[T]{}
}

// Compute 函数接受一个数字通道，并计算指定时间段内的TR。
func (a *Tr[T]) Compute(highs, lows, closings <-chan T) <-chan T {
	// 使用之前的收盘价，跳过高点和低点一个。
	highs = helper.Skip(highs, 1)
	lows = helper.Skip(lows, 1)

	tr := helper.Operate3(highs, lows, closings, func(high, low, closing T) T {
		return T(math.Max(float64(high-low), math.Max(float64(high-closing), float64(closing-low))))
	})
	return tr
}

// IdlePeriod is 不会产生任何结果的过程
// TR是 1
func (a *Tr[T]) IdlePeriod() int {
	return 1
}
