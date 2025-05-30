package helper_test

import (
	"strings"
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestJSONToChan(t *testing.T) {
	expected := helper.SliceToChan([]int{2, 4, 6, 8})
	input := "[2, 4, 6, 8]"

	actual := helper.JSONToChan[int](strings.NewReader(input))

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
