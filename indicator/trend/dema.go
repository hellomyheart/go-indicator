package trend

import "github.com/hellomyheart/go-indicator/helper"

// Dema 表示计算双重指数移动平均线（DEMA）的参数。当周期为5天的DEMA高于周期为35天的DEMA时，出现看涨交叉。当35天周期的DEMA高于5天周期的DEMA时，就会出现看跌交叉。
//
//	DEMA = (2 * EMA1(values)) - EMA2(EMA1(values))
//
// Example:
//
//	dema := trend.NewDema[float64]()
//	dema.Ema1.Period = 10
//	dema.Ema2.Period = 16
//
//	result := dema.Compute(input)
type Dema[T helper.Number] struct {
	// Ema1 对原始数据进行一次指数移动平均（Exponential Moving Average）。
	Ema1 *Ema[T]

	// Ema2 对 EMA1 的结果再次进行指数移动平均。
	Ema2 *Ema[T]
}

// NewDema function initializes a new DEMA instance
// with the default parameters.
func NewDema[T helper.Number]() *Dema[T] {
	return &Dema[T]{
		Ema1: NewEma[T](),
		Ema2: NewEma[T](),
	}
}

// Compute function takes a channel of numbers and computes the DEMA
// over the specified period.
func (d *Dema[T]) Compute(c <-chan T) <-chan T {
	// 计算一下Ema 一个周期 -1 后会有第一个数据 20是第一个数据（假如一个周期是20）
	ema1 := helper.Duplicate(d.Ema1.Compute(c), 2)
	// 二次EMA(对EMA进行EMA) 2个周期-3后会有第一个数据 39是第一个数据
	ema2 := d.Ema2.Compute(ema1[1])
	// 将 EMA1 的结果乘以 2， 一个周期 -1 后会有第一个数据
	doubleEma1 := helper.MultiplyBy(ema1[0], 2)
	//EMA1比EMA2 快 一个周期 - 1
	// 快的那个， 缓存慢的那个的差
	doubleEma1 = helper.Buffered(doubleEma1, d.Ema2.Period)

	//再减去 EMA2 的结果，从而消除传统 EMA 的滞后性。
	return helper.Subtract(doubleEma1, ema2)
}

// IdlePeriod is the initial period that DEMA won't yield any results.
func (d *Dema[T]) IdlePeriod() int {
	return d.Ema1.Period + d.Ema2.Period - 2
}
