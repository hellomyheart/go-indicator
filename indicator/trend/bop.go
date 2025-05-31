package trend

import "github.com/hellomyheart/go-indicator/helper"

// Bop gauges the strength of buying and selling forces using
// the Balance of Power (BoP) indicator. A positive BoP value
// suggests an upward trend, while a negative value indicates
// a downward trend. A BoP value of zero implies equilibrium
// between the two forces.
//
//	Formula: BOP = (Closing - Opening) / (High - Low)
type Bop[T helper.Number] struct{}

// NewBop function initializes a new BOP instance
// with the default parameters.
func NewBop[T helper.Number]() *Bop[T] {
	return &Bop[T]{}
}

// Compute processes a channel of open, high, low, and close values,
// computing the BOP for each entry.
func (*Bop[T]) Compute(opening, high, low, closing <-chan T) <-chan T {
	return helper.Divide(helper.Subtract(closing, opening), helper.Subtract(high, low))
}
