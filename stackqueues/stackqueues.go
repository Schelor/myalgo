package stackqueues

import (
	"container/list"
	"strconv"
)

// 用两个栈模拟队列
// 定义2个stack,一个用于输入,一个用以输出
type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

// 输入场景,用inStack
func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

// 输出场景,用outStack
func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		this.inToOut()
	}
	lastIndex := len(this.outStack) - 1
	v := this.outStack[lastIndex]
	this.outStack = this.outStack[:lastIndex] // start默认为0
	return v
}

// 输出场景,用outStack
func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		this.inToOut()
	}
	lastIndex := len(this.outStack) - 1
	v := this.outStack[lastIndex]
	return v
}

func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

// inStack数组内容全部转入到outStack
// 按照先入后出的顺序读取inStack
func (this *MyQueue) inToOut() {
	for len(this.inStack) > 0 {
		// 后进先出
		lastIndex := len(this.inStack) - 1
		v := this.inStack[lastIndex]             // 最后进入stack的元素在末尾
		this.inStack = this.inStack[:lastIndex]  // 用新切片替换原来的数组
		this.outStack = append(this.outStack, v) // 放入outStack
	}
}

// IsValid 有效字符, 当前字符在ASCII码范围,可用byte或int8来表示
func IsValid(s string) bool {
	// 用数组来模拟栈,数组最大下标就是栈顶元素
	stack := make([]byte, 0)
	mapping := map[byte]byte{')': '(', ']': '[', '}': '{'}
	for i := 0; i < len(s); i++ {
		cur := s[i]
		if pair, ok := mapping[cur]; ok { // 如果包含右括号,在栈顶找左括号
			top := byte('#')
			if len(stack) > 0 { // 出栈,更新stack
				top, stack = stack[len(stack)-1], stack[:len(stack)-1]
			}
			if pair != top {
				return false
			}
		} else { // 如果是左括号直接入栈
			stack = append(stack, cur)
		}
	}
	return len(stack) == 0
}

func IsValid2(s string) bool {
	// 用数组来模拟栈,数组最大下标就是栈顶元素
	stack := make([]byte, 0)
	mapping := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	for i := 0; i < len(s); i++ {
		if _, ok := mapping[s[i]]; ok { // 如果包含左括号,压栈
			stack = append(stack, s[i])
			continue
		}
		// 如果是右括号,则从栈顶获取，然后检查是否与当前字符匹配
		if len(stack) == 0 {
			return false
		}
		// 出栈,更新stack
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if mapping[top] != s[i] {
			return false
		}
	}
	return len(stack) == 0
}

// RemoveDuplicates 基于栈来移除重复的字符
func RemoveDuplicates(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		// 栈为空直接入栈,例如i=0或前面都是相同的字符
		if len(stack) == 0 {
			stack = append(stack, s[i])
			continue
		}
		// 检查当前与栈top是否相同,如果相同,则丢弃,如果不同则入栈
		peekTop := stack[len(stack)-1]
		if s[i] == peekTop { // 当前与前一个(位于栈顶)相同,出栈TOP
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

// 后缀表达式
func EvalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err == nil { // 数字则正常转换
			stack = append(stack, num)
		} else {
			// 依次出栈2个数
			num1, num2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num2-num1)
			case "*":
				stack = append(stack, num2*num1)
			case "/":
				stack = append(stack, num2/num1)
			}
		}
	}
	return stack[len(stack)-1]
}

// 定义一个单调队列，内部基于双端链表实现
// 比较耗内存
type MonoQueue struct {
	deque *list.List
}

// 单调递减队列,入队方法
// 如果offer的元素value大于队尾(入口）元素的数值，那么就将队列入口的元素弹出
// 直到offer元素的数值小于等于队列入口元素的数值为止,然后放入队尾
// 这样就保持了队列里的数值是单调从大到小的了。
func (this *MonoQueue) offer(value int) {
	// 将尾部的元素转换为int类型
	for this.deque.Len() != 0 && value > this.deque.Back().Value.(int) {
		this.deque.Remove(this.deque.Back())
	}
	this.deque.PushBack(value)
}

// 从队列中移除指定的元素
func (this *MonoQueue) poll(value int) {
	if this.deque.Len() != 0 && this.deque.Front().Value.(int) == value {
		this.deque.Remove(this.deque.Front())
	}
}

// 获取单调队列中队首的元素,即队列中最大值
func (this *MonoQueue) peek() int {
	return this.deque.Front().Value.(int)
}

// 基于单调递减队列来查找滑动窗口最大值
func MaxSlidingWindow(nums []int, k int) []int {
	n := len(nums)                              // 数组长度
	result, arrayIndex := make([]int, n-k+1), 0 // 窗口大小为k,有n-k+1个窗口
	mq := &MonoQueue{
		list.New(),
	}
	// 初始第一个窗口,找到第一个窗口的最大值
	for i := 0; i < k; i++ {
		mq.offer(nums[i])
	}
	result[arrayIndex] = mq.peek()
	arrayIndex++
	for i := k; i < n; i++ {
		// 从队首滑出一个最左边的下标：i-k
		mq.poll(nums[i-k]) // 滑动窗口移除最前面元素,移除时检查是否在队列里
		// 从队尾滑入一个新的元素：i
		mq.offer(nums[i]) // 滑动窗口加入最后面的元素
		// 当前窗口里最大值即队列的队首,记录对应的最大值
		result[arrayIndex] = mq.peek() // arrayIndex最大到len-k+1
		arrayIndex++
	}
	return result
}

// 定义一个单调队列，内部基于数组切片
type MonoQueueWithArray struct {
	deque []int
}

func NewMyQueue() *MonoQueueWithArray {
	return &MonoQueueWithArray{make([]int, 0)}
}

// 单调递减队列,入队方法
// 如果offer的元素value大于队尾(入口）元素的数值，那么就将队列入口的元素弹出
// 直到offer元素的数值小于等于队列入口元素的数值为止,然后放入队尾
// 这样就保持了队列里的数值是单调从大到小的了。
func (this *MonoQueueWithArray) offer(value int) {
	for !this.Empty() && value > this.Back() {
		this.deque = this.deque[:len(this.deque)-1]
	}
	this.deque = append(this.deque, value)
}

// 从队列中移除指定的元素
func (this *MonoQueueWithArray) poll(value int) {
	if !this.Empty() && this.Front() == value {
		this.deque = this.deque[1:]
	}
}

// 获取单调队列中队首的元素,即队列中最大值
func (this *MonoQueueWithArray) peek() int {
	return this.Front()
}

// 队列是否为空
func (this *MonoQueueWithArray) Empty() bool {
	return len(this.deque) == 0
}

// 队尾元素
func (this *MonoQueueWithArray) Back() int {
	return this.deque[len(this.deque)-1]
}

// 队首元素
func (this *MonoQueueWithArray) Front() int {
	return this.deque[0]
}

// 基于单调递减队列来查找滑动窗口最大值
func MaxSlidingWindow2(nums []int, k int) []int {
	n := len(nums)                              // 数组长度
	result, arrayIndex := make([]int, n-k+1), 0 // 窗口大小为k,有n-k+1个窗口
	mq := NewMyQueue()
	// 初始第一个窗口,找到第一个窗口的最大值
	for i := 0; i < k; i++ {
		mq.offer(nums[i])
	}
	result[arrayIndex] = mq.peek()
	arrayIndex++
	for i := k; i < n; i++ {
		// 从队首滑出一个最左边的下标：i-k
		mq.poll(nums[i-k]) // 滑动窗口移除最前面元素,移除时检查是否在队列里
		// 从队尾滑入一个新的元素：i
		mq.offer(nums[i]) // 滑动窗口加入最后面的元素
		// 当前窗口里最大值即队列的队首,记录对应的最大值
		result[arrayIndex] = mq.peek() // arrayIndex最大到len-k+1
		arrayIndex++
	}
	return result
}
