package indicator

// 定义枚举类型
type TrendType int

// 使用 iota 实现枚举值自动递增
const (
	// _       TrendType = iota - 2 // 这样iota初始值变为-2
	FALLING TrendType = iota - 1 // FALLING = -1 下降
	STABLE                       // STABLE = 0 稳定
	RISING                       // RISING = 1 上升
)

// 实现 Stringer 接口（可选）
func (c TrendType) String() string {
	return [...]string{"FALLING", "STABLE", "RISING"}[c+1]
}

// 实现 int 接口
func (c TrendType) Int() int {
	return int(c)
}
