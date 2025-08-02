package compound

import (
	"fmt"
	"math"

	"github.com/hellomyheart/go-indicator/asset"
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator"
	"github.com/hellomyheart/go-indicator/indicator/trend"
	"github.com/hellomyheart/go-indicator/indicator/volatility"
	"github.com/hellomyheart/go-indicator/strategy"
)

// 波动率趋势一号策略

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

	// 布林带宽度预警值 默认值
	DefaultBollingerBandWidthStart = 0.2
	// 布林带宽度安全值 默认值
	DefaultBollingerBandWidthEnd = 0.1

	// 布林带中轨斜率预警值 默认值
	DefaultBollingerBandWidthSlopeStart = 0.05
	// 布林带中轨斜率安全值 默认值
	DefaultBollingerBandWidthSlopeEnd = 0.01
	// 默认周期
	DefaultPeriod = 20
)

// 波动率趋势一号策略结构体
// macd 未来可以加
type VolatilityTrendOneStrategy struct {

	// BollingerBands表示计算布林带的配置参数。
	BollingerBands *volatility.BollingerBands[float64]
	// 布林带宽度
	BollingerBandWidth *volatility.BollingerBandWidth[float64]
	// sar
	Sar *trend.Sar[float64]
	// 周期
	Period int
	// 布林带宽度预警值
	BollingerBandWidthStart float64
	// 布林带宽度安全值
	BollingerBandWidthEnd float64
	// 布林带中轨斜率预警值
	BollingerBandWidthSlopeStart float64
	// 布林带中轨斜率安全值
	BollingerBandWidthSlopeEnd float64
}

// 获取策略 使用默认参数
func NewVolatilityTrendOneStrategy() *VolatilityTrendOneStrategy {
	return NewVolatilityTrendOneStrategyWith(DefaultPeriod, DefaultBollingerBandWidthStart, DefaultBollingerBandWidthEnd, DefaultBollingerBandWidthSlopeStart, DefaultBollingerBandWidthSlopeEnd)
}

// 获取策略
func NewVolatilityTrendOneStrategyWith(peroid int, BollingerBandWidthStart, BollingerBandWidthEnd, BollingerBandWidthSlopeStart, BollingerBandWidthSlopeEnd float64) *VolatilityTrendOneStrategy {

	return &VolatilityTrendOneStrategy{
		BollingerBands:               volatility.NewBollingerBandsWithPeriod[float64](peroid),
		BollingerBandWidth:           volatility.NewBollingerBandWidthWithPeriod[float64](peroid),
		Sar:                          trend.NewSar[float64](),
		Period:                       peroid,
		BollingerBandWidthStart:      BollingerBandWidthStart,
		BollingerBandWidthEnd:        BollingerBandWidthEnd,
		BollingerBandWidthSlopeStart: BollingerBandWidthSlopeStart,
		BollingerBandWidthSlopeEnd:   BollingerBandWidthSlopeEnd,
	}
}

func (v *VolatilityTrendOneStrategy) Name() string {
	return fmt.Sprintf("Volatility-Trend-One Strategy (%d, %.0f, %.0f, %.0f, %.0f)",
		v.Period,
		v.BollingerBandWidthStart,
		v.BollingerBandWidthEnd,
		v.BollingerBandWidthSlopeStart,
		v.BollingerBandWidthSlopeEnd,
	)

}

// Compute处理所提供的资产快照，并生成一系列可操作的建议。
func (m *VolatilityTrendOneStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan strategy.Action {

	// snapshots2[0 1 2]  计算sar

	// 将输入通道复制为3个通道（只要输入通道的结算价）
	// closings[0]  计算布林带
	// closings[1]  计算布林带宽度

	snapshots2 := helper.Duplicate(snapshots, 4)

	sarHighsShots := asset.SnapshotsAsHighs(snapshots2[0])
	sarLowShots := asset.SnapshotsAsLows(snapshots2[1])
	sarCloseShots := asset.SnapshotsAsClosings(snapshots2[2])

	closings := helper.Duplicate(
		asset.SnapshotsAsClosings(snapshots2[3]),
		2,
	)

	// 布林带指标计算
	upperChan, middleChan, lowerChan := m.BollingerBands.Compute(closings[0])
	// 布林带宽度计算
	widthChan := m.BollingerBandWidth.Compute(closings[1])
	// SAR计算
	TrendTypeChan, sarChan := m.Sar.Compute(sarHighsShots, sarLowShots, sarCloseShots)

	// sar 跳过初始数据 周期数 -1
	TrendTypeChan = helper.Skip(TrendTypeChan, m.IdlePeriod())
	sarChan = helper.Skip(sarChan, m.IdlePeriod())

	// rsiActions := strategy.DenormalizeActions(
	// 	m.RsiStrategy.Compute(snapshotsSplice[1]),
	// )

	// 上一个动作
	lastAction := strategy.Hold
	// 上一个中轨值
	lastMiddle := 0.0
	// 计算
	actions := helper.Operate6(upperChan, middleChan, lowerChan, widthChan, TrendTypeChan, sarChan, func(upper, middle, lower, width float64, TrendType indicator.TrendType, sar float64) strategy.Action {
		switch lastAction {
		case strategy.Hold:
			// 未触发动作，需要寻找入场时机
			// 1.判断布林带宽度
			// 	宽度小于等于预警值
			if width <= m.BollingerBandWidthStart {
				// 更新上一个中轨值
				lastMiddle = middle
				return strategy.Hold
			}
			// 2. 计算斜率
			// 斜率小于等于预警值
			tempSlope := (middle - lastMiddle) / lastMiddle

			if math.Abs(tempSlope) <= m.BollingerBandWidthSlopeStart {
				// 更新上一个中轨值
				lastMiddle = middle
				return strategy.Hold
			}
			// 更新 lastMiddle
			// 3. 判断斜率方向和 sar方向
			slopeTrendType := indicator.NewTrendType(tempSlope)
			if slopeTrendType != TrendType {
				// 更新上一个中轨值
				lastMiddle = middle
				return strategy.Hold
			}
			// 4.符合条件
			// 判断中轨方向
			switch slopeTrendType {
			case indicator.FALLING:
				// 下降趋势
				// 卖
				lastAction = strategy.Sell
				// 更新上一个中轨值
				lastMiddle = middle
				return strategy.Sell
			case indicator.STABLE:
				// 平衡趋势
				// 更新上一个中轨值
				lastMiddle = middle
				return strategy.Hold
			case indicator.RISING:
				// 上升趋势  indicator.RISING:
				// 更新上一个中轨值
				lastMiddle = middle
				lastAction = strategy.Buy
				return strategy.Buy
			default:
				// 更新上一个中轨值
				lastMiddle = middle
				return strategy.Hold
			}
		case strategy.Buy:
			// 已经买入
			// 维持 或者卖出
			// 1. 判断宽度是否小于安全值
			if width <= m.BollingerBandWidthEnd {
				// 更新上一个中轨值
				lastMiddle = middle
				lastAction = strategy.Hold
				return strategy.Sell
			}
			// 2.计算斜率
			tempSlope := (middle - lastMiddle) / lastMiddle

			// 斜率小于等于预警值
			if math.Abs(tempSlope) <= m.BollingerBandWidthSlopeEnd {
				// 更新上一个中轨值
				lastMiddle = middle
				lastAction = strategy.Hold
				return strategy.Sell
			}
			// 3. 判断斜率方向和sar方向
			slopeTrendType := indicator.NewTrendType(tempSlope)
			if slopeTrendType != TrendType {
				// 更新上一个中轨值
				lastMiddle = middle
				lastAction = strategy.Hold
				return strategy.Sell
			}
			// 4.符合条件
			// 判断中轨方向
			switch slopeTrendType {
			case indicator.FALLING:
				//  上升变下降趋势
				// 更新上一个中轨值
				lastMiddle = middle
				lastAction = strategy.Hold
				return strategy.Sell
			case indicator.STABLE:
				// 上升变平衡趋势
				// 更新上一个中轨值
				lastMiddle = middle
				lastAction = strategy.Hold
				return strategy.Sell
			case indicator.RISING:
				// 上升趋势  indicator.RISING:
				// 更新上一个中轨值
				lastMiddle = middle
				lastAction = strategy.Buy
				return strategy.Buy
			default:
				// 更新上一个中轨值
				lastMiddle = middle
				lastAction = strategy.Hold
				return strategy.Sell
			}
		case strategy.Sell:
			// 已经卖出
			// 维持 或者买入
			// 1. 判断宽度是否小于安全值
			if width <= m.BollingerBandWidthEnd {
				// 更新上一个中轨值
				lastMiddle = middle
				lastAction = strategy.Hold
				return strategy.Buy
			}
			// 2.计算斜率
			tempSlope := (middle - lastMiddle) / lastMiddle

			// 斜率小于等于预警值
			if math.Abs(tempSlope) <= m.BollingerBandWidthSlopeEnd {
				// 更新上一个中轨值
				lastMiddle = middle
				lastAction = strategy.Hold
				return strategy.Buy
			}
			// 3. 判断斜率方向和sar方向
			slopeTrendType := indicator.NewTrendType(tempSlope)
			if slopeTrendType != TrendType {
				// 更新上一个中轨值
				lastMiddle = middle
				lastAction = strategy.Hold
				return strategy.Buy
			}
			// 4.符合条件
			// 判断中轨方向
			switch slopeTrendType {
			case indicator.RISING:
				// 下降变上升
				// 更新上一个中轨值
				lastMiddle = middle
				lastAction = strategy.Hold
				return strategy.Buy
			case indicator.STABLE:
				// 上升变平衡趋势
				// 更新上一个中轨值
				lastMiddle = middle
				lastAction = strategy.Hold
				return strategy.Buy
			case indicator.FALLING:
				// 下降趋势
				// 更新上一个中轨值
				lastMiddle = middle
				lastAction = strategy.Sell
				return strategy.Sell
			default:
				// 更新上一个中轨值
				lastMiddle = middle
				lastAction = strategy.Hold
				return strategy.Buy
			}
		default:
			return strategy.Hold
		}
	})
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
