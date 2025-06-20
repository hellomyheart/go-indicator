package trend_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

func TestApo(t *testing.T) {
	type ApoData struct {
		Close float64
		Apo   float64
	}

	input, err := helper.ReadFromCsvFile[ApoData]("testdata/apo.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	inputs := helper.Duplicate(input, 2)
	closing := helper.Map(inputs[0], func(a *ApoData) float64 { return a.Close })
	expected := helper.Map(inputs[1], func(a *ApoData) float64 { return a.Apo })

	apo := trend.NewApo[float64]()
	actual := helper.RoundDigits(apo.Compute(closing), 2)
	expected = helper.Skip(expected, apo.SlowPeriod-1)

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
