package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestLcmWithTwoValues(t *testing.T) {
	actual := helper.Lcm(18, 32)
	expected := 288

	if actual != expected {
		t.Fatalf("actual %d expected %d", actual, expected)
	}
}

func TestLcmWithFourValues(t *testing.T) {
	actual := helper.Lcm(1, 2, 8, 6)
	expected := 24

	if actual != expected {
		t.Fatalf("actual %d expected %d", actual, expected)
	}
}

func TestLcmWithFiveValues(t *testing.T) {
	actual := helper.Lcm(2, 7, 3, 9, 8)
	expected := 504

	if actual != expected {
		t.Fatalf("actual %d expected %d", actual, expected)
	}
}
