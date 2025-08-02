// Package strategy contains the strategy functions.
//
// This package belongs to the Indicator project. Indicator is
// a Golang module that supplies a variety of technical
// indicators, strategies, and a backtesting framework
// for analysis.
//
// # License
//
//	Copyright (c) 2021-2024 Onur Cinar.
//	The source code is provided under GNU AGPLv3 License.
//	https://github.com/cinar/indicator
//
// # Disclaimer
//
// The information provided on this project is strictly for
// informational purposes and is not to be construed as
// advice or solicitation to buy or sell any security.
package strategy

import (
	"github.com/hellomyheart/go-indicator/asset"
	"github.com/hellomyheart/go-indicator/helper"
)

// Strategy defines a shared interface for trading strategies.
type Strategy interface {
	// Name returns the name of the strategy.
	Name() string

	// Compute processes the provided asset snapshots and generates a
	// stream of actionable recommendations.
	Compute(snapshots <-chan *asset.Snapshot) <-chan Action

	// Report processes the provided asset snapshots and generates a
	// report annotated with the recommended actions.
	Report(snapshots <-chan *asset.Snapshot) *helper.Report
}

// ComputeWithOutcome使用给定的策略来处理提供的资产快照，并生成一系列可操作的建议和结果。
// 返回动作 和
func ComputeWithOutcome(s Strategy, c <-chan *asset.Snapshot) (<-chan Action, <-chan float64) {
	// 输入通道复制为2个
	snapshots := helper.Duplicate(c, 2)

	// 使用第一个输入通道获取计算动作 
	// 并把计算动作复制为2个通道
	actions := helper.Duplicate(s.Compute(snapshots[0]), 2)
	// 使用第二个输入通道获取收盘价
	closings := asset.SnapshotsAsClosings(snapshots[1])

	// 使用收盘价和计算动作 获取预计收益通道
	outcomes := Outcome(closings, actions[1])
	// 返回计算动作和预计收益
	return actions[0], outcomes
}

// AllStrategies returns a slice containing references to all available base strategies.
func AllStrategies() []Strategy {
	return []Strategy{
		NewBuyAndHoldStrategy(),
	}
}

// ActionSources creates a slice of action channels, one for each strategy, where each channel emits actions
// computed by its corresponding strategy based on snapshots from the provided snapshot channel.
func ActionSources(strategies []Strategy, snapshots <-chan *asset.Snapshot) []<-chan Action {
	snapshotsSplice := helper.Duplicate(snapshots, len(strategies))
	sources := make([]<-chan Action, len(strategies))

	for i, strategy := range strategies {
		sources[i] = DenormalizeActions(
			strategy.Compute(snapshotsSplice[i]),
		)
	}

	return sources
}
