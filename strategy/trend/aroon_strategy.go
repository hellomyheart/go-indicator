package trend

import (
	"github.com/hellomyheart/go-indicator/asset"
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/indicator/trend"
	"github.com/hellomyheart/go-indicator/strategy"
)

// AroonStrategy represents the configuration parameters for calculating the Aroon strategy.
// Aroon is a technical analysis tool that gauges trend direction and strength in asset
// prices. It comprises two lines: Aroon Up and Aroon Down. Aroon Up measures uptrend
// strength, while Aroon Down measures downtrend strength. When Aroon Up exceeds
// Aroon Down, it suggests a bullish trend; when Aroon Down surpasses Aroon Up,
// it indicates a bearish trend.
type AroonStrategy struct {
	// Aroon represent the configuration for calculating the Aroon indicator.
	Aroon *trend.Aroon[float64]
}

// NewAroonStrategy function initializes a new Aroon strategy instance
// with the default parameters.
func NewAroonStrategy() *AroonStrategy {
	return &AroonStrategy{
		Aroon: trend.NewAroon[float64](),
	}
}

// Name returns the name of the strategy.
func (*AroonStrategy) Name() string {
	return "Aroon Strategy"
}

// Compute processes the provided asset snapshots and generates a
// stream of actionable recommendations.
func (a *AroonStrategy) Compute(c <-chan *asset.Snapshot) <-chan strategy.Action {
	snapshots := helper.Duplicate(c, 2)

	highs := asset.SnapshotsAsHighs(snapshots[0])
	lows := asset.SnapshotsAsLows(snapshots[1])

	ups, downs := a.Aroon.Compute(highs, lows)

	actions := helper.Operate(ups, downs, func(up, down float64) strategy.Action {
		if up > down {
			return strategy.Buy
		}

		if down > up {
			return strategy.Sell
		}

		return strategy.Hold
	})

	// Aroon starts only after a full period.
	actions = helper.Shift(actions, a.Aroon.Period-1, strategy.Hold)

	return actions
}

// Report processes the provided asset snapshots and generates a
// report annotated with the recommended actions.
func (a *AroonStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	//
	// snapshots[0] -> dates
	// snapshots[1] -> highs    |> ups, downs
	// snapshots[2] -> lows     |
	// snapshots[3] -> closings
	// snapshots[4] -> Compute -> actions  -> annotations
	//                            outcomes
	//
	snapshots := helper.Duplicate(c, 5)

	dates := asset.SnapshotsAsDates(snapshots[0])
	highs := asset.SnapshotsAsHighs(snapshots[1])
	lows := asset.SnapshotsAsLows(snapshots[2])
	closings := asset.SnapshotsAsClosings(snapshots[3])

	ups, downs := a.Aroon.Compute(highs, lows)
	ups = helper.Shift(ups, a.Aroon.Period-1, 0)
	downs = helper.Shift(downs, a.Aroon.Period-1, 0)

	actions, outcomes := strategy.ComputeWithOutcome(a, snapshots[4])
	annotations := strategy.ActionsToAnnotations(actions)
	outcomes = helper.MultiplyBy(outcomes, 100)

	report := helper.NewReport(a.Name(), dates)
	report.AddChart()
	report.AddChart()

	report.AddColumn(helper.NewNumericReportColumn("Close", closings))
	report.AddColumn(helper.NewNumericReportColumn("Aroon Up", ups), 1)
	report.AddColumn(helper.NewNumericReportColumn("Aroon Down", downs), 1)
	report.AddColumn(helper.NewAnnotationReportColumn(annotations), 0, 1)

	report.AddColumn(helper.NewNumericReportColumn("Outcome", outcomes), 2)

	return report
}
