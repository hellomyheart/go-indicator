package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestAbs(t *testing.T) {
	input := helper.SliceToChan([]int{-10, 20, -4, -5})
	expected := helper.SliceToChan([]int{10, 20, 4, 5})

	actual := helper.Abs(input)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
