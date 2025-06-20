package volatility_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/volatility"
)

func TestDonchianChannel(t *testing.T) {
	type Data struct {
		Close  float64
		Upper  float64
		Middle float64
		Lower  float64
	}

	input, err := helper.ReadFromCsvFile[Data]("testdata/donchian_channel.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 4)
	closings := helper.Map(inputs[0], func(d *Data) float64 { return d.Close })
	expectedUpper := helper.Map(inputs[1], func(d *Data) float64 { return d.Upper })
	expectedMiddle := helper.Map(inputs[2], func(d *Data) float64 { return d.Middle })
	expectedLower := helper.Map(inputs[3], func(d *Data) float64 { return d.Lower })

	dc := volatility.NewDonchianChannel[float64]()
	actualUpper, actualMiddle, actualLower := dc.Compute(closings)
	actualUpper = helper.RoundDigits(actualUpper, 2)
	actualMiddle = helper.RoundDigits(actualMiddle, 2)
	actualLower = helper.RoundDigits(actualLower, 2)

	expectedUpper = helper.Skip(expectedUpper, dc.IdlePeriod())
	expectedMiddle = helper.Skip(expectedMiddle, dc.IdlePeriod())
	expectedLower = helper.Skip(expectedLower, dc.IdlePeriod())

	err = helper.CheckEquals(actualUpper, expectedUpper, actualMiddle, expectedMiddle, actualLower, expectedLower)
	if err != nil {
		t.Fatal(err)
	}
}
