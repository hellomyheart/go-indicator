package trend

import "github.com/hellomyheart/go-indicator/helper"

// Ma represents the interface for the Moving Average (MA) indicators.
type Ma[T helper.Number] interface {
	// Compute function takes a channel of numbers and computes the MA.
	Compute(<-chan T) <-chan T

	// IdlePeriod is the initial period that MA won't yield any results.
	IdlePeriod() int

	// String is the string representation of the MA instance.
	String() string
}
