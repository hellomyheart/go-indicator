package helper_test

import (
	"testing"
	"time"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestDaysBetween(t *testing.T) {
	from := time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2024, 9, 15, 0, 0, 0, 0, time.UTC)

	actual := helper.DaysBetween(from, from)
	expected := 0

	if actual != expected {
		t.Fatalf("actual %d expected %d", actual, expected)
	}

	actual = helper.DaysBetween(from, to)
	expected = 14

	if actual != expected {
		t.Fatalf("actual %d expected %d", actual, expected)
	}
}
