package trend_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/asset"
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/strategy"
	"github.com/hellomyheart/go-indicator/strategy/trend"
)

func TestKdjStrategy(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	results, err := helper.ReadFromCsvFile[strategy.Result]("testdata/kdj_strategy.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	expected := helper.Map(results, func(r *strategy.Result) strategy.Action { return r.Action })

	kdj := trend.NewKdjStrategy()
	actual := kdj.Compute(snapshots)

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}

func TestKdjStrategyReport(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	kdj := trend.NewKdjStrategy()

	report := kdj.Report(snapshots)

	fileName := "kdj_strategy.html"
	defer helper.Remove(t, fileName)

	err = report.WriteToFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
}
