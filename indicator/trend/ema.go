package trend

import (
	"fmt"

	"github.com/hellomyheart/go-indicator/helper"
)

const (
	// DefaultEmaPeriod 默认的均线周期为20。
	DefaultEmaPeriod = 20

	// DefaultEmaSmoothing 默认均线平滑为2。
	DefaultEmaSmoothing = 2
)

// Ema 加权指数移动平均线的结构体
//
// 例子:
//
//	ema := trend.NewEma[float64]()
//	ema.Period = 10
//
//	result := ema.Compute(c)
type Ema[T helper.Number] struct {
	// 周期
	Period int

	// 平滑系数
	Smoothing T
}

// NewEma 用默认参数初始化一个新的EMA实例。
func NewEma[T helper.Number]() *Ema[T] {
	return &Ema[T]{
		Period:    DefaultEmaPeriod,
		Smoothing: DefaultEmaSmoothing,
	}
}

// NewEmaWithPeriod 函数用给定的周期初始化新的EMA实例。
func NewEmaWithPeriod[T helper.Number](period int) *Ema[T] {
	ema := NewEma[T]()
	ema.Period = period

	return ema
}

// Compute 函数接受一个数字通道并计算指定时间段内的EMA。
func (e *Ema[T]) Compute(c <-chan T) <-chan T {
	// 创建一个结果chan
	result := make(chan T, cap(c))

	// 启动一个协程
	go func() {
		// 最后关闭协程
		defer close(result)

		// 初始化一个简单移动平均线
		sma := NewSma[T]()
		// 简单移动平均线的周期是 ema的周期
		sma.Period = e.Period

		// 取前周期数个元素 ，并计算简单移动平均线
		before := <-sma.Compute(helper.Head(c, e.Period))
		// 将结果返回
		result <- before

		// 权重乘数 = 平滑系数 / 周期+ 1
		// 最新价格 权重乘数的 权重，
		// 上一个EMA价格权重是 （1- 权重乘数）
		multiplier := e.Smoothing / T(e.Period+1)

		for n := range c {
			// (当前值 - 上一个值) *  权重乘数 + 上一个值
			before = (n-before)*multiplier + before
			// 将结果返回
			result <- before
		}
	}()

	// 返回chan
	return result
}

// IdlePeriod 是EMA产生任何结果的初始阶段。 返回周期-1
func (e *Ema[T]) IdlePeriod() int {
	return e.Period - 1
}

// String 是EMA的字符串表示形式。
func (e *Ema[T]) String() string {
	return fmt.Sprintf("EMA(%d)", e.Period)
}
