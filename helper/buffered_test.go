package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestBuffered(_ *testing.T) {
	c := make(chan int, 1)
	b := helper.Buffered(c, 4)

	c <- 1
	c <- 2
	c <- 3
	c <- 4

	close(c)

	helper.Drain(b)
}
