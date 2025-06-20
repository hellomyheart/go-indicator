package volatility_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/volatility"
)

func TestSuperTrend(t *testing.T) {
	type Data struct {
		High       float64
		Low        float64
		Close      float64
		SuperTrend float64
	}

	input, err := helper.ReadFromCsvFile[Data]("testdata/super_trend.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 4)
	highs := helper.Map(inputs[0], func(d *Data) float64 { return d.High })
	lows := helper.Map(inputs[1], func(d *Data) float64 { return d.Low })
	closings := helper.Map(inputs[2], func(d *Data) float64 { return d.Close })
	expectedSuperTrends := helper.Map(inputs[3], func(d *Data) float64 { return d.SuperTrend })

	superTrend := volatility.NewSuperTrend[float64]()
	actualSuperTrends := superTrend.Compute(highs, lows, closings)

	actualSuperTrends = helper.RoundDigits(actualSuperTrends, 2)
	expectedSuperTrends = helper.Skip(expectedSuperTrends, superTrend.IdlePeriod())

	err = helper.CheckEquals(actualSuperTrends, expectedSuperTrends)
	if err != nil {
		t.Fatal(err)
	}
}
