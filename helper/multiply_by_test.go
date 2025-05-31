package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestMultiplyBy(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := helper.SliceToChan([]int{2, 4, 6, 8})

	actual := helper.MultiplyBy(helper.SliceToChan(input), 2)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
