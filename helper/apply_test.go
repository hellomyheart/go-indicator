package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestApply(t *testing.T) {
	input := helper.SliceToChan([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	expected := helper.SliceToChan([]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20})

	actual := helper.Apply(input, func(n int) int {
		return n * 2
	})

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
