package strategy_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/strategy"
)

func TestOutcome(t *testing.T) {
	values := helper.SliceToChan([]float64{
		10, 15, 12, 12, 18,
		20, 22, 25, 24, 20,
	})

	actions := helper.SliceToChan([]strategy.Action{
		strategy.Hold, strategy.Hold, strategy.Buy, strategy.Buy, strategy.Hold,
		strategy.Hold, strategy.Hold, strategy.Sell, strategy.Hold, strategy.Hold,
	})

	expected := helper.SliceToChan([]float64{
		0, 0, 0, 0, 0.5,
		0.67, 0.83, 1.08, 1.08, 1.08,
	})

	actual := helper.RoundDigits(strategy.Outcome(values, actions), 2)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
