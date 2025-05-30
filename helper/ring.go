package helper

// 泛型循环缓冲区（Ring Buffer）数据结构
// 使用环的方法， 数据存满后，根据索引，逐步改变切片一个元素的值，调整指针完成缓存功能
type Ring[T any] struct {
	buffer []T  // 底层切片 保存缓存数据
	begin  int  // 当前有效数据起始索引
	end    int  // 当前有效数据结束索引 下一个要插入元素的索引
	empty  bool // 缓冲区是否为空
}

// 构造函数，创建具有给定缓存大小的实例
func NewRing[T any](size int) *Ring[T] {
	return &Ring[T]{
		buffer: make([]T, size),
		begin:  0,
		end:    0,
		empty:  true,
	}
}

// 将指定的值插入到环中并返回老数据（老数据可能为nil）
// value that was previously stored at that index.
func (r *Ring[T]) Put(t T) T {
	// 缓冲区已满，则将begin索引向右移动一位
	if r.IsFull() {
		r.begin = r.nextIndex(r.begin)
	}

	// 取出老数据
	o := r.buffer[r.end]
	// 放入元素到缓冲区
	r.buffer[r.end] = t

	// 更新end索引 end + 1
	r.end = r.nextIndex(r.end)
	// 更新空标志为false
	r.empty = false

	// 返回老数据
	return o
}

// Get从环缓冲区中检索可用值。如果为空，则返回默认值(T)和false。
// 从头获取，并且移动头
func (r *Ring[T]) Get() (T, bool) {
	var t T

	// 如果为空， 返回默认值T 和false
	if r.empty {
		return t, false
	}

	// 获取数据
	t = r.buffer[r.begin]
	// begin右移一位
	r.begin = r.nextIndex(r.begin)

	//如果空， 则将empty标志设置为true
	if r.begin == r.end {
		r.empty = true
	}

	// 返回头数据和true
	return t, true
}

// At返回给定索引处的值。
// 缓存只读取，不更改
// 下标从0开始
func (r *Ring[T]) At(index int) T {
	return r.buffer[(r.begin+index)%len(r.buffer)]
}

// IsEmpty 检查当前环缓冲区是否为空。
func (r *Ring[T]) IsEmpty() bool {
	return r.empty
}

// 判断缓冲区是否已满
func (r *Ring[T]) IsFull() bool {
	// 判断条件： 缓冲区不为空且
	// 				begin索引 == end索引
	return !r.empty && (r.end == r.begin)
}

// nextIndex返回环缓冲区中的下一个索引
func (r *Ring[T]) nextIndex(i int) int {
	return (i + 1) % len(r.buffer)
}
