package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestChangePercent(t *testing.T) {
	input := helper.SliceToChan([]float64{1, 2, 5, 5, 8, 2, 1, 1, 3, 4})
	expected := helper.SliceToChan([]float64{400, 150, 60, -60, -87.5, -50, 200, 300})

	actual := helper.ChangePercent(input, 2)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
