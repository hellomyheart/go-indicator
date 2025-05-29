package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestDrain(_ *testing.T) {
	input := helper.SliceToChan([]int{2, 4, 6, 8})
	helper.Drain(input)
}
