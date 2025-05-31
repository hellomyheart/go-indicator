package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestFirst(t *testing.T) {
	input := helper.SliceToChan([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	expected := helper.SliceToChan([]int{1, 2, 3, 4})

	actual := helper.First(input, 4)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFirstLessValues(t *testing.T) {
	input := helper.SliceToChan([]int{1, 2})
	expected := helper.SliceToChan([]int{1, 2})

	actual := helper.First(input, 4)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
