package trend_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/asset"
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/strategy"
	"github.com/hellomyheart/go-indicator/strategy/trend"
)

func TestEnvelopeStrategy(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	results, err := helper.ReadFromCsvFile[strategy.Result]("testdata/envelope_strategy.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	expected := helper.Map(results, func(r *strategy.Result) strategy.Action { return r.Action })

	envelope := trend.NewEnvelopeStrategy()
	actual := envelope.Compute(snapshots)

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEnvelopeStrategyReport(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	envelope := trend.NewEnvelopeStrategy()
	report := envelope.Report(snapshots)

	fileName := "envelope_strategy.html"
	defer helper.Remove(t, fileName)

	err = report.WriteToFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
}
