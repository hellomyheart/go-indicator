package trend_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

func TestMovingMax(t *testing.T) {
	input := helper.SliceToChan([]int{-10, 20, -4, -5, 1, 5, 8, 10, -20, 4})
	expected := helper.SliceToChan([]int{20, 20, 5, 8, 10, 10, 10})

	movingMax := trend.NewMovingMaxWithPeriod[int](4)
	actual := movingMax.Compute(input)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
