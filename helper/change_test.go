package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestChange(t *testing.T) {
	input := helper.SliceToChan([]int{1, 2, 5, 5, 8, 2, 1, 1, 3, 4})
	expected := helper.SliceToChan([]int{4, 3, 3, -3, -7, -1, 2, 3})

	actual := helper.Change(input, 2)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
