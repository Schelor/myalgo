package binarytrees

import "container/list"

// 定义树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历：先遍历根结点,然后再遍历左子树,最后遍历右子树
func PreorderTraversal(root *TreeNode) []int {
	return PreorderTraversalByRecursive(root)
}

func PreorderTraversalByRecursive(root *TreeNode) []int {
	// 存放遍历结果
	list := make([]int, 0)
	// 定义一个函数类型变量,内部递归调用该函数,函数内部访问外部的list,即闭包调用
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		list = append(list, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return list
}

// 循环方式前序遍历
func PreorderTraversalByLoop(root *TreeNode) []int {
	list := make([]int, 0)
	stack := make([]*TreeNode, 0)
	node := root
	for node != nil || len(stack) != 0 {
		for node != nil {
			list = append(list, node.Val) // 前序遍历,先记录当前节点
			stack = append(stack, node)   // 当前节点入栈,方便后续遍历其右子树
			node = node.Left              // 继续遍历左子树
		}
		p := stack[len(stack)-1]     // 后进先出,出栈最后加入的结点,p可能是一个最左叶子,也可能是一个父节点(左子节点为空)
		stack = stack[:len(stack)-1] // golang中没有Stack这种数据结构,用数组来模拟栈,出栈后要更新当前栈
		node = p.Right               // 继续遍历p的右子树
	}
	return list
}

func preorderTraversal(root *TreeNode) (vals []int) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

// 中序遍历,递归实现
func InorderTraversalByRecursive(root *TreeNode) []int {
	// 存放遍历结果
	list := make([]int, 0)
	// 定义一个函数类型变量,内部递归调用该函数,函数内部访问外部的list,即闭包调用
	var traversal func(node *TreeNode) // 先定义
	traversal = func(node *TreeNode) { // 函数变量赋值
		if node == nil { // 当前函数栈退出
			return
		}
		traversal(node.Left)          // 先遍历当前节点左子树，函数入栈
		list = append(list, node.Val) // 再遍历当前节点
		traversal(node.Right)         // 再遍历当前节点右子树,函数入栈
	}
	traversal(root)
	return list
}

// 中序遍历,循环实现,借助栈
func InorderTraversalByLoop(root *TreeNode) []int {
	list := make([]int, 0)
	stack := make([]*TreeNode, 0)
	current := root
	for current != nil || len(stack) != 0 {
		for current != nil {
			stack = append(stack, current) // 入栈，直到找到最左子节点
			current = current.Left
		}
		p := stack[len(stack)-1]     // 出栈,该结点表示最左子节点,或左节点为空的父节点
		list = append(list, p.Val)   // 记录该结点(表示先遍历左子节点)
		stack = stack[:len(stack)-1] // 出栈栈顶结点后更新栈
		current = p.Right            // 继续遍历右子树
	}
	return list
}

// 后续遍历,递归实现
func PostorderTraversalByRecursive(root *TreeNode) []int {
	list := make([]int, 0)             // 存储遍历结果
	var traversal func(*TreeNode)      // 定义函数变量
	traversal = func(node *TreeNode) { // 函数变量赋值/函数实现
		if node == nil {
			return
		}
		traversal(node.Left)          // 先遍历当前节点的左子树
		traversal(node.Right)         // 再遍历当前节点的右子树
		list = append(list, node.Val) // 最后遍历当前节点
	}
	traversal(root)
	return list
}

// 后续遍历,循环实现
func PostorderTraversalByLoop(root *TreeNode) []int {
	list := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var current = root
	// 由于在某颗子树访问完成以后，接着就要回溯到其父节点去
	// 因此可以用prev来记录访问历史，在回溯到父节点时，可以由此来判断，上一个访问的节点是否为右子树。如果是则表示当前为一个父节点
	var prev *TreeNode // 所谓访问过,即记录过该结点的值,因此用prev来标识该结点
	for current != nil || len(stack) != 0 {
		for current != nil {
			stack = append(stack, current) // 入栈，直到找到最左子节点
			current = current.Left
		}
		// 从栈中弹出的元素，左子树一定是访问完了的
		current = stack[len(stack)-1] // 出栈,该结点表示最左子节点,或左节点为空的父节点
		stack = stack[:len(stack)-1]  // 出栈栈顶结点后更新栈
		// 现在需要确定的是是否有右子树，或者右子树是否访问过
		// 如果没有右子树，或者右子树访问完了，也就是上一个访问的节点是右子节点时 说明可以访问当前节点
		if current.Right == nil || prev == current.Right {
			list = append(list, current.Val)
			// 更新历史访问记录，这样回溯的时候父节点可以由此判断右子树是否访问完成
			prev = current
			current = nil
		} else {
			// 入栈，遍历右子树
			stack = append(stack, current)
			current = current.Right
		}
	}
	return list
}

// 层序遍历,借助队列(先使用container下的双端链表
func LevelOrder(root *TreeNode) [][]int {
	nodeList := make([][]int, 0) // 存放每层遍历的结果
	if root == nil {
		return nodeList
	}
	var queue *list.List = list.New() // 定义一个队列
	queue.PushBack(root)              // 根结点入队，队尾入队
	for queue.Len() > 0 {
		size := queue.Len() // 当前层节点数量,由于可能入队下一层的结点,为了出队固定数量的结点,这里只能使用Len()
		levelNodes := make([]int, 0)
		for i := 0; i < size; i++ {
			var node *TreeNode = queue.Remove(queue.Front()).(*TreeNode) // 出队队首结点
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			levelNodes = append(levelNodes, node.Val)
		}
		nodeList = append(nodeList, levelNodes)
	}
	return nodeList
}

// 层序遍历，通过数组来模拟队列数据结构,通过2个数组来模拟队列,其中一个用来遍历当前层,其中一个用来记录下一层
func LevelOrderByArray(root *TreeNode) [][]int {
	nodeList := make([][]int, 0) // 存放每层遍历的结果
	if root == nil {
		return nodeList
	}
	var currentLevelQueue = []*TreeNode{root} // 数组类型为*TreeNode,默认入队root,当前层
	for len(currentLevelQueue) > 0 {
		currentLevelNodes := make([]int, 0)
		nextLevelQueue := make([]*TreeNode, 0) // 准备通过当前层生成下一层
		for i := 0; i < len(currentLevelQueue); i++ {
			node := currentLevelQueue[i] // 当前层依次出队
			currentLevelNodes = append(currentLevelNodes, node.Val)
			if node.Left != nil { // 下一层有节点
				nextLevelQueue = append(nextLevelQueue, node.Left)
			}
			if node.Right != nil {
				nextLevelQueue = append(nextLevelQueue, node.Right)
			}

		}
		nodeList = append(nodeList, currentLevelNodes) // 当前层加入结果
		currentLevelQueue = nextLevelQueue             // 将下一层更新为当前层
	}
	return nodeList
}

// 获取最大深度,用数组来模拟队列
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	max := 0
	for len(queue) > 0 {
		size := len(queue) // 出队当前层的节点
		for i := 0; i < size; i++ {
			node := queue[0]  // 队首出队
			queue = queue[1:] // 更新队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		max++
	}
	return max
}

func LevelOrderBottom(root *TreeNode) [][]int {
	levelOrder := LevelOrder(root)
	for i, j := 0, len(levelOrder)-1; i < j; i, j = i+1, j-1 {
		levelOrder[i], levelOrder[j] = levelOrder[j], levelOrder[i]
	}
	return levelOrder
}

// 二叉树右视图
func RightSideView(root *TreeNode) []int {

}
