package volume_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/volume"
)

func TestAd(t *testing.T) {
	type AdData struct {
		High   float64
		Low    float64
		Close  float64
		Volume int64
		Ad     float64
	}

	input, err := helper.ReadFromCsvFile[AdData]("testdata/ad.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 5)
	highs := helper.Map(inputs[0], func(m *AdData) float64 { return m.High })
	lows := helper.Map(inputs[1], func(m *AdData) float64 { return m.Low })
	closings := helper.Map(inputs[2], func(m *AdData) float64 { return m.Close })
	volumes := helper.Map(inputs[3], func(m *AdData) float64 { return float64(m.Volume) })
	expected := helper.Map(inputs[4], func(m *AdData) float64 { return m.Ad })

	ad := volume.NewAd[float64]()
	actual := helper.RoundDigits(ad.Compute(highs, lows, closings, volumes), 2)
	expected = helper.Skip(expected, ad.IdlePeriod())

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
