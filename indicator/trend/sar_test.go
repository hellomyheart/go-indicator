package trend_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

func TestSar(t *testing.T) {

	// 定义并初始化三个切片（slice）存储价格数据
	highs := []float64{
		3836.86, 3766.57, 3576.17, 3513.55, 3529.75, 3756.17, 3717.17, 3572.62,
		3612.43,
	}

	lows := []float64{
		3643.25, 3542.73, 3371.75, 3334.02, 3314.75, 3558.21, 3517.79, 3447.9,
		3494.39,
	}

	closings := []float64{
		3790.55, 3546.2, 3507.31, 3340.81, 3529.6, 3717.41, 3544.35, 3478.14,
		3612.08,
	}

	h := helper.SliceToChan(highs)
	l := helper.SliceToChan(lows)
	c := helper.SliceToChan(closings)

	sarV := []float64{
		3836.86, 3836.86, 3836.86, 3808.95, 3770.96, 3314.75, 3314.75, 3323.58,
		3332.23,
	}

	sv := helper.SliceToChan(sarV)

	sarT := []indicator.TrendType{
		indicator.FALLING,
		indicator.FALLING,
		indicator.FALLING,
		indicator.FALLING,
		indicator.FALLING,
		indicator.RISING,
		indicator.RISING,
		indicator.RISING,
		indicator.RISING,
	}
	st := make(chan indicator.TrendType)
	go func() {
		defer close(st)
		for _, n := range sarT {
			st <- n
		}
	}()

	sar := trend.NewSar[float64]()

	r1, r2 := sar.Compute(h, l, c)
	r22 := helper.RoundDigits(r2, 2)

	for r222 := range r22 {

		tempT := <-st
		tempV := <-sv

		r11 := <-r1

		if tempV != r222 {
			t.Fatalf("actual %v expected %v", r222, tempV)
		}

		if r11 != tempT {
			t.Fatalf("actual %v expected %v", r11, tempT)
		}
	}
}
