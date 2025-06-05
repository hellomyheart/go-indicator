package indicator_test

import (
	"fmt"
	"testing"

	"github.com/hellomyheart/go-indicator/indicator"
)

func TestTrendType(t *testing.T) {
	var c indicator.TrendType = -1
	fmt.Println("TrendType:", c.String())    // 输出：TrendType: FALLING
	fmt.Println("TrendType value:", c.Int()) // 输出：TrendType value: -1
	fmt.Print(1)
}
