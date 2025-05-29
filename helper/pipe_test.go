package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestPipe(t *testing.T) {
	data := []int{2, 4, 6, 8}
	expected := helper.SliceToChan(data)

	input := helper.SliceToChan(data)
	actual := make(chan int)

	go helper.Pipe(input, actual)

	err := helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}
