package helper

// 将切片转为chan
// 定义一个chan，用go启动一个协程，然后通过for循环将切片中的元素一个一个写入chan中
// 在函数主线程返回chan.

// 用法示例：
//	slice := []float64{2, 4, 6, 8}
//	c := helper.SliceToChan(slice)
//	fmt.Println(<- c)  // 2
//	fmt.Println(<- c)  // 4
//	fmt.Println(<- c)  // 6
//	fmt.Println(<- c)  // 8
func SliceToChan[T any](slice []T) <-chan T {
	c := make(chan T)

	go func() {
		defer close(c)

		for _, n := range slice {
			c <- n
		}
	}()

	return c
}
