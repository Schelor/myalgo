package stackqueues

import "strconv"

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

func RemoveDuplicates(s string) string {
	i := len(s)
	for i > 0 {
		s = RemoveDuplicates0(s)
		i--
	}
	return s
}

func RemoveDuplicates0(s string) string {
	queue := make([]byte, 0)
	i := 0
	for i < len(s) {
		if j := i + 1; j < len(s) && s[i] == s[j] {
			i += 2
			continue
		}
		queue = append(queue, s[i])
		i++
	}
	return string(queue)
}

func RemoveDuplicates2(s string) string {
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
