package helper

// 类似linux 管道命令
// Pipe函数接受一个输入通道和一个输出通道，并将输入通道中的所有元素复制到输出通道中。
// 将一个只读chan的元素复制到只写chan
// 例子：
//	input := helper.SliceToChan([]int{2, 4, 6, 8})
//	output := make(chan int)
//	helper.Pipe(input, output)
//	fmt.println(helper.ChanToSlice(output)) // [2, 4, 6, 8]
func Pipe[T any](f <-chan T, t chan<- T) {
	defer close(t)
	for n := range f {
		t <- n
	}
}
