package backtest

import (
	"github.com/hellomyheart/go-indicator/asset"
	"github.com/hellomyheart/go-indicator/strategy"
)

// Report is the backtest report interface.
type Report interface {
	// Begin is called when the backtest begins.
	Begin(assetNames []string, strategies []strategy.Strategy) error

	// AssetBegin is called when backtesting for the given asset begins.
	AssetBegin(name string, strategies []strategy.Strategy) error

	// Write writes the given strategy actions and outomes to the report.
	Write(assetName string, currentStrategy strategy.Strategy, snapshots <-chan *asset.Snapshot, actions <-chan strategy.Action, outcomes <-chan float64) error

	// AssetEnd is called when backtesting for the given asset ends.
	AssetEnd(name string) error

	// End is called when the backtest ends.
	End() error
}
