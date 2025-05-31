package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestChangeRatio(t *testing.T) {
	input := helper.SliceToChan([]float64{1, 2, 5, 5, 8, 2, 1, 1, 3, 4})
	expected := helper.SliceToChan([]float64{4, 1.5, 0.6, -0.6, -0.875, -0.5, 2, 3})

	actual := helper.ChangeRatio(input, 2)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
