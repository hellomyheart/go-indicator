package helper_test

import (
	"testing"

	"github.com/hellomyheart/go-indicator/helper"
)

func TestAbs(t *testing.T) {
	// 将切片转为通道 测试数据
	input := helper.SliceToChan([]int{-10, 20, -4, -5})
	// 将切片转换为通道 预期结果
	expected := helper.SliceToChan([]int{10, 20, 4, 5})

	// 调用abs函数计算绝对值
	actual := helper.Abs(input)

	// 检查结果
	err := helper.CheckEquals(actual, expected)
	// 如果有错误，使用t.Fatal输出错误信息
	if err != nil {
		t.Fatal(err)
	}
}
