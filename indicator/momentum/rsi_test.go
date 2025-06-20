package momentum_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/momentum"
)

func TestRsi(t *testing.T) {
	type Data struct {
		Close float64
		Rsi   float64
	}

	input, err := helper.ReadFromCsvFile[Data]("testdata/rsi.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 2)
	closings := helper.Map(inputs[0], func(d *Data) float64 { return d.Close })
	expectedRsi := helper.Map(inputs[1], func(d *Data) float64 { return d.Rsi })

	rsi := momentum.NewRsi[float64]()
	actualRsi := rsi.Compute(closings)
	actualRsi = helper.RoundDigits(actualRsi, 2)

	expectedRsi = helper.Skip(expectedRsi, rsi.IdlePeriod())

	err = helper.CheckEquals(actualRsi, expectedRsi)
	if err != nil {
		t.Fatal(err)
	}
}
