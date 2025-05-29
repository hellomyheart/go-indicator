package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestMap(t *testing.T) {
	type Row struct {
		High int
		Low  int
	}

	input := []Row{
		{High: 10, Low: 5},
		{High: 20, Low: 15},
	}

	expected := helper.SliceToChan([]int{5, 15})

	actual := helper.Map(helper.SliceToChan(input), func(r Row) int {
		return r.Low
	})

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
