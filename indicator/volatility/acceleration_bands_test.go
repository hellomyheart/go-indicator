package volatility_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/volatility"
)

func TestAccelerationBands(t *testing.T) {
	type Data struct {
		High   float64
		Low    float64
		Close  float64
		Upper  float64
		Middle float64
		Lower  float64
	}

	input, err := helper.ReadFromCsvFile[Data]("testdata/acceleration_bands.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 6)
	highs := helper.Map(inputs[0], func(d *Data) float64 { return d.High })
	lows := helper.Map(inputs[1], func(d *Data) float64 { return d.Low })
	closings := helper.Map(inputs[2], func(d *Data) float64 { return d.Close })
	upper := helper.Map(inputs[3], func(d *Data) float64 { return d.Upper })
	middle := helper.Map(inputs[4], func(d *Data) float64 { return d.Middle })
	lower := helper.Map(inputs[5], func(d *Data) float64 { return d.Lower })

	ab := volatility.NewAccelerationBands[float64]()
	actualUpper, actualMiddle, actualLower := ab.Compute(highs, lows, closings)
	actualUpper = helper.RoundDigits(actualUpper, 2)
	actualMiddle = helper.RoundDigits(actualMiddle, 2)
	actualLower = helper.RoundDigits(actualLower, 2)

	upper = helper.Skip(upper, ab.IdlePeriod())
	middle = helper.Skip(middle, ab.IdlePeriod())
	lower = helper.Skip(lower, ab.IdlePeriod())

	err = helper.CheckEquals(actualUpper, upper, actualMiddle, middle, actualLower, lower)
	if err != nil {
		t.Fatal(err)
	}
}
