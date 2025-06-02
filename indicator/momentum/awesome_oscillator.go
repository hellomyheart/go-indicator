package momentum

import (
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

// 绝密振荡器 AO
const (
	// DefaultAwesomeOscillatorShortPeriod 是AO的默认短周期。
	DefaultAwesomeOscillatorShortPeriod = 5

	// DefaultAwesomeOscillatorLongPeriod 是AO的默认长周期。
	DefaultAwesomeOscillatorLongPeriod = 34
)

// AwesomeOscillator 表示计算AO （Awesome Oscillator）的配置参数。它通过比较短期价格走势（5个周期平均值）和长期趋势（34个周期平均值）来衡量市场势头。它在零线附近的价值反映了上方的看涨和下方的看跌。越过零线可能预示着潜在的趋势逆转。交易者使用AO来确认现有趋势，确定进入/退出点，并了解动量变化。
//
//	中位价格 = ((Low + High) / 2).
//	AO = 5周期均值 - 34周期均值
//
// Example:
//
//	ao := momentum.AwesomeOscillator[float64]()
//	values := ao.Compute(lows, highs)
type AwesomeOscillator[T helper.Number] struct {
	// ShortSma is the SMA for the short period.
	ShortSma *trend.Sma[T]

	// LongSma is the SMA for the long period.
	LongSma *trend.Sma[T]
}

// NewAwesomeOscillator function initializes a new Awesome Oscillator instance.
func NewAwesomeOscillator[T helper.Number]() *AwesomeOscillator[T] {
	return &AwesomeOscillator[T]{
		ShortSma: trend.NewSmaWithPeriod[T](DefaultAwesomeOscillatorShortPeriod),
		LongSma:  trend.NewSmaWithPeriod[T](DefaultAwesomeOscillatorLongPeriod),
	}
}

// Compute function takes a channel of numbers and computes the AwesomeOscillator.
func (a *AwesomeOscillator[T]) Compute(highs, lows <-chan T) <-chan T {
	medianSplice := helper.Duplicate(
		helper.DivideBy(
			helper.Add(highs, lows),
			2,
		),
		2,
	)

	shortSma := a.ShortSma.Compute(medianSplice[0])
	longSma := a.LongSma.Compute(medianSplice[1])

	shortSma = helper.Skip(shortSma, a.LongSma.IdlePeriod()-a.ShortSma.IdlePeriod())

	return helper.Subtract(
		shortSma,
		longSma,
	)
}

// IdlePeriod is the initial period that Awesome Oscillator won't yield any results.
func (a *AwesomeOscillator[T]) IdlePeriod() int {
	return a.LongSma.IdlePeriod()
}
