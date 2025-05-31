package strategy_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/asset"
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/strategy"
	"github.com/hellomyheart/go-indicator/strategy/trend"
)

func TestSplitStrategy(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/repository/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	results, err := helper.ReadFromCsvFile[strategy.Result]("testdata/split.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	expected := helper.Map(results, func(r *strategy.Result) strategy.Action { return r.Action })

	split := strategy.NewSplitStrategy(
		trend.NewMacdStrategy(),
		trend.NewApoStrategy(),
	)

	actual := split.Compute(snapshots)

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSplitStrategyReport(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/repository/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	split := strategy.NewSplitStrategy(
		trend.NewMacdStrategy(),
		trend.NewApoStrategy(),
	)

	report := split.Report(snapshots)

	fileName := "split.html"
	defer helper.Remove(t, fileName)

	err = report.WriteToFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAllSplitStrategies(t *testing.T) {
	strategies := []strategy.Strategy{
		strategy.NewBuyAndHoldStrategy(),
		strategy.NewMajorityStrategyWith("", []strategy.Strategy{
			strategy.NewBuyAndHoldStrategy(),
		}),
	}

	allSplitStrategies := strategy.AllSplitStrategies(strategies)

	expected := len(strategies)*len(strategies) - len(strategies)
	actual := len(allSplitStrategies)

	if actual != expected {
		t.Fatalf("actual=%d expected=%d", actual, expected)
	}
}
