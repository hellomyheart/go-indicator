package volatility_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/volatility"
)

func TestBollingerBandWidth(t *testing.T) {
	type Data struct {
		Close float64
		Width float64
	}

	input, err := helper.ReadFromCsvFile[Data]("testdata/bollinger_band_width.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 2)
	closings := helper.Map(inputs[0], func(d *Data) float64 { return d.Close })
	expected := helper.Map(inputs[1], func(d *Data) float64 { return d.Width })

	bbw := volatility.NewBollingerBandWidth[float64]()
	actual := bbw.Compute(closings)
	actual = helper.RoundDigits(actual, 2)

	expected = helper.Skip(expected, bbw.IdlePeriod())

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
