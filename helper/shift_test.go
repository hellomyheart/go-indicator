package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestShift(t *testing.T) {
	input := helper.SliceToChan([]int{2, 4, 6, 8})
	expected := helper.SliceToChan([]int{0, 0, 0, 0, 2, 4, 6, 8})

	actual := helper.Shift(input, 4, 0)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
