package trend_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

func TestMovingSum(t *testing.T) {
	input := helper.SliceToChan([]int{-10, 20, -4, -5, 1, 5, 8, 10, -20, 4})
	expected := helper.SliceToChan([]int{1, 12, -3, 9, 24, 3, 2})

	sum := trend.NewMovingSum[int]()
	sum.Period = 4

	actual := sum.Compute(input)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
