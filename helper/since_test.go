package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestSince(t *testing.T) {
	input := helper.SliceToChan([]int{1, 1, 2, 2, 2, 1, 2, 3, 3, 4})
	expected := helper.SliceToChan([]int{0, 1, 0, 1, 2, 0, 0, 0, 1, 0})

	actual := helper.Since[int, int](input)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
