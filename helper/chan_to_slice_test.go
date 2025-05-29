package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestChanToSlice(t *testing.T) {
	input := []int{2, 4, 6, 8}
	expected := helper.SliceToChan(input)

	actual := make(chan int, len(input))
	for _, n := range input {
		actual <- n
	}
	close(actual)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
