package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestSync(t *testing.T) {
	input1 := helper.Skip(helper.SliceToChan([]int{0, 0, 0, 0, 1, 2, 3, 4}), 4)
	input2 := helper.Skip(helper.SliceToChan([]int{0, 0, 1, 2, 3, 4, 5, 6}), 2)
	input3 := helper.Skip(helper.SliceToChan([]int{0, 0, 0, 1, 2, 3, 4, 5}), 3)

	// 同步周期是4
	commonPeriod := helper.CommonPeriod(4, 2, 3)

	// 跳过部分元素实现周期对齐， 跳过0个元素
	actual1 := helper.SyncPeriod(commonPeriod, 4, input1)
	expected1 := helper.SliceToChan([]int{1, 2, 3, 4})

	// 跳过部分元素实现周期对齐 跳过2个元素
	actual2 := helper.SyncPeriod(commonPeriod, 2, input2)
	expected2 := helper.SliceToChan([]int{3, 4, 5, 6})

	// 跳过部分元素实现周期对齐 跳过1个元素
	actual3 := helper.SyncPeriod(commonPeriod, 3, input3)
	expected3 := helper.SliceToChan([]int{2, 3, 4, 5})

	err := helper.CheckEquals(actual1, expected1, actual2, expected2, actual3, expected3)
	if err != nil {
		t.Fatal(err)
	}
}
