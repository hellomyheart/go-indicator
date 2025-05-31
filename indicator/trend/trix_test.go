package trend_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

func TestTrix(t *testing.T) {
	type Data struct {
		Close float64
		Trix  float64
	}

	input, err := helper.ReadFromCsvFile[Data]("testdata/trix.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 2)
	closing := helper.Map(inputs[0], func(d *Data) float64 { return d.Close })
	expected := helper.Map(inputs[1], func(d *Data) float64 { return d.Trix })

	trix := trend.NewTrix[float64]()

	actual := trix.Compute(closing)
	actual = helper.RoundDigits(actual, 4)

	expected = helper.Skip(expected, trix.IdlePeriod())

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
