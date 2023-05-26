package linklists

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveElements(head *ListNode, val int) *ListNode {
	var dummy = &ListNode{-1, head}
	// 找到被删除结点的前一个结点
	p := dummy
	for p.Next != nil {
		x := p.Next       // x结点
		if x.Val == val { // 匹配删除的结点
			p.Next = x.Next
			x = nil
			continue // 这里需要重新检查p
		}
		p = p.Next

	}
	return dummy.Next
}

func ReverseList(head *ListNode) *ListNode {
	return reverseList(head)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev *ListNode = nil
	var current *ListNode = head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func swapPairs(head *ListNode) *ListNode {
	// 借助一个虚拟的dummy结点作为哨兵结点
	var dummy *ListNode = &ListNode{0, head}
	var prev *ListNode = dummy
	var p *ListNode = prev.Next
	for p != nil && p.Next != nil {
		// 需要画图理解，
		q := p.Next
		// 第一步
		prev.Next = q
		// 第二步
		p.Next = q.Next
		// 第三步
		q.Next = p

		// 第四步:移动prev 到p 移动p到p.next
		prev = p
		p = p.Next
	}
	return dummy.Next
}

// 时间复杂度O(2N)
func RemoveKthFromEnd1(head *ListNode, k int) *ListNode {
	var dummy = &ListNode{0, head}
	p := dummy
	n := getListLength(dummy.Next)
	var find = n - k
	for find > 0 {
		p = p.Next
		find--
	}
	x := p.Next // x是要被删除的结点,即倒数第k
	p.Next = x.Next
	x.Next = nil // help GC
	return dummy.Next
}

// 获取链表长度
func getListLength(p *ListNode) int {
	length := 0
	for p != nil {
		length++
		p = p.Next
	}
	return length
}

// 时间复杂度O(1)
func removeKthFromEnd2(head *ListNode, k int) *ListNode {
	var dummy = &ListNode{0, head}
	p1, p2 := dummy, dummy
	// p1先移动k个结点
	for step := 0; step < k; step++ {
		p1 = p1.Next
	}
	// p1,p2同时移动,直到p1到达末尾（非NULL结点）
	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	x := p2.Next
	p2.Next = x.Next
	x.Next = nil // help GC
	return dummy.Next
}

// 查找链表中间结点
// 采用快慢双指针,快指针一次移动2个结点(走2步),慢指针一次移动1个结点(走1步).当快指针到达末尾时,慢指针在中间结点位置
// 什么叫一次移动2个结点:即如果当前节点p且p.next不为空,直接走到p.next.next结点
// 如何定义到达末尾,即到达尾部为NULL的结点
func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// 链表环检查
// 采用快慢指针, 快指针一次移动2个结点(走2步),慢指针一次移动1个结点(走1步)
// 什么叫一次移动2个结点:即如果当前节点p且p.next不为空,直接走到p.next.next结点
// 如果链表存在环,快指针走到若干圈后,一定能碰上慢指针
// 龟兔赛跑算法
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false // 不存在环直接退出
}

// O(N)时间复杂度,O(N)空间复杂度
func getIntersectionNode_hash(headA, headB *ListNode) *ListNode {
	// 定义哈希表,默认容量
	var traversed = make(map[*ListNode]bool)
	// 将链表A的所有结点放入哈希表中
	for p := headA; p != nil; p = p.Next {
		traversed[p] = true
	}
	// 遍历链表B,检查是否存在
	for p := headB; p != nil; p = p.Next {
		v := traversed[p]
		if v {
			return p
		}
	}
	return nil
}

// O(N)时间复杂度，O(1)空间复杂度
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	// 假设p1代表长一点的链表
	var p1, p2 *ListNode
	lenA, lenB := getListLength(headA), getListLength(headB)
	x := 0
	if lenA >= lenB {
		x = lenA - lenB
		p1, p2 = headA, headB
	} else {
		x = lenB - lenA
		p1, p2 = headB, headA
	}
	// p1先移动x个结点
	for x > 0 {
		p1 = p1.Next
		x--
	}
	for p1 != nil && p2 != nil {
		if p1 == p2 {
			return p1
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return nil
}

func detectCycle1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 快慢指针：fast走2步,slow走1步
	// 如果存在环,fast slow会相遇,slow为相遇的结点
	// 此时让fast重新回到head,然后开始走,如果与slow再次相遇,此时相遇的结点为环入口
	// 如果不存在环,fast 不会等于slow, 返回null
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
