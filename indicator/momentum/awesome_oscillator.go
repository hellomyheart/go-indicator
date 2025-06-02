package momentum

import (
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
)

// 绝密振荡器 AO
const (
	// DefaultAwesomeOscillatorShortPeriod 是AO的默认短周期。
	DefaultAwesomeOscillatorShortPeriod = 5

	// DefaultAwesomeOscillatorLongPeriod 是AO的默认长周期。
	DefaultAwesomeOscillatorLongPeriod = 34
)

// AwesomeOscillator 表示计算AO （Awesome Oscillator）的配置参数。它通过比较短期价格走势（5个周期平均值）和长期趋势（34个周期平均值）来衡量市场势头。它在零线附近的价值反映了上方的看涨和下方的看跌。越过零线可能预示着潜在的趋势逆转。交易者使用AO来确认现有趋势，确定进入/退出点，并了解动量变化。
//
//	中位价格 = ((Low + High) / 2).
//	AO = 5周期均值 - 34周期均值
//
// Example:
//
//	ao := momentum.AwesomeOscillator[float64]()
//	values := ao.Compute(lows, highs)
type AwesomeOscillator[T helper.Number] struct {
	// 短周期sma
	ShortSma *trend.Sma[T]

	// 长周期sma
	LongSma *trend.Sma[T]
}

// NewAwesomeOscillator 函数初始化一个新的Awesome振荡器实例。
func NewAwesomeOscillator[T helper.Number]() *AwesomeOscillator[T] {
	return &AwesomeOscillator[T]{
		ShortSma: trend.NewSmaWithPeriod[T](DefaultAwesomeOscillatorShortPeriod),
		LongSma:  trend.NewSmaWithPeriod[T](DefaultAwesomeOscillatorLongPeriod),
	}
}

// Compute 函数接受2个数字通道并计算AwesomeOscillator。
func (a *AwesomeOscillator[T]) Compute(highs, lows <-chan T) <-chan T {
	// 将一个chan复制为多个chan （2）
	// 复制了两个平均价格
	medianSplice := helper.Duplicate(
		// 除法 (high s + lows)/2 求出平均价格
		helper.DivideBy(
			helper.Add(highs, lows),
			2,
		),
		2,
	)

	// 计算平均价格的 短均线
	shortSma := a.ShortSma.Compute(medianSplice[0])
	// 计算平均价格的 长均线
	longSma := a.LongSma.Compute(medianSplice[1])

	// 跳过短均线的前（长周期-1 -短周期-1）周期个数据
	// 实现了数据对齐
	// 例如：
	// 短周期5 长周期34 33 -4 = 29
	// 长周期在34才有数据，短周期在5
	// 短周期跳过29个，数据就实现了对齐
	shortSma = helper.Skip(shortSma, a.LongSma.IdlePeriod()-a.ShortSma.IdlePeriod())

	// 返回 短均线减去长均线 的结果
	// 这里的均线都是最高价格和最低价格的平均值的均线
	return helper.Subtract(
		shortSma,
		longSma,
	)
}

// IdlePeriod 是绝妙振荡器不会产生任何结果的初始周期。
func (a *AwesomeOscillator[T]) IdlePeriod() int {
	// 长周期sma的IdlePeriod
	return a.LongSma.IdlePeriod()
}
