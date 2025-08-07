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
	// 读取数据
	snapshot, err := helper.ReadFromCsvFile[asset.Snapshot]("testdata/lc2511-60.csv", true)
	if err != nil {
		t.Fatal(err)
	}

	// 创建策略
	volatilityTrendOne := compound.NewVolatilityTrendOneStrategy()

	// 源数据复制3份
	snapshots := helper.Duplicate(snapshot, 3)

	action := helper.Duplicate(volatilityTrendOne.Compute(snapshots[0]), 2)

	asset := asset.SnapshotsAsClosings(snapshots[1])

	outcomes := strategy.Outcome(asset, action[1])

	// for o := range outcomes {
	// 	fmt.Println(o)
	// }

	for i := 0; i < 3000000; i++ {
		a, ok := <-action[0]
		o, ok := <-outcomes
		if !ok {
			return
		}
		v := <-snapshots[2]
		if a.Annotation() == "" {
			continue
		}
		fmt.Print(changeAction(a).Annotation(), "  ")
		fmt.Print(v.Close, "   ")
		fmt.Println(v.Date)
		fmt.Println(o)
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
