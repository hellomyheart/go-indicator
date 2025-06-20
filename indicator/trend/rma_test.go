package trend_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

func TestRma(t *testing.T) {
	type Data struct {
		Close float64
		Rma   float64
	}

	input, err := helper.ReadFromCsvFile[Data]("testdata/rma.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 2)
	closing := helper.Map(inputs[0], func(d *Data) float64 { return d.Close })
	expected := helper.Map(inputs[1], func(d *Data) float64 { return d.Rma })

	rma := trend.NewRma[float64]()
	rma.Period = 15

	actual := rma.Compute(closing)
	actual = helper.RoundDigits(actual, 2)

	expected = helper.Skip(expected, rma.IdlePeriod())

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
