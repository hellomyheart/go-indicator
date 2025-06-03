package volatility

import (
	"math"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

// DefaultAdxPeriod 默认ADX周期参数
const (
	DefaultAdxPeriod = 14
)

// Adx 表示计算平均趋向指数(Average Directional Index)的配置参数
// ADX值越高表明趋势越强，通常高于25视为明显趋势
// 计算公式:
// TR = max(high-low, |high-close_prev|, |low-close_prev|)
// +DI = 100 * EMA(+DM)/EMA(TR)
// -DI = 100 * EMA(-DM)/EMA(TR)
// DX = 100 * |+DI - -DI|/(+DI + -DI)
// ADX = EMA(DX)

// EMA(TR) 使用atr
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

	// TR 有结果的地方是1 + 一个周期 -1
	// highs[0]、lows[0] 消费了 1
	atr := NewAtrWithMa(trend.NewEmaWithPeriod[T](adx.Period))
	atrChan := atr.Compute(highs[0], lows[0], closes[0])

	// 复制两个TR
	atrsChan := helper.Duplicate(atrChan, 2)

	// EMA(+DM) 有结果的地方是1 + 一个周期 -1
	// highs[1] 消费了1
	admPlus := NewAdmWithMa(trend.NewEmaWithPeriod[T](adx.Period), true)
	admPlusChan := admPlus.Compute(highs[1], closes[1])

	// EMA(-DM) 有结果的地方是1+ 一个周期 -1
	// lows[1] 消费了1
	admMinus := NewAdmWithMa(trend.NewEmaWithPeriod[T](adx.Period), false)
	admMinusChan := admMinus.Compute(closes[2], lows[1])

	// +DI = 100 * EMA(+DM)/EMA(TR)
	// 有结果的地方是 周期-1
	diPlusChan := helper.MultiplyBy(helper.Divide(admPlusChan,
		atrsChan[0]), 100)
	diPlusChan = helper.Buffered(diPlusChan, adx.Period)

	// -DI = 100 * EMA(-DM)/EMA(TR)
	// 有结果的地方是 周期-1
	diMinusChan := helper.MultiplyBy(helper.Divide(admMinusChan,
		atrsChan[1]), 100)

	diMinusChan = helper.Buffered(diMinusChan, adx.Period)

	// EMA
	ema := trend.NewEmaWithPeriod[T](adx.Period)

	// ADX = EMA(DX)
	// DX =100 * |+DI - -DI|/(+DI + -DI)
	return ema.Compute(helper.Operate(diPlusChan, diMinusChan, func(diPlus, diMinus T) T {
		return T(100 * math.Abs(float64(diPlus-diMinus)) / (float64(diPlus + diMinus)))
	}))
}

// IdlePeriod 2 * 周期-1
func (a *Adx[T]) IdlePeriod() int {
	return 2*a.Period - 1
}
