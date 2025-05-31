package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestDivide(t *testing.T) {
	ac := helper.SliceToChan([]int{2, 4, 6, 8, 10})
	bc := helper.SliceToChan([]int{2, 1, 3, 2, 5})

	expected := helper.SliceToChan([]int{1, 4, 2, 4, 2})

	actual := helper.Divide(ac, bc)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
