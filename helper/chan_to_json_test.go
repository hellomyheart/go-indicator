package helper_test

import (
	"bytes"
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestChanToJSON(t *testing.T) {
	input := helper.SliceToChan([]int{2, 4, 6, 8})
	expected := "[2,4,6,8]"

	var buffer bytes.Buffer

	err := helper.ChanToJSON(input, &buffer)
	if err != nil {
		t.Fatal(err)
	}

	actual := buffer.String()
	if actual != expected {
		t.Fatalf("actual=%s expected=%s", actual, expected)
	}
}
