package binarytrees

import (
	"container/list"
	"math"
	"strconv"
)

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

// 二叉树右视图,用数组模拟队列
func RightSideView(root *TreeNode) []int {
	view := make([]int, 0)
	if root == nil {
		return view
	}
	queue := []*TreeNode{root} // 定义一个队列
	for len(queue) > 0 {
		size := len(queue)
		for i := 1; i <= size; i++ {
			node := queue[0]  // 队首出队结点,队首结点固定在下标为0
			queue = queue[1:] // 出队后,更新队列
			if i == size {    // 当前层最后一个节点
				view = append(view, node.Val)
			}
			// 下一层
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return view
}

func AverageOfLevels(root *TreeNode) []float64 {
	avgs := make([]float64, 0)
	if root == nil {
		return avgs
	}
	queue := list.New() // 通过双端链表来作为队列
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		sum := 0
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			sum += node.Val
			if node.Left != nil { // 下一层入队列
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		// 计算当前层的平均值
		avgs = append(avgs, float64(sum)/float64(size))

	}
	return avgs
}

//type Node struct {
//	Val      int
//	Children []*Node
//}

//func NTreeLevelOrder(root *Node) [][]int {
//	nodeList := make([][]int, 0) // 存放每层遍历的结果
//	if root == nil {
//		return nodeList
//	}
//	var queue *list.List = list.New() // 定义一个队列
//	queue.PushBack(root)              // 根结点入队，队尾入队
//	for queue.Len() > 0 {
//		// 当前层节点数量,由于可能入队下一层的结点,为了出队固定数量的结点,这里只能使用Len()
//		size := queue.Len()
//		levelNodes := make([]int, 0)
//		for i := 0; i < size; i++ {
//			var node *Node = queue.Remove(queue.Front()).(*Node) // 出队队首结点
//			// 检查下一层的所有结点，如果存在则入队
//			if node.Children != nil && len(node.Children) > 0 {
//				for _, v := range node.Children {
//					queue.PushBack(v)
//				}
//			}
//			levelNodes = append(levelNodes, node.Val)
//		}
//		nodeList = append(nodeList, levelNodes)
//	}
//	return nodeList
//}

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	min := 0
	for len(queue) > 0 {
		size := len(queue) // 出队当前层的节点
		min++
		for i := 0; i < size; i++ {
			node := queue[0]  // 队首出队
			queue = queue[1:] // 更新队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			// 当前层的某个结点,无子节点,其所对应的层次即最小深度
			if node.Left == nil && node.Right == nil {
				return min
			}
		}
	}
	return min
}
func MinDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	depth := math.MaxInt
	if root.Left != nil {
		depth = min(depth, MinDepth2(root.Left))
	}
	if root.Right != nil {
		depth = min(depth, MinDepth2(root.Right))
	}
	return depth + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func LargestValues(root *TreeNode) []int {
	nodeList := make([]int, 0) // 存放每层遍历的结果
	if root == nil {
		return nodeList
	}
	var queue *list.List = list.New() // 定义一个队列
	queue.PushBack(root)              // 根结点入队，队尾入队
	for queue.Len() > 0 {
		size := queue.Len() // 当前层节点数量,由于可能入队下一层的结点,为了出队固定数量的结点,这里只能使用Len()
		maxVal := math.MinInt
		for i := 0; i < size; i++ {
			var node *TreeNode = queue.Remove(queue.Front()).(*TreeNode) // 出队队首结点
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			maxVal = max(maxVal, node.Val)
		}
		nodeList = append(nodeList, maxVal)
	}
	return nodeList
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 找到每层结点，放入到列表，然后遍历这个列表,依次设置节点next
// 时间复杂度O(n),遍历N个结点, 空间复杂度O(N) 存放每层的结点
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := list.New() // 定义队列
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()                       // 当前层的节点数量,只出队固定数量的结点
		var levelNodes []*Node = make([]*Node, 0) // 存放每层的结点
		for i := 0; i < size; i++ {
			var node *Node = queue.Remove(queue.Front()).(*Node) // 出队队首结点
			// 检查下一层的所有结点，如果存在则入队
			if node.Left != nil { // 下一层入队列
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			levelNodes = append(levelNodes, node)
		}
		// 遍历当前层节点,依次设置next
		for i := 0; i < len(levelNodes)-1; i++ {
			levelNodes[i].Next = levelNodes[i+1]
		}
	}
	return root
}

func connect2(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := list.New() // 定义队列
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len() // 当前层的节点数量,只出队固定数量的结点
		prevNode := queue.Front().Value.(*Node)
		for i := 0; i < size; i++ {
			var node *Node = queue.Remove(queue.Front()).(*Node) // 出队队首结点
			// 本层依次出队的结点,从第二个开始，让前一个依次执行当前节点
			if i >= 1 {
				prevNode.Next = node
				prevNode = node
			}
			// 检查下一层的所有结点，如果存在则入队
			if node.Left != nil { // 下一层入队列
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}

		}
	}
	return root
}

func InvertTree(root *TreeNode) *TreeNode {
	return InvertTree1(root)
}

// 翻转二叉树,递归法
// 时间复杂度：O(N), 空间复杂度O(N)
func InvertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = InvertTree1(root.Right), InvertTree1(root.Left)
	return root
}

// 翻转二叉树,层序遍历
// 时间复杂度：O(N), 空间复杂度O(N)
func InvertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := list.New() // 定义队列
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len() // 当前层的节点数量,只出队固定数量的结点
		for i := 0; i < size; i++ {
			node := (queue.Remove(queue.Front())).(*TreeNode) // 当前层的每一个结点
			node.Left, node.Right = node.Right, node.Left     // 交换本层左右结点
			if node.Left != nil {                             // 下一层入队列
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return root
}

// 翻转二叉树,递归法-前序遍历模式
// 时间复杂度：O(N), 空间复杂度O(N)
func InvertTreeV1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// 先遍历当前节点,直接读取left,right并交换
	root.Left, root.Right = root.Right, root.Left // 当前节点的左右子节点交换
	InvertTreeV1(root.Left)                       // 反转左子树
	InvertTreeV1(root.Right)                      // 反正右子树
	return root
}

// 翻转二叉树,迭代法-前序遍历模式-借助栈
// 时间复杂度：O(N), 空间复杂度O(N)
func InvertTreeV2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := make([]*TreeNode, 0) // 用数组模拟栈
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			// 前序遍历，遇到一个结点直接处理，自顶向下
			node.Left, node.Right = node.Right, node.Left
			stack = append(stack, node)
			node = node.Left
		}
		// 到达最左子节
		node = stack[len(stack)-1] // 出栈栈顶结点
		stack = stack[0 : len(stack)-1]
		node = node.Right // 继续考察是否还有右子节点
	}
	return root
}

// 翻转二叉树,迭代法-后续序遍历模式-借助栈
// 时间复杂度：O(N), 空间复杂度O(N)
func InvertTreeV3(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := make([]*TreeNode, 0) // 用数组模拟栈
	node := root
	var pre *TreeNode = nil // 上一次遍历过/上一次处理过的结点
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node) // 自顶向下遍历,上层的节点先入栈
			node = node.Left
		}
		// 到达最左子节点或最左子节点的父节点且该父节点有右子结点
		node = stack[len(stack)-1] // 出栈栈顶结点
		stack = stack[0 : len(stack)-1]
		if node.Right == nil || node.Right == pre {
			node.Left, node.Right = node.Right, node.Left
			pre = node // 当前节点已处理完成,标记为上次处理的结点
			node = nil // 继续出栈下一个结点
		} else {
			stack = append(stack, node) // 存在右子结点，继续入栈
			node = node.Right
		}
	}
	return root
}

// 对称二叉树
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 通过递归来实现
	return compareNode(root.Left, root.Right)
}

// 递归三步骤
// 第一步：确定递归函数的参数和返回值： 确定哪些参数是递归的过程中需要处理的，那么就在递归函数里加上这个参数，
// 并且还要明确每次递归的返回值是什么进而确定递归函数的返回类型。
// 第二步：确定终止条件： 写完了递归算法, 运行的时候，经常会遇到栈溢出的错误，
//
//	就是没写终止条件或者终止条件写的不对，操作系统也是用一个栈的结构来保存每一层递归的信息，
//	如果递归没有终止，操作系统的内存栈必然就会溢出。
//
// 第三步：确定单层递归的逻辑，确定每一层递归需要处理的信息。在这里也就会重复调用自己来实现递归的过程。
func compareNode(left *TreeNode, right *TreeNode) bool {
	// 递归终止分解条件
	if left == nil && right == nil { // 左右子节点都为空，对称
		return true
	}
	if left != nil && right == nil { // 左不为空,右为空,不对称
		return false
	}
	if left == nil && right != nil { // 左为空,右不为空,不对称
		return false
	}
	if left.Val != right.Val { // 左右不为空,但值不相同
		return false
	}
	// 到这里表示左右都不为空,此时需要向下分解(递归)
	// 逻辑是要判断当前左右节点是否对称 先判断外侧其左子节点与右子节点是否对称，
	// 再判断内测右子节点与左子节点是否对称
	outside := compareNode(left.Left, right.Right)
	inside := compareNode(left.Right, right.Left)
	return outside && inside // 外侧相同且内侧相同才对称
}

// 对称二叉树,通过迭代来实现，这里借住队列
// 依次入队外侧两结点，内侧两节点
func IsSymmetricV2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := list.New() // 用双向链表来模拟队列
	queue.PushBack(root.Left)
	queue.PushBack(root.Right)
	for queue.Len() > 0 {
		// 依次从队列里取出2个结点（这两个节点分别代表外侧，内侧对称的两个结点)
		left := queue.Remove(queue.Front()).(*TreeNode)
		right := queue.Remove(queue.Front()).(*TreeNode)
		// 如果左右都为空,这种属于对称,提前处理
		if left == nil && right == nil {
			continue
		}
		// 不对称的场景条件
		// 1.左不为空,右为空
		// 2.左为空,右不为空
		// 3.左右值不同
		if left != nil && right == nil {
			return false
		}
		if left == nil && right != nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		// 剩下的情况为左右都不为空,且值相同,继续看下一层
		// 依次入队外侧2结点，内侧2节点
		queue.PushBack(left.Left)   // 加入左节点左孩子
		queue.PushBack(right.Right) // 加入右节点右孩子
		queue.PushBack(left.Right)  // 加入左节点右孩子
		queue.PushBack(right.Left)  // 加入右节点左孩子
	}
	return true // 迭代就自顶向下,如果都对称,返回true
}

// 对称二叉树,通过迭代来实现，基于数组来模拟队列
// 依次入队外侧两结点，内侧两节点
func IsSymmetricV3(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		// 依次从队列里取出2个结点（这两个节点分别代表外侧，内侧对称的两个结点)
		left, right := queue[0], queue[1]
		queue = queue[2:] // 出队2个结点
		// 如果左右都为空,这种属于对称,提前处理
		if left == nil && right == nil {
			continue
		}
		// 不对称的场景条件
		// 1.左不为空,右为空
		// 2.左为空,右不为空
		// 3.左右值不同
		if left != nil && right == nil {
			return false
		}
		if left == nil && right != nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		// 剩下的情况为左右都不为空,且值相同,继续看下一层
		// 依次入队外侧2结点，内侧2节点
		// 加入左节点左孩子
		// 加入右节点右孩子
		// 加入左节点右孩子
		// 加入右节点左孩子
		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}
	return true // 迭代就自顶向下,如果都对称,返回true
}

// 求完全二叉树的节点数，直接遍历
func countNodesV1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	count := 0
	for len(queue) > 0 {
		size := len(queue) // 当前层的节点数
		count += size      // 累计节点数
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return count
}

// 求完全二叉树的节点数-递归法
func countNodesV2(root *TreeNode) int {
	return sumNode(root)
}

// 递归三步骤:第一步确定递归函数的入参返回值
// 求解当前节点的子树节点数,入参为一个节点,返回值为该节点的包含的所有子节点数
func sumNode(node *TreeNode) int {
	// 第二步：确定结束分解的条件
	// 如果当前节点为空,不能再继续分解了则结束
	if node == nil {
		return 0
	}
	// 第三步：确认递归过程
	// 确定单层递归的逻辑，确定每一层递归需要处理的信息。在这里也就会重复调用自己来实现递归的过程
	sumLeft := sumNode(node.Left)   // 获取左子节点
	sumRight := sumNode(node.Right) // 获取右子节点数
	return sumLeft + sumRight + 1   // 加上当前节点数-固定为1
}

// 判断是否为平衡二叉树,高度差绝对值小于等于1
// 递归求解
func isBalancedV1(root *TreeNode) bool {
	// 返回以该节点为根节点的二叉树的高度，如果不是平衡二叉树了则返回-1
	h := getHeight(root)
	if h == -1 {
		return false
	}
	return true
}

// 递归三步骤：1.确定函数入参/返回值
// 入参为每个节点,返回值为当前节点的高度
func getHeight(node *TreeNode) int {
	// 2.确定终止分解的条件
	if node == nil {
		return 0
	}
	leftHeight := getHeight(node.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := getHeight(node.Right)
	if rightHeight == -1 {
		return -1
	}
	// 绝对值差大于1，也不满足平衡树
	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1 // 返回当前节点的实际高度
}

// 二叉树所有路径-递归-自顶向下分解-自底向上返回
func BinaryTreePathsV1(root *TreeNode) []string {
	return nodePaths(root)
}
func nodePaths(node *TreeNode) []string {
	if node == nil {
		return []string{}
	}
	leftPaths := nodePaths(node.Left)
	rightPaths := nodePaths(node.Right)
	if len(leftPaths) == 0 && len(rightPaths) == 0 {
		paths := []string{strconv.Itoa(node.Val)}
		return paths
	}
	paths := make([]string, 0)
	if len(leftPaths) > 0 {
		for _, v := range leftPaths {
			paths = append(paths, strconv.Itoa(node.Val)+"->"+v)
		}
	}
	if len(rightPaths) > 0 {
		for _, v := range rightPaths {
			paths = append(paths, strconv.Itoa(node.Val)+"->"+v)
		}
	}
	return paths
}

// 二叉树所有路径-递归-自顶向下-分解过程中直接拼接路径,然后加入结果集
// 这种递归没有返回值,是直接访问一个公共内存变量,基于语言特性才有的用法
// 这种方式不用每个节点都申请一份数组空间,复用同一个结果存放空间
func BinaryTreePathsV2(root *TreeNode) []string {
	result := make([]string, 0)
	// 递归三步骤:1.定义递归函数入参/返回值
	var nodePaths func(node *TreeNode, path string) // 这里定义一个函数变量,走闭包访问
	nodePaths = func(node *TreeNode, path string) {
		// 递归三步骤: 2.定义终止条件
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil { // 到叶子结点了
			path = path + strconv.Itoa(node.Val)
			result = append(result, path)
			return
		}
		path = path + strconv.Itoa(node.Val) + "->"
		// 递归三步骤: 3.定义本层调用逻辑,递归过程是怎样的
		if node.Left != nil {
			nodePaths(node.Left, path)
		}
		if node.Right != nil {
			nodePaths(node.Right, path)
		}
	}
	nodePaths(root, "")
	return result
}

// 左叶子节点之和-通过层序遍历
// 遍历当前层时,需要识别下一层节点是否为左叶子
func sumOfLeftLeavesV2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	sum := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil { // 左叶子一定在Left节点
				// node.Left左叶子,加入累加和
				if node.Left.Left == nil && node.Left.Right == nil {
					sum += node.Left.Val
				}
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return sum
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归-前序遍历
func sumOfLeftLeavesV3(root *TreeNode) int {
	res := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 如果节点为左叶子,加入累计和
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			res += node.Left.Val
		}
		dfs(node.Left)  // 继续看左叶子的子节点
		dfs(node.Right) // 继续看右叶子的子节点
	}
	dfs(root)

	return res
}

// 找树左下角的值-层序遍历取第一个节点,并依次覆盖
func FindBottomLeftValueV1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	target := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			if i == 0 {
				target = queue[0].Val
			}
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return target
}

// 路径总和-层序遍历
func HasPathSumV1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	nodeQueue := list.New() // 用双端链表来模拟队列
	sumQueue := list.New()  // 用双端链表来模拟队列
	nodeQueue.PushBack(root)
	sumQueue.PushBack(root.Val)
	for nodeQueue.Len() > 0 {
		// 这种出队模式,与用一个size是类似的,都是要把所有节点出队完成
		// 如果用一个size来出队,表示可以控制当前层节点个数
		node := (nodeQueue.Remove(nodeQueue.Front())).(*TreeNode)
		currentNodeSum := (sumQueue.Remove(sumQueue.Front())).(int)
		if node.Left == nil && node.Right == nil {
			if currentNodeSum == targetSum {
				return true
			}
		}
		if node.Left != nil {
			nodeQueue.PushBack(node.Left)
			sumQueue.PushBack(currentNodeSum + node.Left.Val)
		}
		if node.Right != nil {
			nodeQueue.PushBack(node.Right)
			sumQueue.PushBack(currentNodeSum + node.Right.Val)
		}
	}
	return false
}

// 路径总和-递归-语言自带的闭包语法
func HasPathSumV2(root *TreeNode, targetSum int) bool {
	// 递归三部曲:1.定义函数入参/返回值
	var hasPathSum func(node *TreeNode, sum int) bool
	hasPathSum = func(node *TreeNode, sum int) bool {
		// 2.定义终止条件
		if node == nil {
			return false
		}
		sum += node.Val
		// 当前节点为子节点,且满足路径和等于target
		if node.Left == nil && node.Right == nil && sum == targetSum {
			return true
		}
		// 3.定义递归逻辑
		// 当前节点不满足,继续分解检查左右子节点
		checkLeft := hasPathSum(node.Left, sum)
		checkRight := hasPathSum(node.Right, sum)
		return checkLeft || checkRight
	}
	return hasPathSum(root, 0)
}

// 路径和II-层序遍历
func PathSumV1(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	nodeQueue := list.New() // 用双端链表来模拟队列
	sumQueue := list.New()  // 用双端链表来模拟队列
	nodeQueue.PushBack(root)
	sumQueue.PushBack([]int{root.Val})
	for nodeQueue.Len() > 0 {
		// 这种出队模式,与用一个size是类似的,都是要把所有节点出队完成
		// 如果用一个size来出队,表示可以控制当前层节点个数
		node := (nodeQueue.Remove(nodeQueue.Front())).(*TreeNode)
		paths := (sumQueue.Remove(sumQueue.Front())).([]int)
		// 当前节点为叶子节点，路径和匹配目标值
		if node.Left == nil && node.Right == nil && sumInt(paths) == targetSum {
			result = append(result, paths)
		}
		if node.Left != nil {
			nodeQueue.PushBack(node.Left)
			// 创建新的path,不能复用原有path,append是创建一个新的对象
			leftPaths := make([]int, len(paths))
			copy(leftPaths, paths)
			leftPaths = append(leftPaths, node.Left.Val)
			sumQueue.PushBack(leftPaths)
		}
		if node.Right != nil {
			nodeQueue.PushBack(node.Right)
			// 创建新的path,不能复用原有path,append是创建一个新的对象
			rightPaths := make([]int, len(paths))
			copy(rightPaths, paths)
			rightPaths = append(rightPaths, node.Right.Val)
			sumQueue.PushBack(rightPaths)
		}
	}
	return result
}

func sumInt(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}

// 路径和II-递归
func PathSumV3(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	var pathSum func(node *TreeNode, sum []int)
	pathSum = func(node *TreeNode, sum []int) {
		if node == nil {
			return
		}
		// 这里也要申请新内存,切片是引用传递,底层是同一个数组
		newSum := make([]int, len(sum))
		copy(newSum, sum)
		newSum = append(newSum, node.Val)
		if node.Left == nil && node.Right == nil && sumInt(newSum) == targetSum {
			result = append(result, newSum)
		}
		pathSum(node.Left, newSum)
		pathSum(node.Right, newSum)
	}
	pathSum(root, []int{})
	return result
}

// 路径和II-层序遍历
func PathSumV2(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	nodeQueue := []*TreeNode{root} // 这里用数组来模拟队列,存储每个节点
	sumQueue := []int{root.Val}    // 这里用数组来模拟队列,存储每个节点的路径和
	nodeMap := make(map[*TreeNode]*TreeNode)
	for len(nodeQueue) > 0 {
		// 这种出队模式,与用一个size是类似的,都是要把所有节点出队完成
		// 如果用一个size来出队,表示可以控制当前层节点个数
		node, pathsSum := nodeQueue[0], sumQueue[0]
		nodeQueue, sumQueue = nodeQueue[1:], sumQueue[1:] // 出队,更新队列
		// 当前节点为叶子节点，路径和匹配目标值
		if node.Left == nil && node.Right == nil && pathsSum == targetSum {
			result = append(result, findPath(node, nodeMap))
		}
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			sumQueue = append(sumQueue, pathsSum+node.Left.Val)
			nodeMap[node.Left] = node
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			sumQueue = append(sumQueue, pathsSum+node.Right.Val)
			nodeMap[node.Right] = node
		}
	}
	return result
}

// 查找当前节点的完整路径
func findPath(node *TreeNode, nodeMap map[*TreeNode]*TreeNode) []int {
	path := make([]int, 0) // 默认放入的是自底向上的路径,最后需要做一次反转
	for node != nil {
		path = append(path, node.Val)
		node = nodeMap[node]
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
