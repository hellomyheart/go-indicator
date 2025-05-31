package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestDecrementBy(t *testing.T) {
	input := []int{2, 3, 4, 5}
	expected := helper.SliceToChan([]int{1, 2, 3, 4})

	actual := helper.DecrementBy(helper.SliceToChan(input), 1)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
