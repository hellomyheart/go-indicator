package volatility

import (
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

const (
	// DefaultAtrPeriod atr的默认周期
	DefaultAtrPeriod = 14
)

// Atr 表示计算ATR （Average True Range）的配置参数。它是一种技术分析指标，通过分解该时期的整个股票价格范围来衡量市场波动性。
//
//	TR = Max((High - Low), (High - Previous Closing), (Previous Closing - Low))
//	ATR = MA TR
//
// By default, SMA is used as the MA.
//
// Example:
//
//	atr := volatility.NewAtr()
//	atr.Compute(highs, lows, closings)
type Atr[T helper.Number] struct {
	// Ma is the moving average for the ATR.
	Ma trend.Ma[T]
}

// NewAtr function initializes a new ATR instance with the default parameters.
func NewAtr[T helper.Number]() *Atr[T] {
	return NewAtrWithPeriod[T](DefaultAtrPeriod)
}

// NewAtrWithPeriod function initializes a new ATR instance with the given period.
func NewAtrWithPeriod[T helper.Number](period int) *Atr[T] {
	return NewAtrWithMa(trend.NewSmaWithPeriod[T](period))
}

// NewAtrWithMa function initializes a new ATR instance with the given moving average instance.
func NewAtrWithMa[T helper.Number](ma trend.Ma[T]) *Atr[T] {
	return &Atr[T]{
		Ma: ma,
	}
}

// Compute function takes a channel of numbers and computes the ATR over the specified period.
func (a *Atr[T]) Compute(highs, lows, closings <-chan T) <-chan T {
	tr := NewTr[T]()

	atr := a.Ma.Compute(tr.Compute(highs, lows, closings))
	return atr
}

// IdlePeriod is the initial period that Acceleration Bands won't yield any results.
func (a *Atr[T]) IdlePeriod() int {
	// Ma idle period and for using the previous closing.
	return a.Ma.IdlePeriod() + 1
}
