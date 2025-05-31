package trend_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

func TestTrimaWithOddPeriod(t *testing.T) {
	type Data struct {
		Close float64
		Trima float64
	}

	input, err := helper.ReadFromCsvFile[Data]("testdata/trima_odd.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 2)
	closing := helper.Map(inputs[0], func(d *Data) float64 { return d.Close })
	expected := helper.Map(inputs[1], func(d *Data) float64 { return d.Trima })

	trima := trend.NewTrima[float64]()
	trima.Period = 15

	actual := trima.Compute(closing)
	actual = helper.RoundDigits(actual, 2)

	expected = helper.Skip(expected, trima.IdlePeriod())

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTrimaWithEvenPeriod(t *testing.T) {
	type Data struct {
		Close float64
		Trima float64
	}

	input, err := helper.ReadFromCsvFile[Data]("testdata/trima_even.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 2)
	closing := helper.Map(inputs[0], func(d *Data) float64 { return d.Close })
	expected := helper.Map(inputs[1], func(d *Data) float64 { return d.Trima })

	trima := trend.NewTrima[float64]()
	trima.Period = 20

	actual := trima.Compute(closing)
	actual = helper.RoundDigits(actual, 2)

	expected = helper.Skip(expected, trima.IdlePeriod())

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
