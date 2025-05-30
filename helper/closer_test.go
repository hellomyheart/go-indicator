package helper_test

import (
	"os"
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestCloseAndLogError(t *testing.T) {
	file, err := os.CreateTemp(os.TempDir(), "closer")
	if err != nil {
		t.Fatal(err)
	}

	helper.CloseAndLogError(file, "")
	helper.CloseAndLogError(file, "")
}
