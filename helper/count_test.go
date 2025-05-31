package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestCount(t *testing.T) {
	input := helper.SliceToChan([]int{1, 1, 1, 1})
	expected := helper.SliceToChan([]int{1, 2, 3, 4})

	actual := helper.Count(1, input)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
