package trend

import (
	"math"

	"github.com/hellomyheart/go-indicator/helper"
)

// DefaultAdxPeriod 默认ADX周期参数
const DefaultAdxPeriod = 14

// Adx 表示计算平均趋向指数(Average Directional Index)的配置参数
// ADX值越高表明趋势越强，通常高于25视为明显趋势
// 计算公式:
// TR = max(high-low, |high-close_prev|, |low-close_prev|)
// +DI = 100 * EMA(+DM)/EMA(TR)
// -DI = 100 * EMA(-DM)/EMA(TR)
// DX = 100 * |+DI - -DI|/(+DI + -DI)
// ADX = EMA(DX)
type Adx[T helper.Number] struct {
	Period int // 计算周期，默认14
}

// NewAdx 创建具有默认参数的新ADX实例
func NewAdx[T helper.Number]() *Adx[T] {
	return &Adx[T]{
		Period: DefaultAdxPeriod,
	}
}

// NewAdx 创建指定参数的新ADX实例
func NewAdxWithPeriod[T helper.Number](period int) *Adx[T] {
	return &Adx[T]{
		Period: period,
	}
}

// Compute 计算ADX指标值，需要传入最高价、最低价、收盘价三个通道
func (adx *Adx[T]) Compute(high, low, close <-chan T) <-chan T {

	// 复制chan TR +DM -DM
	highs := helper.Duplicate(high, 2)
	lows := helper.Duplicate(low, 2)
	closes := helper.Duplicate(close, 3)

	// TR
	tr := NewTr[T]()
	trChan := tr.Compute(highs[0], lows[0], closes[0])
	// 复制两个TR
	trsChan := helper.Duplicate(trChan, 2)

	// +DM
	highs[1] = helper.Skip(highs[1], 1)
	dmPlusChan := helper.Operate(highs[1], closes[1], func(high, close T) T {
		return T(math.Max(0, float64(high-close)))
	})

	// -DM
	lows[1] = helper.Skip(lows[1], 1)
	dmMinusChan := helper.Operate(lows[1], closes[2], func(low, close T) T {
		return T(math.Max(0, float64(close-low)))
	})

	// +DI = 100 * EMA(+DM)/EMA(TR)
	diPlusEma1 := NewEmaWithPeriod[T](adx.Period)
	diPlusEma2 := NewEmaWithPeriod[T](adx.Period)
	diPlusChan := helper.Divide(helper.MultiplyBy(diPlusEma1.Compute(dmPlusChan), 100),
		diPlusEma2.Compute(trsChan[0]))

	// -DI = 100 * EMA(-DM)/EMA(TR)
	diMinusEma1 := NewEmaWithPeriod[T](adx.Period)
	diMinusEma2 := NewEmaWithPeriod[T](adx.Period)
	diMinusChan := helper.Divide(helper.MultiplyBy(diMinusEma1.Compute(dmMinusChan), 100),
		diMinusEma2.Compute(trsChan[1]))

	// EMA
	ema := NewEmaWithPeriod[T](adx.Period)

	// ADX = EMA(DX)
	// DX =100 * |+DI - -DI|/(+DI + -DI)
	return ema.Compute(helper.Operate(diPlusChan, diMinusChan, func(diPlus, diMinus T) T {
		return T(100 * math.Abs(float64(diPlus-diMinus)) / (float64(diPlus + diMinus)))
	}))
}

// IdlePeriod 周期-1
func (a *Adx[T]) IdlePeriod() int {
	return a.Period - 1
}
