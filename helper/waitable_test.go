package helper_test

import (
	"sync"
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestWaitable(_ *testing.T) {
	wg := &sync.WaitGroup{}
	c := make(chan int)

	helper.Waitable[int](wg, c)
	close(c)

	wg.Wait()
}
