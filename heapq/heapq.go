package heapq

import (
	"container/heap"
)

type Item struct {
	Key int // priority
	Val int // payload
}

type Heap struct {
	data []Item
	min  bool
}

func (h Heap) Len() int { return len(h.data) }
func (h Heap) Less(i, j int) bool {
	if h.min {
		return h.data[i].Key < h.data[j].Key
	}
	return h.data[i].Key > h.data[j].Key
}
func (h Heap) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *Heap) Push(x interface{}) { h.data = append(h.data, x.(Item)) }
func (h *Heap) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[:n-1]
	return x
}

func NewMinHeap() *Heap { h := &Heap{min: true}; heap.Init(h); return h }
func NewMaxHeap() *Heap { h := &Heap{min: false}; heap.Init(h); return h }
func (h *Heap) Heappush(item Item) { heap.Push(h, item) }
func (h *Heap) Heappop() Item      { return heap.Pop(h).(Item) }
func (h *Heap) Len2() int          { return len(h.data) }
