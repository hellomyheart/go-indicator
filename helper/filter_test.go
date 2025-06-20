package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := helper.SliceToChan([]int{2, 4, 6, 8, 10})

	actual := helper.Filter(helper.SliceToChan(input), func(n int) bool {
		return n%2 == 0
	})

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
