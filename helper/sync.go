package helper

import "slices"

// CommonPeriod 计算所有数据通道可以同步的最小周期
// 		返回最大值
//
// 例如:
//
//	// Synchronize channels with periods 4, 2, and 3.
//	commonPeriod := helper.CommonPeriod(4, 2, 3) // commonPeriod = 4
//
//	// Synchronize the first channel
//	c1 := helper.Sync(commonPeriod, 4, c1)
//
//	// Synchronize the second channel
//	c2 := helper.Sync(commonPeriod, 2, c2)
//
//	// Synchronize the third channel
//	c3 := helper.Sync(commonPeriod, 3, c3)
func CommonPeriod(periods ...int) int {
	return slices.Max(periods)
}

// SyncPeriod 调整给定的通道以匹配给定的公共周期。
// 通过跳过部分元素实现周期对齐，
// 数据的频率是一致的，但是起始点不一致
// commonPeriod 目标周期
// 原始周期
// chan
func SyncPeriod[T any](commonPeriod, period int, c <-chan T) <-chan T {
	forwardPeriod := commonPeriod - period

	if forwardPeriod > 0 {
		// 跳过部分原始
		c = Skip(c, forwardPeriod)
	}

	return c
}
