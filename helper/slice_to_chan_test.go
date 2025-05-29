package helper_test

import (
	"reflect"
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestSliceToChan(t *testing.T) {
	expected := []int{2, 4, 6, 8}
	actual := helper.ChanToSlice(helper.SliceToChan(expected))

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("actual %v expected %v", actual, expected)
	}
}
