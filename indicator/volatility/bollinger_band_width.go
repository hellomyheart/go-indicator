package volatility

import "github.com/hellomyheart/go-indicator/helper"

// BollingerBandWidth表示计算布林带宽的配置参数。
// 它测量上波段和下波段之间的百分比差。随着
// 布林带变窄，并随着布林带变宽而增大。

// 在价格波动上升的时期，带宽变宽，在市场低迷的时期
// 波动带宽合约。

// 带宽=（上带-下带）/中布林带宽
//
// 例如:
//
//	bbw := NewBollingerBandWidth[float64]()
//	bbw.Compute(c)
type BollingerBandWidth[T helper.Number] struct {
	// Bollinger bands.
	BollingerBands *BollingerBands[T]
}

// NewBollingerBandWidth 函数 使用默认参数初始化一个新的布林带宽度实例。
func NewBollingerBandWidth[T helper.Number]() *BollingerBandWidth[T] {
	return &BollingerBandWidth[T]{
		BollingerBands: NewBollingerBands[T](),
	}
}

// NewBollingerBandWidthWithPeriod 函数 使用指定周期初始化一个新的布林带宽度实例。
func NewBollingerBandWidthWithPeriod[T helper.Number](period int) *BollingerBandWidth[T] {
	return &BollingerBandWidth[T]{
		BollingerBands: NewBollingerBandsWithPeriod[T](period),
	}
}

// 计算函数取一个数字通道并计算布林带宽度。
func (b *BollingerBandWidth[T]) Compute(c <-chan T) <-chan T {
	// 获取布林带指标
	upper, middle, lower := b.BollingerBands.Compute(c)

	// (上轨 - 下轨) / 中轨
	return helper.Divide(
		helper.Subtract(upper, lower),
		middle,
	)
}

// IdlePeriod  函数返回一个数字，表示该指标的初始延迟。
func (b *BollingerBandWidth[T]) IdlePeriod() int {
	return b.BollingerBands.IdlePeriod()
}
