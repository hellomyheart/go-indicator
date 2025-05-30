package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestRoundDigit(t *testing.T) {
	input := 10.1234
	expected := 10.12

	actual := helper.RoundDigit(input, 2)

	if actual != expected {
		t.Fatalf("actual %v expected %v", actual, expected)
	}
}
