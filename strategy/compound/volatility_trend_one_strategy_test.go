package compound_test

import (
	"fmt"
	"testing"

	"github.com/hellomyheart/go-indicator/asset"
	"github.com/hellomyheart/go-indicator/helper"
	"github.com/hellomyheart/go-indicator/strategy"
	"github.com/hellomyheart/go-indicator/strategy/compound"
)

func TestVolatilityTrendOneStrategy(t *testing.T) {
	// 获取测试数据
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/SA2509.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	// 获取预期结果
	// results, err := helper.ReadFromCsvFile[strategy.Result]("testdata/macd_rsi_strategy.csv", true)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// expected := helper.Map(results, func(r *strategy.Result) strategy.Action { return r.Action })

	volatilityTrend := compound.NewVolatilityTrendOneStrategy()
	actual := volatilityTrend.Compute(snapshots)

	actual = helper.Shift(actual, volatilityTrend.IdlePeriod(), strategy.Hold)

	for item := range actual {
		fmt.Println(item.Annotation())
	}

	// 检查结果
	// err = helper.CheckEquals(actual, expected)
	// if err != nil {
	// 	t.Fatal(err)
	// }
}

func TestVolatilityTrendOneStrategyReport(t *testing.T) {
	snapshots, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/AG2510-5.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	volatilityTrendOne := compound.NewVolatilityTrendOneStrategy()

	report := volatilityTrendOne.Report(snapshots)

	fileName := "volatility_trend_one_strategy.html"
	defer helper.Remove(t, fileName)

	err = report.WriteToFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVolatilityTrendOneStrategyOutComes(t *testing.T) {
	snapshot, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/AG2510-5.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	volatilityTrendOne := compound.NewVolatilityTrendOneStrategy()

	snapshots := helper.Duplicate(snapshot, 2)

	action := volatilityTrendOne.Compute(snapshots[0])

	asset := snapshots[1]

	for i := 0; i < 3000; i++ {
		a, ok := <-action
		if !ok {
			return
		}
		v := <-asset
		fmt.Print(changeAction(a).Annotation(), "  ")
		fmt.Print(v.Close, "   ")
		fmt.Println(v.Date)
	}
}

// 反向
func changeAction(action strategy.Action) strategy.Action {
	// switch action {
	// case strategy.Buy:
	// 	return strategy.Sell

	// case strategy.Sell:
	// 	return strategy.Buy

	// default:
	// 	return strategy.Hold
	// }
	return action
}
