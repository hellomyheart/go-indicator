package trend

import "github.com/hellomyheart/go-indicator/helper"

const (
	// DefaultVwmaPeriod is the default period for the VWMA.
	DefaultVwmaPeriod = 20
)

// Vwma 表示计算成交量加权移动平均线（VWMA）的配置参数。它对价格数据进行平均，强调成交量，这意味着成交量越大的区域将具有更大的权重。
//
//	VWMA = Sum(价格 * 成交量) / Sum(成交量)
type Vwma[T helper.Number] struct {
	// 周期
	Period int
}

// NewVwma 使用默认参数初始化新的VWMA实例。
func NewVwma[T helper.Number]() *Vwma[T] {
	return &Vwma[T]{
		Period: DefaultVwmaPeriod,
	}
}

// Compute 函数接受一个数字通道并计算VWMA和信号线。
// closing 价格chan
// volume 成交量chan
func (v *Vwma[T]) Compute(closing, volume <-chan T) <-chan T {
	// 复制成交量chan
	volumes := helper.Duplicate(volume, 2)

	// 创建移动和
	sum := NewMovingSum[T]()
	// 设置移动和周期
	sum.Period = v.Period

	// 除法运算
	return helper.Divide(
		// 移动和
		// 价格 * 成交量
		sum.Compute(
			helper.Multiply(closing, volumes[0]),
		),
		// 成交量
		sum.Compute(volumes[1]),
	)
}

// IdlePeriod 是VWMA不会产生任何结果的初始阶段。 周期-1
func (v *Vwma[T]) IdlePeriod() int {
	return v.Period - 1
}
