package helper

import "sync"

// Waitable在从通道读取之前增加等待组
// 当通道关闭时表示完成。
func Waitable[T any](wg *sync.WaitGroup, c <-chan T) <-chan T {
	result := make(chan T, cap(c))

	// +1
	wg.Add(1)

	go func() {
		defer close(result)
		// -1
		defer wg.Done()

		for n := range c {
			result <- n
		}
	}()

	return result
}
