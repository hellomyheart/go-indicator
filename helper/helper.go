// 帮助类包
package helper

// 定义了 Integer 接口，表示可以接受任意一种类型
type Integer interface {
	int | int8 | int16 | int32 | int64
}

// 定义了Float 接口，表示可以接受任意一种类型
type Float interface {
	float32 | float64
}

// 定义了Number 接口，表示可以接受任意一种类型
type Number interface {
	Integer | Float
}
