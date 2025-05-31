package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestHead(t *testing.T) {
	input := []int{2, 4, 6, 8}
	expected := helper.SliceToChan([]int{2, 4})

	actual := helper.Head(helper.SliceToChan(input), 2)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHeadEarly(t *testing.T) {
	input := []int{2}
	expected := helper.SliceToChan([]int{2})

	actual := helper.Head(helper.SliceToChan(input), 2)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
