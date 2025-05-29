package helper

// 将协程转换为切片
// 会阻塞主线程，只有chan关闭后，才返回切片

//	c := make(chan int, 4)
//	c <- 1
//	c <- 2
//	c < -3
//	c <- 4
//	close(c)
func ChanToSlice[T any](c <-chan T) []T {
	var slice []T

	for n := range c {
		slice = append(slice, n)
	}

	return slice
}
