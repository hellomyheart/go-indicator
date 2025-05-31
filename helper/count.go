package helper

// Count 生成一个数字序列，从指定的值开始,递增1，直到给定的其他通道继续产生值。
// 基于通道信号的递增计数器生成器，其核心作用是根据另一个通道的输入频率生成递增序列。
// Example:
//
//	other := make(chan int, 4)
//	other <- 1
//	other <- 1
//	other <- 1
//	other <- 1
//	close(other)
//
//	c := Count(0, other)
//
//	fmt.Println(<- s) // 1
//	fmt.Println(<- s) // 2
//	fmt.Println(<- s) // 3
//	fmt.Println(<- s) // 4
func Count[T Number, O any](from T, other <-chan O) <-chan T {
	c := make(chan T)

	go func() {
		defer close(c)

		for i := from; ; i++ {
			_, ok := <-other
			if !ok {
				break
			}

			c <- i
		}
	}()

	return c
}
