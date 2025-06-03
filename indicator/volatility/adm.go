package volatility

import (
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

const (
	// DefaultAdmPeriod adm的默认周期
	DefaultAdmPeriod = 14
)

// Adm dm的平均
//
// Example:
//
//	atr := volatility.NewAtr()
//	atr.Compute(highs, lows, closings)
type Adm[T helper.Number] struct {
	// Ma is the moving average for the Adm.
	Ma     trend.Ma[T]
	DmType bool
}

// NewAtrWithPeriod function initializes a new ATR instance with the given period.
func NewAdmWithPeriod[T helper.Number](period int, dmType bool) *Adm[T] {
	return NewAdmWithMa(trend.NewSmaWithPeriod[T](period), dmType)
}

// NewAtrWithMa function initializes a new ATR instance with the given moving average instance.
func NewAdmWithMa[T helper.Number](ma trend.Ma[T], dmType bool) *Adm[T] {
	return &Adm[T]{
		Ma:     ma,
		DmType: dmType,
	}
}

// Compute function takes a channel of numbers and computes the ATR over the specified period.
func (a *Adm[T]) Compute(f, l <-chan T) <-chan T {
	dm := trend.NewDm[T](a.DmType)

	adm := a.Ma.Compute(dm.Compute(f, l))
	return adm
}

// IdlePeriod is the initial period that Acceleration Bands won't yield any results.
func (a *Adm[T]) IdlePeriod() int {
	// Ma idle period and for using the previous closing.
	return a.Ma.IdlePeriod() + 1
}
