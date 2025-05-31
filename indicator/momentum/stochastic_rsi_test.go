package momentum_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/momentum"
)

func TestStochasticRsi(t *testing.T) {
	type Data struct {
		Close         float64
		StochasticRsi float64
	}

	input, err := helper.ReadFromCsvFile[Data]("testdata/stochastic_rsi.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 2)
	closings := helper.Map(inputs[0], func(d *Data) float64 { return d.Close })
	expected := helper.Map(inputs[1], func(d *Data) float64 { return d.StochasticRsi })

	stochasticRsi := momentum.NewStochasticRsi[float64]()
	actual := stochasticRsi.Compute(closings)
	actual = helper.RoundDigits(actual, 2)

	expected = helper.Skip(expected, stochasticRsi.IdlePeriod())

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
