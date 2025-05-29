package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestCheckEqualsNotPairs(t *testing.T) {
	c := helper.SliceToChan([]int{1, 2, 3, 4})

	err := helper.CheckEquals(c)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCheckEqualsNotEndedTheSame(t *testing.T) {
	a := helper.SliceToChan([]int{1, 2, 3, 4})
	b := helper.SliceToChan([]int{1, 2})

	err := helper.CheckEquals(a, b)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCheckEqualsNotEquals(t *testing.T) {
	a := helper.SliceToChan([]int{1, 2, 3, 4})
	b := helper.SliceToChan([]int{1, 3, 3, 4})

	err := helper.CheckEquals(a, b)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCheckEquals(t *testing.T) {
	a := helper.SliceToChan([]int{1, 2, 3, 4})
	b := helper.SliceToChan([]int{1, 2, 3, 4})

	err := helper.CheckEquals(a, b)
	if err != nil {
		t.Fatal(err)
	}
}
