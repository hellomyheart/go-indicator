package helper

// Seq 生成以指定值开头的数字序列，从指定的量开始递增，递增，直到达到或超过指定值to。序列包括from和to。
//
// Example:
//
//	s := Seq(1, 5, 1)
//	defer close(s)
//
//	fmt.Println(<- s) // 1
//	fmt.Println(<- s) // 2
//	fmt.Println(<- s) // 3
//	fmt.Println(<- s) // 4
func Seq[T Number](from, to, increment T) <-chan T {
	c := make(chan T)

	go func() {
		for i := from; i < to; i += increment {
			c <- i
		}

		close(c)
	}()

	return c
}
