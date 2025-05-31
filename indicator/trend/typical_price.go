package trend

import "github.com/hellomyheart/go-indicator/helper"

// TypicalPrice represents the configuration parameters for calculating the Typical Price.
// It is another approximation of average price for each period and can be used as a
// filter for moving average systems.
//
//	Typical Price = (High + Low + Closing) / 3
type TypicalPrice[T helper.Number] struct{}

// NewTypicalPrice function initializes a new Typical Price instance with the default parameters.
func NewTypicalPrice[T helper.Number]() *TypicalPrice[T] {
	return &TypicalPrice[T]{}
}

// Compute function takes a channel of numbers and computes the Typical Price and the signal line.
func (*TypicalPrice[T]) Compute(high, low, closing <-chan T) <-chan T {
	return helper.DivideBy(
		helper.Add(
			helper.Add(high, low),
			closing,
		),
		3,
	)
}
