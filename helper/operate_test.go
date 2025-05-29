package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestOperate(t *testing.T) {
	ac := helper.SliceToChan([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	bc := helper.SliceToChan([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	expected := helper.SliceToChan([]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20})

	actual := helper.Operate(ac, bc, func(a, b int) int {
		return a + b
	})

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOperateFirstEnds(t *testing.T) {
	ac := helper.SliceToChan([]int{1, 2, 3, 4, 5, 6, 7, 8})
	bc := helper.SliceToChan([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	expected := helper.SliceToChan([]int{2, 4, 6, 8, 10, 12, 14, 16})

	actual := helper.Operate(ac, bc, func(a, b int) int {
		return a + b
	})

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOperateSecondEnds(t *testing.T) {
	ac := helper.SliceToChan([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	bc := helper.SliceToChan([]int{1, 2, 3, 4, 5, 6, 7, 8})

	expected := helper.SliceToChan([]int{2, 4, 6, 8, 10, 12, 14, 16})

	actual := helper.Operate(ac, bc, func(a, b int) int {
		return a + b
	})

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
