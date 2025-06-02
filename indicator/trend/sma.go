package trend

import (
	"fmt"

	"github.com/hellomyheart/go-indicator/helper"
)

const (
	// DefaultSmaPeriod 默认的SMA周期 50
	DefaultSmaPeriod = 50
)

// Sma 表示计算简单移动平均线的参数。
//
// Example:
//
//	sma := trend.NewSma[float64]()
//	sma.Period = 10
//
//	result := sma.Compute(c)
type Sma[T helper.Number] struct {
	// Period is the time period for the SMA.
	Period int
}

// NewSma 使用默认参数初始化一个新的SMA实例。
func NewSma[T helper.Number]() *Sma[T] {
	return NewSmaWithPeriod[T](DefaultSmaPeriod)
}

// NewSmaWithPeriod 使用自定义参数初始化一个新的SMA实例。
func NewSmaWithPeriod[T helper.Number](period int) *Sma[T] {
	return &Sma[T]{
		Period: period,
	}
}

// Compute sma计算方法
func (s *Sma[T]) Compute(c <-chan T) <-chan T {
	// 移动和
	sum := NewMovingSum[T]()
	// 设置移动和周期为sma周期
	sum.Period = s.Period

	// sma =  移动和/周期
	return helper.Apply(sum.Compute(c), func(sum T) T {
		return sum / T(s.Period)
	})
}

// IdlePeriod 是SMA不会产生任何结果的初始阶段。（周期值-1）
func (s *Sma[T]) IdlePeriod() int {
	return s.Period - 1
}

// String 是SMA的字符串表示形式。
func (s *Sma[T]) String() string {
	return fmt.Sprintf("SMA(%d)", s.Period)
}
