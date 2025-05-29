package helper

import (
	"errors"
	"fmt"
	"reflect"
)

// 检查多对chan的数据是否一致
// 第一个chen是参照对象，之后的chan都是被检查对象
func CheckEquals[T comparable](inputs ...<-chan T) error {
	// 非成对报错
	if len(inputs)%2 != 0 {
		return errors.New("not pairs")
	}

	//数据下标
	i := 0

	// 死循环，一直比较到chan关闭
	for {
		// j是偶数（是第奇数个）
		for j := 0; j < len(inputs); j += 2 {
			actual, actualOk := <-inputs[j]
			expected, expectedOk := <-inputs[j+1]

			// 至少一个通道关闭
			if !actualOk || !expectedOk {
				// 有一个通道没有关闭
				if actualOk != expectedOk {
					return fmt.Errorf("not ended the same actual %v expected %v", actualOk, expectedOk)
				}

				return nil
			}

			//数值比较
			if !reflect.DeepEqual(actual, expected) {
				return fmt.Errorf("index %d pair %d actual %v expected %v", i, j/2, actual, expected)
			}
		}

		i++
	}
}
