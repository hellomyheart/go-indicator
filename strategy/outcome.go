package strategy

import "github.com/hellomyheart/go-indicator/helper"

// Outcome simulates the potential result of executing the given actions based on the provided values.
func Outcome[T helper.Number](values <-chan T, actions <-chan Action) <-chan float64 {
	balance := 1.0
	shares := 0.0

	return helper.Operate(values, actions, func(value T, action Action) float64 {
		if balance > 0 && action == Buy {
			shares = balance / float64(value)
			balance = 0
		} else if shares > 0 && action == Sell {
			balance = shares * float64(value)
			shares = 0
		}

		return balance + (shares * float64(value)) - 1.0
	})
}
