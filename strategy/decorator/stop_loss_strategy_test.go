package decorator_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/asset"
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/strategy"
	"github.com/hellomyheart/go-indicator/strategy/decorator"
	"github.com/hellomyheart/go-indicator/strategy/trend"
)

func TestStopLossStrategy(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	results, err := helper.ReadFromCsvFile[strategy.Result]("testdata/stop_loss_strategy.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	expected := helper.Map(results, func(r *strategy.Result) strategy.Action { return r.Action })

	innerStrategy := trend.NewAroonStrategy()
	stopLossStrategy := decorator.NewStopLossStrategy(innerStrategy, 0.02)

	actual := stopLossStrategy.Compute(snapshots)

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStopLossStrategyReport(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	innerStrategy := trend.NewAroonStrategy()
	stopLossStrategy := decorator.NewStopLossStrategy(innerStrategy, 0.02)

	report := stopLossStrategy.Report(snapshots)

	fileName := "stop_loss_strategy.html"
	defer helper.Remove(t, fileName)

	err = report.WriteToFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
}
