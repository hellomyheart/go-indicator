package volume_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/asset"
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/strategy"
	"github.com/hellomyheart/go-indicator/strategy/volume"
)

func TestEaseOfMovementStrategy(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	results, err := helper.ReadFromCsvFile[strategy.Result]("testdata/ease_of_movement_strategy.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	expected := helper.Map(results, func(r *strategy.Result) strategy.Action { return r.Action })

	emvs := volume.NewEaseOfMovementStrategy()
	actual := emvs.Compute(snapshots)

	err = helper.CheckEquals(actual, expected)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEaseOfMovementStrategyReport(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/brk-b.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	emvs := volume.NewEaseOfMovementStrategy()
	report := emvs.Report(snapshots)

	fileName := "ease_of_movement_strategy.html"
	defer helper.Remove(t, fileName)

	err = report.WriteToFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
}
