package volatility

import (
	"github.com/hellomyheart/go-indicator/asset"
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/volatility"
	"github.com/hellomyheart/go-indicator/strategy"
)

// BollingerBandsStrategy表示计算布林带策略的配置参数。
// 收盘价高于上带表明买入信号，而低于下带
// 表示卖出信号。
type BollingerBandsStrategy struct {
	// BollingerBands表示计算布林带的配置参数。
	BollingerBands *volatility.BollingerBands[float64]
}

// NewBollingerBandsStrategy 函数 初始化一个新的布林带策略实例。 使用默认周期（20）
func NewBollingerBandsStrategy() *BollingerBandsStrategy {
	return &BollingerBandsStrategy{
		BollingerBands: volatility.NewBollingerBands[float64](),
	}
}

// Name 返回策略的名称。
func (*BollingerBandsStrategy) Name() string {
	return "Bollinger Bands Strategy"
}

// Compute处理所提供的资产快照，并生成一系列可操作的建议。
func (b *BollingerBandsStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan strategy.Action {
	// 将输入通道复制为2个通道（只要输入通道的结算价）
	closings := helper.Duplicate(
		asset.SnapshotsAsClosings(snapshots),
		2,
	)

	// 第一个输入通道计算布林带指标，返回三个轨的输出通道
	uppers, middles, lowers := b.BollingerBands.Compute(closings[0])
	// 创建一个协程 排空中轨通道
	go helper.Drain(middles)

	// 第二个输入通道（复制过来的）， 跳过一个周期-1的数据
	closings[1] = helper.Skip(closings[1], b.BollingerBands.IdlePeriod())

	// 计算动作
	// 输入 上轨 下轨 输入通道（结算价）
	actions := helper.Operate3(uppers, lowers, closings[1], func(upper, lower, closing float64) strategy.Action {
		// 如果结算价大于上轨
		if closing > upper {
			// 返回 买信号
			return strategy.Buy
		}

		// 如果结算价小于下轨
		if lower > closing {
			// 返回 卖信号
			return strategy.Sell
		}

		// 否则返回 持有信号(不操作)
		return strategy.Hold
	})

	// 布林带只在一个完整的周期后开始。
	// 为了保证输入通道和输出通道数据量相同，（1:1）
	// 输出通道右移数据填充之前空缺的 1个周期-1
	actions = helper.Shift(actions, b.BollingerBands.IdlePeriod(), strategy.Hold)

	// 返回输出通道
	return actions
}

// Report处理提供的资产快照，并生成带有建议操作注释的报告。
func (b *BollingerBandsStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	//
	// snapshots[0] -> dates 				   		// 时间数据
	// snapshots[1] -> closings[0] -> closings 		// 收盘价
	//                 closings[1] -> upper    		// 布林带上轨
	//                             -> middle  		// 布林带中轨
	//                             -> lower			// 布林带下轨
	// snapshots[2] -> actions     -> annotations   // 建议操作
	//              -> outcomes						//
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
