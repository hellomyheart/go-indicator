package strategy

import "github.com/hellomyheart/go-indicator/helper"

// Outcome基于提供的值模拟执行给定操作的潜在结果。
// 这个是理想模拟，没有延迟的计算？
func Outcome[T helper.Number](values <-chan T, actions <-chan Action) <-chan float64 {
	// 资金
	balance := 1.0
	// 持仓数
	shares := 0.0

	return helper.Operate(values, actions, func(value T, action Action) float64 {
		// 如果资金 > 0 && action == Buy
		if balance > 0 && action == Buy {
			// 持仓数 = 资金 /value
			shares = balance / float64(value)
			// 资金 = 0
			balance = 0
			// 持仓数 > 0 && action == Sell
		} else if shares > 0 && action == Sell {
			// 资金 = 持仓数 * value
			balance = shares * float64(value)
			// 持仓数 = 0
			shares = 0
		}
		// 返回预计收益 = 资金 + 持仓 * 价值 - 1.0
		return balance + (shares * float64(value)) - 1.0
	})
}
