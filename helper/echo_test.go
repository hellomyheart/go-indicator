package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestEcho(t *testing.T) {
	input := helper.SliceToChan([]int{2, 4, 6, 8})
	expected := helper.SliceToChan([]int{2, 4, 6, 8, 6, 8, 6, 8, 6, 8, 6, 8})

	actual := helper.Echo(input, 2, 4)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
