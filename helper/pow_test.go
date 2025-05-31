package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestPow(t *testing.T) {
	input := helper.SliceToChan([]int{2, 3, 5, 10})
	expected := helper.SliceToChan([]int{4, 9, 25, 100})

	actual := helper.Pow(input, 2)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
