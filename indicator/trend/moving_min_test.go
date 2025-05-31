package trend_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

func TestMovingMin(t *testing.T) {
	input := helper.SliceToChan([]int{-10, 20, -4, -5, 1, 5, 8, 10, -20, 4})
	expected := helper.SliceToChan([]int{-10, -5, -5, -5, 1, -20, -20})

	movingMin := trend.NewMovingMinWithPeriod[int](4)
	actual := movingMin.Compute(input)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
