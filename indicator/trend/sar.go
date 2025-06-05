package trend

import (
	"math"

	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator"
)

// SarParameters SAR指标参数配置
// sar 是迭代计算指标，为了保证每次计算结果一致，每次都要从第一个数据开始计算
// （或者每次计算的起始数据都一致，不然不保证每次的计算结果一致）

// 抛物线SAR。它是一个流行的技术指标，用于识别趋势和作为跟踪止损。
// PSAR = PSAR[i - 1] - ((PSAR[i - 1] - EP) * AF)
// 如果趋势是下跌：
//  - PSAR取PSAR与前两个高点的最大值
//  - 如果当前高点大于等于PSAR，则使用EP
// 如果趋势是上升：
//  - PSAR取PSAR与前两个低点的最小值
//  - 如果当前低点小于等于PSAR，则使用EP
// 如果PSAR大于收盘价，则趋势为下跌，EP取EP与低点的最小值
// 如果PSAR小于等于收盘价，则趋势为上升，EP取EP与高点的最大值
// 如果趋势保持不变且AF小于0.20，则递增0.02
// 如果趋势发生变化，则将AF重置为0.02
// 注：EP代表Extreme Point（极端点），AF代表Acceleration Factor（加速因子），这些都是技术分析领域的专业术语。

const (

	// DefaultSarAcceleration 默认初始加速因子
	DefaultSarAcceleration = 0.02

	// DefaultSarMaxAcceleration 默认最大加速度因子
	DefaultSarMaxAcceleration = 0.2
)

// sar参数
type Sar[T helper.Number] struct {
	// 初始加速度因子 (默认 0.02)
	Acceleration float64

	// 最大加速度因子 (默认 0.2)
	MaxAcceleration float64
}

// sar无参构造函数
func NewSar[T helper.Number]() *Sar[T] {
	return NewSarWithParams[T](
		DefaultSarAcceleration,
		DefaultSarMaxAcceleration,
	)
}

// sar有参构造函数
func NewSarWithParams[T helper.Number](acceleration, maxAcceleration float64) *Sar[T] {
	return &Sar[T]{
		Acceleration:    acceleration,
		MaxAcceleration: maxAcceleration,
	}
}

// // Sar 实现抛物线转向指标
// type Sar[T helper.Number] struct {
// 	params SarParameters[T]

// 	// 内部状态缓存
// 	highs []T
// 	lows  []T
// 	sars  []T
// 	ep    T       // 极值点
// 	af    float64 // 当前加速度
// 	trend bool    // 当前趋势 true=上升 false=下降
// 	first bool    // 初始状态标志
// }

// Compute 实现SAR指标计算逻辑
func (s *Sar[T]) Compute(high, low, cl <-chan T) (<-chan indicator.TrendType, <-chan T) {
	// Sar值Chan
	results := make(chan T, 2)

	// Sar方向Chan
	trends := make(chan indicator.TrendType, 2)

	// 默认下降趋势
	trends <- indicator.FALLING
	// 上一个趋势
	lastTrend := indicator.FALLING

	// sar i -1
	tempS := <-high
	lastSar := float64(tempS)
	// 上一个high
	lastH := lastSar
	// h -2
	var lastH2 float64
	// sar 默认值为h[0]
	results <- tempS

	// 舍弃 clo[0]
	helper.Skip(cl, 1)

	// 设置af 为 step (默认0.02)
	af := s.Acceleration

	// 设置ep 为 low[0]
	ep := float64(<-low)

	// lastL
	lastL := ep
	// l -2
	var lastL2 float64

	// 输入：high low clo
	// af ep

	// 第一次循环
	isFirst := true

	// 当前H
	var tempH float64
	// 当前L
	var tempL float64
	// 当前C
	var tempC float64
	// 当前Trend
	var tempTrend indicator.TrendType

	// 三个chan 合并为一个
	tempChan := make(chan T, 3)
	go func() {
		defer close(tempChan)
		for {
			// 依次从三个通道读取
			h, ok := <-high
			if !ok {
				return
			}
			tempChan <- h

			l, ok := <-low
			if !ok {
				return
			}
			tempChan <- l

			c, ok := <-cl
			if !ok {
				return
			}
			tempChan <- c
		}
	}()

	go func() {
		defer close(trends)
		defer close(results)
		for {
			value, ok := <-tempChan
			if !ok {
				// tempChan 已关闭，处理退出或错误逻辑
				return
			}
			tempH = float64(value)

			value, ok = <-tempChan
			if !ok {
				// tempChan 已关闭，处理退出或错误逻辑
				return
			}
			tempL = float64(value)

			value, ok = <-tempChan
			if !ok {
				// tempChan 已关闭，处理退出或错误逻辑
				return
			}
			tempC = float64(value)

			// sar[i] = sar[i-1] -(sar[i-1] - ep) * af
			tempSar := lastSar - (lastSar-ep)*af
			if lastTrend == indicator.FALLING {
				//如果上一个趋势是 下降
				// sar[i] = max(sar[i], highs[i - 1])
				tempSar = math.Max(tempSar, lastH)
				// 如果不是第一次循环
				if !isFirst {
					tempSar = math.Max(tempSar, lastH2)
				} else {
					// 第一次循环 false
					isFirst = false
				}
				// 如果当前highs[i] 大于等于 sar[i]
				if tempH >= tempSar {
					// sar[i] = ep
					tempSar = ep
				}
			} else {
				// 上一个趋势是上升
				// sar[i] =  min(sar[i], lows[i -1])
				tempSar = math.Min(tempSar, lastL)
				if !isFirst {
					// 如果不是第一次循环
					// sar[i] = min(sar[i], lows[i - 2])
					tempSar = math.Min(tempSar, lastL2)
				} else {
					// 第一次循环 true
					isFirst = false
				}
				// 如果当前lows[i] 小于等于 sar[i]
				if tempL <= tempSar {
					// sar[i] = ep
					tempSar = ep
				}

			}
			// 上一个ep
			lastEp := ep

			// 如果sar[i] > close[i]
			if tempSar > tempC {
				// 当前趋势是下降趋势
				tempTrend = indicator.FALLING
				// ep = min(ep, lows[i])
				ep = math.Min(ep, tempL)
			} else {
				// 当前趋势是上升趋势
				tempTrend = indicator.RISING
				// ep = max(ep, highs[i])
				ep = math.Max(ep, tempH)
			}

			// 如果趋势发生了反转
			if tempTrend != lastTrend {
				// af = step 重置af
				af = s.Acceleration
			} else if lastEp != ep && af < s.MaxAcceleration {
				// 如果趋势没有发生发转 并且 上一个ep != 当前ep 并且af < max (默认max = 0.2)
				// af = af + step
				af += s.Acceleration
			}

			// 更新临时值
			lastSar = tempSar
			lastTrend = tempTrend
			lastH2 = lastH
			lastH = tempH

			lastL2 = lastL
			lastL = tempL

			// 发送结果
			trends <- tempTrend
			results <- T(tempSar)
		}
	}()

	// 输出：trends results

	return trends, results

}

// IdlePeriod sar是0
func (m *Sar[T]) IdlePeriod() int {
	return 0
}

// 参考了 indicatorts 代码的实现
// // Copyright (c) 2022 Onur Cinar. All Rights Reserved.
// // https://github.com/cinar/indicatorts

// import { checkSameLength } from '../../helper/numArray';
// import { Trend } from '../types';

// /**
//  * Parabolic SAR result object.
//  */
// export interface PSARResult {
//   trends: Trend[];
//   psarResult: number[];
// }

// /**
//  * Optional configuration of parabolic SAR parameters.
//  */
// export interface PSARConfig {
//   step?: number;
//   max?: number;
// }

// /**
//  * The default configuration of parabolic SAR.
//  */
// export const PSARDefaultConfig: Required<PSARConfig> = {
//   step: 0.02,
//   max: 0.2,
// };

// /**
//  * 抛物线SAR。它是一个流行的技术指标，用于识别趋势和作为跟踪止损。
//  * PSAR = PSAR[i - 1] - ((PSAR[i - 1] - EP) * AF)
//  *
//  * If the trend is Falling:
//  *  - PSAR is the maximum of PSAR or the previous two high values.
//  *  - If the current high is greather than or equals to PSAR, use EP.
//  *
//  * If the trend is Rising:
//  *  - PSAR is the minimum of PSAR or the previous two low values.
//  *  - If the current low is less than or equals to PSAR, use EP.
//  *
//  * If PSAR is greater than the closing, trend is falling, and the EP
//  * is set to the minimum of EP or the low.
//  *
//  * If PSAR is lower than or equals to the closing, trend is rising, and the EP
//  * is set to the maximum of EP or the high.
//  *
//  * If the trend is the same, and AF is less than 0.20, increment it by 0.02.
//  * If the trend is not the same, set AF to 0.02.
//  *
//  * Based on video https://www.youtube.com/watch?v=MuEpGBAH7pw&t=0s.
//  *
//  * @param highs high values.
//  * @param lows low values.
//  * @param closings closing values.
//  * @param config configuration.
//  * @return psar result.
//  */
// export function psar(
//   highs: number[],
//   lows: number[],
//   closings: number[],
//   config: PSARConfig = {}
// ): PSARResult {
//   checkSameLength(highs, lows, closings);

//   const { step, max } = {
//     ...PSARDefaultConfig,
//     ...config,
//   };
//   // 返回两个数组
//   const trends = new Array<Trend>(highs.length);
//   const psarResult = new Array<number>(highs.length);

//   // 默认是下降趋势
//   trends[0] = Trend.FALLING;
//   // 默认sar[0]值为highs[0]
//   psarResult[0] = highs[0];

//   // 设置af 为 step (默认0.02)
//   let af = step;
//   // 设置ep为 low[0]
//   let ep = lows[0];

//   // 此时close[0] 还没有遍历
//   // 代码里遍历close是从1开始， 0应该是舍弃了

//   for (let i = 1; i < psarResult.length; i++) {
//     // sar[i] = sar[i-1] -(sar[i-1] - ep) * af
//     psarResult[i] = psarResult[i - 1] - (psarResult[i - 1] - ep) * af;
//     //如果上一个趋势是 下降
//     if (trends[i - 1] === Trend.FALLING) {
//       // sar[i] = max(sar[i], highs[i - 1])
//       psarResult[i] = Math.max(psarResult[i], highs[i - 1]);
//       // 如果不是第一次循环(第一次循环没有上上个highs)
//       if (i > 1) {
//         // sar[i] = max(sar[i], highs[i - 2])
//         psarResult[i] = Math.max(psarResult[i], highs[i - 2]);
//       }
//       // 如果当前highs[i] 大于等于 sar[i]
//       if (highs[i] >= psarResult[i]) {
//         // sar[i] = ep
//         psarResult[i] = ep;
//       }
//     } else {
//       // 上个趋势是上升
//       // sar[i] =  min(sar[i], lows[i -1])
//       psarResult[i] = Math.min(psarResult[i], lows[i - 1]);
//       if (i > 1) {
//         // 如果不是第一次循环
//         // sar[i] = min(sar[i], lows[i - 2])
//         psarResult[i] = Math.min(psarResult[i], lows[i - 2]);
//       }
//       // 如果当前lows[i] 小于等于 sar[i]
//       if (lows[i] <= psarResult[i]) {
//         // sar[i] = ep
//         psarResult[i] = ep;
//       }
//     }

//     // 上一个ep
//     const prevEp = ep;

//     // 如果sar[i] > close[i]
//     if (psarResult[i] > closings[i]) {
//       // 当前趋势是下降趋势
//       trends[i] = Trend.FALLING;
//       // ep = min(ep, lows[i])
//       ep = Math.min(ep, lows[i]);
//     } else {
//       // 当前趋势是上升趋势
//       trends[i] = Trend.RISING;
//       // ep = max(ep, highs[i])
//       ep = Math.max(ep, highs[i]);
//     }

//     // 如果趋势发生了反转
//     if (trends[i] !== trends[i - 1]) {
//       // af = step 重置af
//       af = step;
//     } else if (prevEp !== ep && af < max) {
//       // 如果趋势没有发生发转 并且 上一个ep != 当前ep 并且af < max (默认max = 0.2)
//       // af = af + step
//       af += step;
//     }
//   }

//   // 返回结果
//   return {
//     trends,
//     psarResult,
//   };
// }

// // Export full name
// export { psar as parabolicSAR };
