package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestSubstract(t *testing.T) {
	ac := helper.SliceToChan([]int{2, 4, 6, 8, 10})
	bc := helper.SliceToChan([]int{1, 2, 3, 4, 5})

	expected := helper.SliceToChan([]int{1, 2, 3, 4, 5})

	actual := helper.Subtract(ac, bc)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
