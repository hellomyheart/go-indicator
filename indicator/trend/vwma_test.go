package trend_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

func TestVwma(t *testing.T) {
	type Data struct {
		Close  float64
		Volume int64
		Vwma   float64
	}

	input, err := helper.ReadFromCsvFile[Data]("testdata/vwma.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 3)
	closing := helper.Map(inputs[0], func(d *Data) float64 { return d.Close })
	volume := helper.Map(inputs[1], func(d *Data) float64 { return float64(d.Volume) })
	expected := helper.Map(inputs[2], func(d *Data) float64 { return d.Vwma })

	vwma := trend.NewVwma[float64]()

	actual := vwma.Compute(closing, volume)
	actual = helper.RoundDigits(actual, 2)

	expected = helper.Skip(expected, vwma.IdlePeriod())

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
