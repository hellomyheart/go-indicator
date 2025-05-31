package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestSeq(t *testing.T) {
	expected := helper.SliceToChan([]int{2, 3, 4, 5})
	actual := helper.Seq(2, 6, 1)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
