package trend_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

func TestBop(t *testing.T) {
	type BopData struct {
		Open  float64
		High  float64
		Low   float64
		Close float64
		Bop   float64
	}

	input, err := helper.ReadFromCsvFile[BopData]("testdata/bop.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 5)
	opening := helper.Map(inputs[0], func(row *BopData) float64 { return row.Open })
	high := helper.Map(inputs[1], func(row *BopData) float64 { return row.High })
	low := helper.Map(inputs[2], func(row *BopData) float64 { return row.Low })
	closing := helper.Map(inputs[3], func(row *BopData) float64 { return row.Close })
	expected := helper.Map(inputs[4], func(row *BopData) float64 { return row.Bop })

	bop := trend.NewBop[float64]()
	actual := bop.Compute(opening, high, low, closing)

	actual = helper.RoundDigits(actual, 0)

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
