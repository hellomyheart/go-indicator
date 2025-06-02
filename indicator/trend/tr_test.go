package trend_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

func TestTr(t *testing.T) {
	type Data struct {
		High  float64
		Low   float64
		Close float64
		Atr   float64
	}

	input, err := helper.ReadFromCsvFile[Data]("testdata/tr.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 4)
	highs := helper.Map(inputs[0], func(d *Data) float64 { return d.High })
	lows := helper.Map(inputs[1], func(d *Data) float64 { return d.Low })
	closings := helper.Map(inputs[2], func(d *Data) float64 { return d.Close })
	expected := helper.Map(inputs[3], func(d *Data) float64 { return d.Atr })

	tr := trend.NewTr[float64]()
	actual := tr.Compute(highs, lows, closings)
	actual = helper.RoundDigits(actual, 2)

	expected = helper.Skip(expected, tr.IdlePeriod())

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
