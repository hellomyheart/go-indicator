package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestKeepPositives(t *testing.T) {
	input := []int{-10, 20, 4, -5}
	expected := helper.SliceToChan([]int{0, 20, 4, 0})

	actual := helper.KeepPositives(helper.SliceToChan(input))

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
