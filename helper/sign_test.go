package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestSign(t *testing.T) {
	input := helper.SliceToChan([]int{-10, 20, -4, 0})
	expected := helper.SliceToChan([]int{-1, 1, -1, 0})

	actual := helper.Sign(input)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
