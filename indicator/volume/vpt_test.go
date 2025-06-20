package volume_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/volume"
)

func TestVpt(t *testing.T) {
	type VptData struct {
		Close  float64
		Volume int64
		Vpt    float64
	}

	input, err := helper.ReadFromCsvFile[VptData]("testdata/vpt.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 3)
	closings := helper.Map(inputs[0], func(m *VptData) float64 { return m.Close })
	volumes := helper.Map(inputs[1], func(m *VptData) float64 { return float64(m.Volume) })
	expected := helper.Map(inputs[2], func(m *VptData) float64 { return m.Vpt })

	vpt := volume.NewVpt[float64]()
	actual := helper.RoundDigits(vpt.Compute(closings, volumes), 2)
	expected = helper.Skip(expected, vpt.IdlePeriod())

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
