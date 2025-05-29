package helper

// 排空chan
func Drain[T any](c <-chan T) {
	for {
		_, ok := <-c
		if !ok {
			break
		}
	}
}
