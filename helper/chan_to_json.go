package helper

import (
	"encoding/json"
	"io"
)

// ChanToJSON将值通道转换为JSON格式，并将其写入指定的写入器。
// 感觉是将对象数组转为json数组
//
// 例如:
//
//	input := helper.SliceToChan([]int{2, 4, 6, 8})
//
//	var buffer bytes.Buffer
//	err := helper.ChanToJSON(input, &buffer)
//
//	fmt.Println(buffer.String())
//	// Output: [2,4,6,8,9]
func ChanToJSON[T any](c <-chan T, w io.Writer) error {
	first := true

	_, err := w.Write([]byte{'['})
	if err != nil {
		return err
	}

	for n := range c {
		if !first {
			_, err = w.Write([]byte{','})
			if err != nil {
				return err
			}
		} else {
			first = false
		}

		encoded, err := json.Marshal(n)
		if err != nil {
			return err
		}

		_, err = w.Write(encoded)
		if err != nil {
			return err
		}
	}

	_, err = w.Write([]byte{']'})

	return err
}
