package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestDuplicate(t *testing.T) {
	expecteds := []float64{-10, 20, -4, -5}

	outputs := helper.Duplicate[float64](helper.SliceToChan(expecteds), 4)

	for i, expected := range expecteds {
		for _, output := range outputs {
			actual := <-output
			if actual != expected {
				t.Fatalf("index %d actual %v expected %v", i, actual, expected)
			}
		}
	}
}
