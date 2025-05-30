package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestSkip(t *testing.T) {
	input := helper.SliceToChan([]int{2, 4, 6, 8})
	expected := helper.SliceToChan([]int{6, 8})

	actual := helper.Skip(input, 2)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
