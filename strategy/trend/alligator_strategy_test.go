package trend_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/asset"
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/strategy"
	"github.com/hellomyheart/go-indicator/strategy/trend"
)

func TestAlligatorStrategy(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	results, err := helper.ReadFromCsvFile[strategy.Result]("testdata/alligator_strategy.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	expected := helper.Map(results, func(r *strategy.Result) strategy.Action { return r.Action })

	alligator := trend.NewAlligatorStrategy()
	actual := alligator.Compute(snapshots)

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAlligatorStrategyReport(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	alligator := trend.NewAlligatorStrategy()
	report := alligator.Report(snapshots)

	fileName := "alligator_strategy.html"
	defer helper.Remove(t, fileName)

	err = report.WriteToFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
}
