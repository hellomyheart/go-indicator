package trend_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

func TestAroon(t *testing.T) {
	type AroonData struct {
		High float64
		Low  float64
		Up   float64
		Down float64
	}

	aroon := trend.NewAroon[float64]()

	input, err := helper.ReadFromCsvFile[AroonData]("testdata/aroon.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 4)
	high := helper.Map(inputs[0], func(row *AroonData) float64 { return row.High })
	low := helper.Map(inputs[1], func(row *AroonData) float64 { return row.Low })
	expectedUp := helper.Map(inputs[2], func(row *AroonData) float64 { return row.Up })
	expectedDown := helper.Map(inputs[3], func(row *AroonData) float64 { return row.Down })

	expectedUp = helper.Skip(expectedUp, aroon.Period-1)
	expectedDown = helper.Skip(expectedDown, aroon.Period-1)

	actualUp, actualDown := aroon.Compute(high, low)

	err = helper.CheckEquals(actualUp, expectedUp, actualDown, expectedDown)
	if err != nil {
		t.Fatal(err)
	}
}
