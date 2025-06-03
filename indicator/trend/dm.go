package trend

import (
	"math"

	"github.com/hellomyheart/go-indicator/helper"
)

// dm 计算dm, 辅助计算adx
//

type Dm[T helper.Number] struct {
	DmType bool // false. -dm true. +dm
}

// NewTr 函数使用默认参数初始化新的TR实例。
func NewDm[T helper.Number](dmType bool) *Dm[T] {
	return &Dm[T]{DmType: dmType}
}

// Compute 函数接受一个数字通道，并计算指定时间段内的TR。

func (a *Dm[T]) Compute(f, l <-chan T) <-chan T {
	// 跳过f一个
	if a.DmType {
		// + dm
		f = helper.Skip(f, 1)
	} else {
		// -dm
		l = helper.Skip(l, 1)
	}
	dm := helper.Operate(f, l, func(a, b T) T {
		return T(math.Max(0, float64(a-b)))
	})
	return dm
}

// IdlePeriod is 不会产生任何结果的过程
// Dm是 1
func (a *Dm[T]) IdlePeriod() int {
	return 1
}
