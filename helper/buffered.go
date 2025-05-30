package helper

// Buffered接受任何类型的通道，并返回具有指定大小的缓冲区的相同类型的新通道。这允许原始通道继续发送数据，即使接收端暂时不可用。

// 将一个无缓冲chan变成一个有缓冲chan
func Buffered[T any](c <-chan T, size int) <-chan T {
	result := make(chan T, size)

	go Pipe(c, result)

	return result
}
