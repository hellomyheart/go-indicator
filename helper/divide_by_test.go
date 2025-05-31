package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestDivideBy(t *testing.T) {
	input := []int{2, 4, 6, 8}
	expected := helper.SliceToChan([]int{1, 2, 3, 4})

	actual := helper.DivideBy(helper.SliceToChan(input), 2)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
