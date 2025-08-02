package volatility

import (
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

const (
	// DefaultBollingerBandsPeriod是布林带的默认周期。
	DefaultBollingerBandsPeriod = 20
)

// BollingerBands 表示计算布林带的配置参数。这是一个技术分析工具，
// 用于衡量市场的波动率并识别超买和超卖条件。返回上轨、中轨和下轨。
//
// 中轨 = 20周期简单移动平均线（SMA）
// 上轨 = 20周期SMA + 2倍（20周期标准差）
// 下轨 = 20周期SMA - 2倍（20周期标准差）
//
// 示例：
//
// bollingerBands := NewBollingerBands[float64]()
// bollingerBands.Compute(values)
type BollingerBands[T helper.Number] struct {
	// Time period.
	Period int
}

// NewBollingerBands 函数使用默认参数初始化一个新的布林带实例。
func NewBollingerBands[T helper.Number]() *BollingerBands[T] {
	return NewBollingerBandsWithPeriod[T](DefaultBollingerBandsPeriod)
}

// NewBollingerBandsWithPeriod 函数 用给定的周期初始化一个新的布林带实例。
func NewBollingerBandsWithPeriod[T helper.Number](period int) *BollingerBands[T] {
	return &BollingerBands[T]{
		Period: period,
	}
}

// 计算函数接受一个数字通道，并计算指定期间的布林带。
// 功能：输入一个数值通道 c，输出三条布林带（上轨、中轨、下轨）的数值通道。
func (b *BollingerBands[T]) Compute(c <-chan T) (<-chan T, <-chan T, <-chan T) {
	// 复制通道以便在计算中使用
	cs := helper.Duplicate(c, 2)
	// 创建一个简单移动平均线(sma)结构体实例 这个就是布林带的中轨
	sma := trend.NewSmaWithPeriod[T](b.Period)
	// 创建一个标准差(std)计算实例
	std := NewMovingStdWithPeriod[T](b.Period)

	// 将sma(中轨)复制三份
	middleBands := helper.Duplicate(
		sma.Compute(cs[0]),
		3,
	)

	// 将标准差复制2分
	std2s := helper.Duplicate(
		helper.MultiplyBy(
			std.Compute(cs[1]),
			2,
		),
		2,
	)

	// 上轨 = 中轨 + 2倍的标准差
	upperBand := helper.Add(
		middleBands[0],
		std2s[0],
	)

	// 下轨 = 中轨 - 2倍的标准差
	lowerBand := helper.Subtract(
		middleBands[1],
		std2s[1],
	)
	// 返回上轨、中轨和下轨
	return upperBand, middleBands[2], lowerBand
}

// IdlePeriod 是布林带不会产生任何结果的初始阶段。
// 周期数 - 1，因为在计算布林带时需要至少一个完整的周期来计算中轨和标准差。
func (b *BollingerBands[T]) IdlePeriod() int {
	return b.Period - 1
}
