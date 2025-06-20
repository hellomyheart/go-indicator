package volume_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/volume"
)

func TestVwap(t *testing.T) {
	type VwapData struct {
		Close  float64
		Volume int64
		Vwap   float64
	}

	input, err := helper.ReadFromCsvFile[VwapData]("testdata/vwap.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 3)
	closings := helper.Map(inputs[0], func(m *VwapData) float64 { return m.Close })
	volumes := helper.Map(inputs[1], func(m *VwapData) float64 { return float64(m.Volume) })
	expected := helper.Map(inputs[2], func(m *VwapData) float64 { return m.Vwap })

	vwap := volume.NewVwap[float64]()
	actual := helper.RoundDigits(vwap.Compute(closings, volumes), 2)
	expected = helper.Skip(expected, vwap.IdlePeriod())

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
