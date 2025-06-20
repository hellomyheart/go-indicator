package trend_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

func TestMls(t *testing.T) {
	type Data struct {
		X float64
		Y float64
		M float64
		B float64
	}

	input, err := helper.ReadFromCsvFile[Data]("testdata/mls.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 4)
	x := helper.Map(inputs[0], func(d *Data) float64 { return d.X })
	y := helper.Map(inputs[1], func(d *Data) float64 { return d.Y })
	expectedM := helper.Map(inputs[2], func(d *Data) float64 { return d.M })
	expectedB := helper.Map(inputs[3], func(d *Data) float64 { return d.B })

	mls := trend.NewMlsWithPeriod[float64](5)

	actualM, actualB := mls.Compute(x, y)
	actualM = helper.RoundDigits(actualM, 2)
	actualB = helper.RoundDigits(actualB, 2)

	expectedM = helper.Skip(expectedM, mls.IdlePeriod())
	expectedB = helper.Skip(expectedB, mls.IdlePeriod())

	err = helper.CheckEquals(actualM, expectedM, actualB, expectedB)
	if err != nil {
		t.Fatal(err)
	}
}
