package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestKeepNegatives(t *testing.T) {
	input := []int{-10, 20, 4, -5}
	expected := helper.SliceToChan([]int{-10, 0, 0, -5})

	actual := helper.KeepNegatives(helper.SliceToChan(input))

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
