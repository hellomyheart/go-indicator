package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestGcdWithTwoValues(t *testing.T) {
	actual := helper.Gcd(1220, 512)
	expected := 4

	if actual != expected {
		t.Fatalf("actual %d expected %d", actual, expected)
	}
}

func TestGcdWithThreeValues(t *testing.T) {
	actual := helper.Gcd(1, 2, 5)
	expected := 1

	if actual != expected {
		t.Fatalf("actual %d expected %d", actual, expected)
	}
}

func TestGcdWithFourValues(t *testing.T) {
	actual := helper.Gcd(2, 4, 6, 12)
	expected := 2

	if actual != expected {
		t.Fatalf("actual %d expected %d", actual, expected)
	}
}
