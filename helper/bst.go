package helper

// BstNode 二叉搜索树节点 结构体
type BstNode[T Number] struct {
	value T
	left  *BstNode[T]
	right *BstNode[T]
}

// Bst 二叉搜索树 根节点 结构体
type Bst[T Number] struct {
	root *BstNode[T]
}

// NewBst 创建一个新的二叉搜索树。
func NewBst[T Number]() *Bst[T] {
	return &Bst[T]{}
}

// Insert 向二叉搜索树添加新值。
func (b *Bst[T]) Insert(value T) {
	// 创建一个新的节点
	// 使用&获取节点指针
	// &取址运算符
	// *取值运算符
	node := &BstNode[T]{
		value: value,
	}

	// 如果没有根节点，则将当前节点作为根节点
	if b.root == nil {
		b.root = node
		return
	}

	// 当前节点默认为root
	cur := b.root

	// 遍历子节点
	for {
		// 如果插入节点比当前节点小
		if node.value <= cur.value {
			// 如果当前节点的left是nil
			// 当前节点的left赋值为插入节点
			if cur.left == nil {
				cur.left = node
				return
			}
			//将当前节点置为 left 继续循环比较
			cur = cur.left
		} else {
			// 插入节点比当前节点大
			if cur.right == nil {
				//如果当前节点的right是nil
				// 当前节点的right赋值为插入节点
				cur.right = node
				return
			}
			// 将当前节点置为 right 继续循环比较
			cur = cur.right
		}
	}
}

// Contains 检查二叉搜索树是否存在该值
func (b *Bst[T]) Contains(value T) bool {
	node, _ := b.searchNode(value)
	return node != nil
}

// Remove 从二叉搜索树中删除指定的值并重新平衡树。
func (b *Bst[T]) Remove(value T) bool {
	// 找到值
	node, parent := b.searchNode(value)
	if node == nil {
		return false
	}

	// 删除值
	b.removeNode(node, parent)
	return true
}

// Min 函数返回二叉搜索树中的最小值。
func (b *Bst[T]) Min() T {
	if b.root == nil {
		return T(0)
	}

	node, _ := getMinNode(b.root)
	return node.value
}

// Max 函数返回二叉搜索树中的最大值。
func (b *Bst[T]) Max() T {
	if b.root == nil {
		return T(0)
	}

	node, _ := getMaxNode(b.root)
	return node.value
}

// searchNode 在二叉搜索树中搜索给定值，并返回第一个匹配节点及其父节点。
func (b *Bst[T]) searchNode(value T) (*BstNode[T], *BstNode[T]) {
	// 定义父节点
	var parent *BstNode[T]
	// 定义返回节点为root
	node := b.root

	// 循环查找 直到node为null
	for node != nil {
		// 如果node的value == value
		// 找到了,直接返回
		diff := value - node.value
		if diff == 0 {
			break
		}

		// 不等，父节点设置为当前node
		parent = node
		if diff < 0 {
			// value比node小， 则将node设置为left,继续循环
			node = node.left
		} else {
			// value比node大， 则将node设置为right,继续循环
			node = node.right
		}
	}

	// 有几种可能， 父节点为nil, node为root
	// 普通情况
	// node为nil 父节点为一个接近value的值
	return node, parent
}

// removeNode 从二叉搜索树中删除指定节点并重新平衡该树。
func (b *Bst[T]) removeNode(node, parent *BstNode[T]) {
	// 如果 要删除的节点的left right 都不为nil
	if node.left != nil && node.right != nil {
		// 获取该节点右子树的最小节点
		minNode, minParent := getMinNode(node.right)
		// 如果最小节点的父节点为nul
		if minParent == nil {
			// 说明右子树节点没有左子树
			// 同时也说明最小节点为右子树的根节点 父节点就是node
			// 最小节点赋值为node
			minParent = node
		}
		// 删除掉最小节点（在下一步会把最小节点的值赋值给node）
		b.removeNode(minNode, minParent)
		// 当前节点的值设置为minNode的值
		node.value = minNode.value
	} else {
		// 当前节点没有子节点或者只有一个子节点
		var child *BstNode[T]

		// 如果有左节点
		// 让child节点尽可能不是nil
		if node.left != nil {
			child = node.left
		} else {
			child = node.right
		}

		// 如果要删除的节点是根节点
		if node == b.root {
			// 根节点重设为child
			b.root = child
		} else if parent.left == node {
			// 如果父节点的left 是node
			// 将父节点的left设置为child
			parent.left = child
		} else {
			// 将父节点的left设置为child
			parent.right = child
		}
	}
}

// getMinNode 函数返回具有最小值的节点及其父节点。
func getMinNode[T Number](root *BstNode[T]) (*BstNode[T], *BstNode[T]) {
	var parent *BstNode[T]
	node := root

	for node.left != nil {
		parent = node
		node = node.left
	}

	return node, parent
}

// getMaxNode 函数返回具有最大值的节点及其父节点。
func getMaxNode[T Number](root *BstNode[T]) (*BstNode[T], *BstNode[T]) {
	var parent *BstNode[T]
	node := root

	for node.right != nil {
		parent = node
		node = node.right
	}

	return node, parent
}
