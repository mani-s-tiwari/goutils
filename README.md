📄 README.md
# Go Utils (goutils)

A collection of **data structures and algorithms utilities in Go**, inspired by Python’s `heapq`, `deque`, and C++ STL.  
Useful for **DSA practice, LeetCode, and interview prep**.

---

## 📦 Installation

```bash
go get github.com/mani-s-tiwari/goutils


Then import what you need:

import "github.com/mani-s-tiwari/goutils/heapq"
import "github.com/mani-s-tiwari/goutils/deque"
import "github.com/mani-s-tiwari/goutils/dsu"
import "github.com/mani-s-tiwari/goutils/trie"

🔎 Quick Reference (Python → Go)
Python	Go (goutils)
heapq.heappush	h.Heappush(heapq.Item{...})
heapq.heappop	h.Heappop()
collections.deque	deque.New()
deque.append	PushBack(x)
deque.appendleft	PushFront(x)
deque.pop	PopBack()
deque.popleft	PopFront()
DSU / Union-Find	dsu.New(n)
Trie	trie.New()
🔹 Included Data Structures (with Code)
1. Heap (heapq)

heapq/heapq.go

package heapq

import (
	"container/heap"
)

type Item struct {
	Key int // priority (used for ordering)
	Val int // payload
}

type Heap struct {
	data []Item
	min  bool // true = min-heap, false = max-heap
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


Example

h := heapq.NewMaxHeap()
h.Heappush(heapq.Item{Key: 5, Val: 100})
h.Heappush(heapq.Item{Key: 9, Val: 200})
fmt.Println(h.Heappop()) // {Key:9, Val:200}

2. Deque (deque)

deque/deque.go

package deque

type Deque struct {
	data []int
}

func New() *Deque { return &Deque{data: []int{}} }
func (d *Deque) PushBack(x int)   { d.data = append(d.data, x) }
func (d *Deque) PushFront(x int)  { d.data = append([]int{x}, d.data...) }
func (d *Deque) PopBack() int     { x := d.data[len(d.data)-1]; d.data = d.data[:len(d.data)-1]; return x }
func (d *Deque) PopFront() int    { x := d.data[0]; d.data = d.data[1:]; return x }
func (d *Deque) Front() int       { return d.data[0] }
func (d *Deque) Back() int        { return d.data[len(d.data)-1] }
func (d *Deque) Empty() bool      { return len(d.data) == 0 }
func (d *Deque) Len() int         { return len(d.data) }


Example

d := deque.New()
d.PushBack(1)
d.PushFront(0)
fmt.Println(d.PopFront()) // 0

3. DSU (Disjoint Set Union)

dsu/dsu.go

package dsu

type DSU struct {
	parent []int
	rank   []int
}

func New(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &DSU{parent: parent, rank: rank}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x]) // path compression
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) bool {
	rx, ry := d.Find(x), d.Find(y)
	if rx == ry {
		return false
	}
	if d.rank[rx] < d.rank[ry] {
		d.parent[rx] = ry
	} else if d.rank[rx] > d.rank[ry] {
		d.parent[ry] = rx
	} else {
		d.parent[ry] = rx
		d.rank[rx]++
	}
	return true
}


Example

uf := dsu.New(5)
uf.Union(0, 1)
uf.Union(1, 2)
fmt.Println(uf.Find(0) == uf.Find(2)) // true

4. Trie (Prefix Tree)

trie/trie.go

package trie

type Node struct {
	children map[rune]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func New() *Trie {
	return &Trie{root: &Node{children: make(map[rune]*Node)}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if node.children[ch] == nil {
			node.children[ch] = &Node{children: make(map[rune]*Node)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return true
}


Example

t := trie.New()
t.Insert("go")
t.Insert("goal")

fmt.Println(t.Search("go"))     // true
fmt.Println(t.Search("god"))    // false
fmt.Println(t.StartsWith("go")) // true

🛠️ What’s Included

✅ heapq: MinHeap / MaxHeap (like Python’s heapq)

✅ deque: Double-ended queue (like collections.deque)

✅ dsu: Disjoint Set Union (Union-Find)

✅ trie: Prefix Tree for strings

🎯 Why?

Go is great for backend systems but lacks built-in data structures for DSA/competitive programming.
This library fills the gaps and makes Go feel closer to Python STL + C++ STL for problem-solving.


---

⚡ This README contains:  
- Intro + install  
- Quick reference table (Python → Go)  
- Full source code for each package (`heapq`, `deque`, `dsu`, `trie`)  
- Example snippets  

---

👉 Do you also want me to include a **section for running tests** (like `go test ./...`) so the repo looks professional?
