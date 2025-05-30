package helper

import (
	"errors"
	"reflect"
)

// Field从struct指针的通道中提取一个特定的字段，并通过一个新的通道传递。
func Field[T, S any](c <-chan *S, name string) (<-chan T, error) {
	// 判断是否为结构体
	st := reflect.TypeOf((*S)(nil)).Elem()
	if st.Kind() != reflect.Struct {
		return nil, errors.New("type not a struct")
	}

	// 获取结构体字段
	f, ok := st.FieldByName(name)
	if !ok {
		return nil, errors.New("field is not found")
	}

	result := make(chan T, cap(c))

	go func() {
		defer close(result)

		// 填充数据
		for n := range c {
			v := reflect.ValueOf(n).Elem()
			result <- v.FieldByIndex(f.Index).Interface().(T)
		}
	}()

	return result, nil
}
