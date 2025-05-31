package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestIncrementBy(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := helper.SliceToChan([]int{2, 3, 4, 5})

	actual := helper.IncrementBy(helper.SliceToChan(input), 1)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
