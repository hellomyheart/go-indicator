package trend_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/trend"
)

func TestTsi(t *testing.T) {
	type Data struct {
		Close float64
		Tsi   float64
	}

	input, err := helper.ReadFromCsvFile[Data]("testdata/tsi.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 2)
	closing := helper.Map(inputs[0], func(d *Data) float64 { return d.Close })
	expected := helper.Map(inputs[1], func(d *Data) float64 { return d.Tsi })

	tsi := trend.NewTsi[float64]()
	actual := tsi.Compute(closing)

	actual = helper.RoundDigits(actual, 2)
	expected = helper.Skip(expected, tsi.IdlePeriod())

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTsiString(t *testing.T) {
	expected := "TSI(EMA(1),EMA(2))"
	actual := trend.NewTsiWith[float64](1, 2).String()

	if actual != expected {
		t.Fatalf("actual %v expected %v", actual, expected)
	}
}
