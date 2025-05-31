package strategy_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/asset"
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/strategy"
	"github.com/hellomyheart/go-indicator/strategy/trend"
)

func TestMajorityStrategy(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/repository/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	results, err := helper.ReadFromCsvFile[strategy.Result]("testdata/majority.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	expected := helper.Map(results, func(r *strategy.Result) strategy.Action { return r.Action })

	majority := strategy.NewMajorityStrategy("Majority Strategy")
	majority.Strategies = append(majority.Strategies, strategy.NewBuyAndHoldStrategy(), trend.NewMacdStrategy())

	actual := majority.Compute(snapshots)

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMajorityStrategyReport(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/repository/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	majority := strategy.NewMajorityStrategy("Majority Strategy")
	majority.Strategies = append(majority.Strategies, strategy.NewBuyAndHoldStrategy(), trend.NewMacdStrategy())

	report := majority.Report(snapshots)

	fileName := "majority.html"
	defer helper.Remove(t, fileName)

	err = report.WriteToFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
}
