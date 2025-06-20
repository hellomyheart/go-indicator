package volatility

import (
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

const (
	// DefaultDonchianChannelPeriod is the default period for the Donchian Channel.
	DefaultDonchianChannelPeriod = 20
)

// DonchianChannel represents the configuration parameters for calculating the Donchian Channel (DC). It calculates
// three lines generated by moving average calculations that comprise an indicator formed by upper and lower bands
// around a midrange or median band. The upper band marks the highest price of an asset while the lower band marks
// the lowest price of an asset, and the area between the upper and lower bands represents the Donchian Channel.
//
//	Upper Channel = Mmax(period, closings)
//	Lower Channel = Mmin(period, closings)
//	Middle Channel = (Upper Channel + Lower Channel) / 2
//
// Example:
//
//	dc := volatility.NewDonchianChannel[float64]()
//	result := dc.Compute(values)
type DonchianChannel[T helper.Number] struct {
	// Max is the Moving Max instance.
	Max *trend.MovingMax[T]

	// Min is the Moving Min instance.
	Min *trend.MovingMin[T]
}

// NewDonchianChannel function initializes a new Donchian Channel instance with the default parameters.
func NewDonchianChannel[T helper.Number]() *DonchianChannel[T] {
	return NewDonchianChannelWithPeriod[T](DefaultDonchianChannelPeriod)
}

// NewDonchianChannelWithPeriod function initializes a new Donchian Channel instance with the given period.
func NewDonchianChannelWithPeriod[T helper.Number](period int) *DonchianChannel[T] {
	return &DonchianChannel[T]{
		Max: trend.NewMovingMaxWithPeriod[T](period),
		Min: trend.NewMovingMinWithPeriod[T](period),
	}
}

// Compute function takes a channel of numbers and computes the Donchian Channel over the specified period.
func (d *DonchianChannel[T]) Compute(c <-chan T) (<-chan T, <-chan T, <-chan T) {
	closings := helper.Duplicate(c, 2)

	uppers := helper.Duplicate(
		d.Max.Compute(closings[0]),
		2,
	)

	lowers := helper.Duplicate(
		d.Min.Compute(closings[1]),
		2,
	)

	middle := helper.DivideBy(
		helper.Add(uppers[0], lowers[0]),
		2,
	)

	return uppers[1], middle, lowers[1]
}

// IdlePeriod is the initial period that Donchian Channel won't yield any results.
func (d *DonchianChannel[T]) IdlePeriod() int {
	return d.Max.IdlePeriod()
}
