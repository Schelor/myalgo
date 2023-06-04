package heaps

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Peek() any {
	return (*h)[0]
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func Demo() string {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("peek minimum: %d\n", h.Peek()) // 1
	fmt.Printf("peek minimum: %d\n", (*h)[0])  // 1
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h)) // 1,2,3,5
	}
	return "IntHeap"
}

// 通过最小堆来实现Top K
func TopKFrequent(nums []int, k int) []int {
	// 先对元素计数
	freq := make(map[int]int, len(nums))
	for _, v := range nums {
		freq[v]++
	}
	return []int{}
}
