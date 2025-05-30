package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestMapWithPrevious(t *testing.T) {
	input := helper.SliceToChan([]int{1, 2, 3, 4})
	expected := helper.SliceToChan([]int{1, 3, 6, 10})

	actual := helper.MapWithPrevious(input, func(p, c int) int {
		return p + c
	}, 0)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
