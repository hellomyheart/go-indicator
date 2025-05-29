package helper

// 将两个chan进行相加，返回一个新的chan
func Add[T Number](ac, bc <-chan T) <-chan T {
	return Operate(ac, bc, func(a, b T) T {
		return a + b
	})
}
