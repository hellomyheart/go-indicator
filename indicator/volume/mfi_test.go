package volume_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/volume"
)

func TestMfi(t *testing.T) {
	type MfiData struct {
		High   float64
		Low    float64
		Close  float64
		Volume int64
		Mfi    float64
	}

	input, err := helper.ReadFromCsvFile[MfiData]("testdata/mfi.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 5)
	highs := helper.Map(inputs[0], func(m *MfiData) float64 { return m.High })
	lows := helper.Map(inputs[1], func(m *MfiData) float64 { return m.Low })
	closings := helper.Map(inputs[2], func(m *MfiData) float64 { return m.Close })
	volumes := helper.Map(inputs[3], func(m *MfiData) float64 { return float64(m.Volume) })
	expected := helper.Map(inputs[4], func(m *MfiData) float64 { return m.Mfi })

	mfi := volume.NewMfi[float64]()
	actual := helper.RoundDigits(mfi.Compute(highs, lows, closings, volumes), 2)
	expected = helper.Skip(expected, mfi.IdlePeriod())

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
