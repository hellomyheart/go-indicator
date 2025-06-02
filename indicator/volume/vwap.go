package volume

import (
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

const (
	// DefaultVwapPeriod 为VWAP的默认周期。
	DefaultVwapPeriod = 14
)

// Vwap保存用于计算成交量加权平均价格（Vwap）的配置参数。它提供了资产交易的平均价格。
//
//	VWAP = Sum(Closing * Volume) / Sum(Volume)
//
// Example:
//
//	vwap := volume.NewVwap[float64]()
//	result := vwap.Compute(closings, volumes)
type Vwap[T helper.Number] struct {
	// Sum is the Moving Sum instance.
	Sum *trend.MovingSum[T]
}

// NewVwap 函数使用默认参数初始化新的VWAP实例。
func NewVwap[T helper.Number]() *Vwap[T] {
	return NewVwapWithPeriod[T](DefaultVwapPeriod)
}

// NewVwapWithPeriod 函数用给定的周期初始化新的VWAP实例。
func NewVwapWithPeriod[T helper.Number](period int) *Vwap[T] {
	return &Vwap[T]{
		Sum: trend.NewMovingSumWithPeriod[T](period),
	}
}

// Compute 函数接受一个数字通道并计算VWAP。
func (v *Vwap[T]) Compute(closings, volumes <-chan T) <-chan T {
	// 复制一个chan
	volumesSplice := helper.Duplicate(volumes, 2)

	// 除法
	return helper.Divide(
		// sum(close * volumes)
		v.Sum.Compute(
			helper.Multiply(
				closings,
				volumesSplice[0],
			),
		),
		// sum(volumes)
		v.Sum.Compute(
			volumesSplice[1],
		),
	)
}

// IdlePeriod 是VWAP不会产生任何结果的初始阶段。 周期数-1
func (v *Vwap[T]) IdlePeriod() int {
	return v.Sum.IdlePeriod()
}
