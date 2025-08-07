package compound

import (
	"fmt"

	"github.com/hellomyheart/go-indicator/asset"
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator"
	"github.com/hellomyheart/go-indicator/indicator/trend"
	"github.com/hellomyheart/go-indicator/indicator/volatility"
	"github.com/hellomyheart/go-indicator/strategy"
	strend "github.com/hellomyheart/go-indicator/strategy/trend"
)

// 波动率趋势一号策略
// ADX 判断是否是趋势， 趋势入场
// 斜率判断方向， 这两个都很好

// 止损有问题，斜率判断，延迟很大

// sar 的缺点， 假的方向，
// 除了趋势，还有震荡，在震荡中，sar的方向判断延迟很大，效果非常差

// ADX 判断是否是趋势
// 布林带 主要使用中轨， 判断斜率

// 布林带宽度入场
//  斜率方向和sar方向相同，二次入场判断
// sar不适合用来做分钟级别的止损（速度太慢）

// 后面还要写 前置策略( 策略执行前有一个前置条件判断)、 延时等待策略（策略信号连续发出多少次后才OK）

// 后置检查策略， 用来在策略发送信号前做一个检查（满足条件才发送）

//  有向无环图、前置后置，只是逻辑，实际计算最好是并行， 这样速度最快

// 分策略， 多个策略一起运行，不要一个大全的策略， 激活一个策略后，其他策略哪些允许激活，哪些不允许，需要思考。
// 策略对某一个品种的胜率， 盈亏比， 这些要记录

// 布林带 布林带宽 sar 三个指标
// bollinger_bands.go   布林带指标   bollinger_bands_strategy.go 布林带策略
// bollinger_band_width.go
// sar.go

// 1.判断布林带带宽发生大变化   预警值

// 2.判断sar方向和 布林带中轨斜率方向

// 1. 一致，趋势确认 当前布林带带宽 大于预警值
// 		按照方向操作

// 2. 不一致， 等待一致

// 3. 布林带小于预警值 sar方向和斜率方向不一致

// 3.等待中轨斜率不再大于指定值（可以设置）或 sar翻转（不利于方向）
// 结束

// 需要的参数 	布林带带宽 带宽预警值  带宽安全值
//  		   布林带中轨 斜率预警值  斜率安全值

const (
	// // 布林带宽度预警值 默认值
	// DefaultBollingerBandWidthStart = 0.002

	// // 默认斜率预警值
	// DefaultSlopeStart = 0.0002
	// 默认周期 20
	DefaultPeriod = 30
)

// 波动率趋势一号策略结构体
// macd 未来可以加
type VolatilityTrendOneStrategy struct {

	// // 布林带宽度
	// BollingerBandWidth *volatility.BollingerBandWidth[float64]

	// // 布林带宽度预警值
	// BollingerBandWidthStart float64
	// // 斜率预警值
	// SlopeStart float64

	// // 周期
	Period int
	// Adx 用来判断是否是趋势
	Adx *volatility.Adx[float64]

	// // BollingerBands表示计算布林带的配置参数。
	BollingerBands *volatility.BollingerBands[float64]

	// sar
	Sar *trend.Sar[float64]
	// macd策略
	MacdStrategy *strend.MacdStrategy
}

// 获取策略 使用默认参数
func NewVolatilityTrendOneStrategy() *VolatilityTrendOneStrategy {
	return NewVolatilityTrendOneStrategyWith(DefaultPeriod)
}

// 获取策略
func NewVolatilityTrendOneStrategyWith(peroid int) *VolatilityTrendOneStrategy {

	return &VolatilityTrendOneStrategy{

		// BollingerBandWidth:      volatility.NewBollingerBandWidthWithPeriod[float64](peroid),

		// BollingerBandWidthStart: BollingerBandWidthStart,
		// SlopeStart:              SlopeStart,
		Period:         peroid,
		Adx:            volatility.NewAdxWithPeriod[float64](peroid),
		BollingerBands: volatility.NewBollingerBandsWithPeriod[float64](peroid),
		Sar:            trend.NewSar[float64](),
		MacdStrategy:   strend.NewMacdStrategy(),
	}
}

func (v *VolatilityTrendOneStrategy) Name() string {
	return fmt.Sprintf("Volatility-Trend-One Strategy (%d)",
		v.Period,
	)
}

// Compute处理所提供的资产快照，并生成一系列可操作的建议。
func (m *VolatilityTrendOneStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan strategy.Action {

	// snapshots2[0 1 2]  计算sar

	// 将输入通道复制为3个通道（只要输入通道的结算价）
	// closings[0]  计算布林带
	// closings[1]  计算布林带宽度

	snapshots2 := helper.Duplicate(snapshots, 5)

	closings := asset.SnapshotsAsClosings(snapshots2[0])

	highsShots := helper.Duplicate(asset.SnapshotsAsHighs(snapshots2[1]), 2)
	lowShots := helper.Duplicate(asset.SnapshotsAsLows(snapshots2[2]), 2)
	closeShots := helper.Duplicate(asset.SnapshotsAsClosings(snapshots2[3]), 2)
	// 布林带指标计算
	upperChan, middleChan, lowerChan := m.BollingerBands.Compute(closings)
	// adx 计算
	adxChan := m.Adx.Compute(highsShots[0], lowShots[0], closeShots[0])

	// sar 计算
	trendTypeChan, sarChan := m.Sar.Compute(highsShots[1], lowShots[1], closeShots[1])

	trendTypeChan = helper.Skip(trendTypeChan, m.BollingerBands.IdlePeriod())
	sarChan = helper.Skip(sarChan, m.BollingerBands.IdlePeriod())
	// adx 的 周期会长
	diff := m.Adx.IdlePeriod() - m.BollingerBands.IdlePeriod()

	adxChan = helper.Shift(adxChan, diff, 0.0)

	// macd 计算
	macdActionChan := m.MacdStrategy.Compute(snapshots2[4])

	// 上一个动作
	lastAction := strategy.Hold
	// 上一个中轨值
	lastMiddle := 0.0
	// 上一个斜率 方向
	lastTrendType := indicator.STABLE
	// 计算
	// 布林带 3个 + adx = 4个

	actions := helper.Operate7(upperChan, middleChan, lowerChan, adxChan, sarChan, trendTypeChan, macdActionChan, func(upper, middle, lower, adx, sar float64, trendType indicator.TrendType, macdAction strategy.Action) strategy.Action {
		// 提前计算斜率
		tempSlope := 0.0
		if 0 != lastMiddle {
			tempSlope = (middle - lastMiddle) / lastMiddle
		}
		slopeTrendType := indicator.NewTrendType(tempSlope)
		switch lastAction {
		case strategy.Hold:
			// 未触发动作，需要寻找入场时机
			// 判断adx 大小
			if adx <= 25 {
				// 更新上一个中轨值
				lastMiddle = middle
				lastTrendType = slopeTrendType
				return strategy.Hold
			}
			// 斜率方向和sar 方向相同
			if slopeTrendType != trendType {
				lastMiddle = middle
				lastTrendType = slopeTrendType
				return strategy.Hold
			}

			if slopeTrendType == indicator.STABLE {
				lastMiddle = middle
				lastTrendType = slopeTrendType
				return strategy.Hold
			}
			// macd 方向和斜率方向不同
			if int(macdAction) != int(slopeTrendType) {
				lastMiddle = middle
				lastTrendType = slopeTrendType
				return strategy.Hold
			}

			if slopeTrendType == indicator.FALLING {
				// 下降趋势
				// 卖
				lastAction = strategy.Sell
				// 更新上一个中轨值
				lastMiddle = middle
				lastTrendType = slopeTrendType
				return strategy.Sell
			} else {
				// 上升趋势
				// 买
				lastAction = strategy.Buy
				// 更新上一个中轨值
				lastMiddle = middle
				lastTrendType = slopeTrendType
				return strategy.Buy
			}
		case strategy.Buy:
			// 已经买入
			// 判断sar方向是否已经改变
			if trendType == lastTrendType {
				// 没有改变
				lastAction = strategy.Buy
				// 更新上一个中轨值
				lastMiddle = middle
				lastTrendType = trendType
				return strategy.Hold
			} else {
				lastAction = strategy.Hold
				// 更新上一个中轨值
				lastMiddle = middle
				lastTrendType = trendType
				return strategy.Sell
			}
		case strategy.Sell:
			// 已经买入
			// 判断sar方向是否已经改变
			if trendType == lastTrendType {
				// 没有改变
				lastAction = strategy.Sell
				// 更新上一个中轨值
				lastMiddle = middle
				lastTrendType = trendType
				return strategy.Hold
			} else {
				lastAction = strategy.Hold
				// 更新上一个中轨值
				lastMiddle = middle
				lastTrendType = trendType
				return strategy.Buy
			}

		default:
			return strategy.Hold
		}
	})
	actions = helper.Shift(actions, m.IdlePeriod(), strategy.Hold)
	return actions
}

// Report处理提供的资产快照，并生成带有建议操作注释的报告。
func (b *VolatilityTrendOneStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	//
	// snapshots[0] -> dates 				   		// 时间数据
	// snapshots[1] -> closings[0] -> closings 		// 收盘价
	//                 closings[1] -> upper    		// 布林带上轨
	//                             -> middle  		// 布林带中轨
	//                             -> lower			// 布林带下轨
	// snapshots[2] -> actions     -> annotations   // 建议操作
	//              -> outcomes						// 营收百分百
	// 复制三个输入通道
	snapshots := helper.Duplicate(c, 3)

	// 时间
	dates := asset.SnapshotsAsDates(snapshots[0])
	// 收盘价 复制两个
	closings := helper.Duplicate(asset.SnapshotsAsClosings(snapshots[1]), 2)

	// 第一个收盘价用来计算布林带
	uppers, middles, lowers := b.BollingerBands.Compute(closings[0])
	// 布林带周期前补0 上轨 中轨下轨
	uppers = helper.Shift(uppers, b.BollingerBands.IdlePeriod(), 0)
	middles = helper.Shift(middles, b.BollingerBands.IdlePeriod(), 0)
	lowers = helper.Shift(lowers, b.BollingerBands.IdlePeriod(), 0)

	// 获取计算动作 和预计收益通道
	actions, outcomes := strategy.ComputeWithOutcome(b, snapshots[2])
	// ActionsToAnnotations接受一个动作建议通道，并返回一个新通道包含这些操作的相应注释。
	// 返回 S 或者B 字符串的通道
	annotations := strategy.ActionsToAnnotations(actions)
	// 乘以100 ，就是营收百分比
	outcomes = helper.MultiplyBy(outcomes, 100)

	// 创建一个报表 使用策略的名称和日期
	report := helper.NewReport(b.Name(), dates)
	// 报表添加一个图表
	report.AddChart()

	// 添加列 收盘价
	report.AddColumn(helper.NewNumericReportColumn("Close", closings[1]))
	// 添加列 上轨
	report.AddColumn(helper.NewNumericReportColumn("Upper", uppers))
	// 添加列 中轨
	report.AddColumn(helper.NewNumericReportColumn("Middle", middles))
	// 添加列 下轨
	report.AddColumn(helper.NewNumericReportColumn("Lower", lowers))
	// 添加列 动作
	report.AddColumn(helper.NewAnnotationReportColumn(annotations))
	// 添加列 收益百分比 新的图表1，其他列都在0号图表？ （不确定）
	report.AddColumn(helper.NewNumericReportColumn("Outcome", outcomes), 1)

	// 返回报表
	return report
}

// IdlePeriod 是不会产生任何结果的初始阶段。
// 周期数 - 1，因为在计算时需要至少一个完整的周期来计算中轨和标准差。
func (m *VolatilityTrendOneStrategy) IdlePeriod() int {
	return m.Period - 1
}
