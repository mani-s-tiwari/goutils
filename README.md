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

```

## 🔎 Quick Reference (Python → Go)

| Python              | Go (goutils)                  | Example (Go)                                  |
|---------------------|-------------------------------|-----------------------------------------------|
| `heapq.heappush`    | `h.Heappush(heapq.Item{...})` | `h.Heappush(heapq.Item{Key: 5, Val: 100})`    |
| `heapq.heappop`     | `h.Heappop()`                 | `item := h.Heappop(); fmt.Println(item.Val)`  |
| `collections.deque` | `deque.New()`                 | `d := deque.New()`                            |
| `deque.append`      | `PushBack(x)`                 | `d.PushBack(10)`                              |
| `deque.appendleft`  | `PushFront(x)`                | `d.PushFront(5)`                              |
| `deque.pop`         | `PopBack()`                   | `x := d.PopBack()`                            |
| `deque.popleft`     | `PopFront()`                  | `x := d.PopFront()`                           |
| DSU / Union-Find    | `dsu.New(n)`                  | `uf := dsu.New(5); uf.Union(0,1)`             |
| Trie                | `trie.New()`                  | `t := trie.New(); t.Insert("go")`             |


🔹 Included Data Structures (with Code)

Example
## 🚀 Full Example Usage

Here’s how you can import and use all utilities:

```go
package main

import (
	"fmt"

	"github.com/mani-s-tiwari/goutils/heapq"
	"github.com/mani-s-tiwari/goutils/deque"
	"github.com/mani-s-tiwari/goutils/dsu"
	"github.com/mani-s-tiwari/goutils/trie"
)

func main() {
	// --- Heap Example ---
	h := heapq.NewMaxHeap()
	h.Heappush(heapq.Item{Key: 5, Val: 100})
	h.Heappush(heapq.Item{Key: 9, Val: 200})
	h.Heappush(heapq.Item{Key: 1, Val: 300})
	fmt.Println("Heap pop:", h.Heappop()) // {Key:9, Val:200}

	// --- Deque Example ---
	d := deque.New()
	d.PushBack(1)
	d.PushFront(0)
	fmt.Println("Deque PopFront:", d.PopFront()) // 0
	fmt.Println("Deque PopBack:", d.PopBack())   // 1

	// --- DSU Example ---
	uf := dsu.New(5)
	uf.Union(0, 1)
	uf.Union(1, 2)
	fmt.Println("DSU Connected(0,2):", uf.Find(0) == uf.Find(2)) // true
	fmt.Println("DSU Connected(3,4):", uf.Find(3) == uf.Find(4)) // false

	// --- Trie Example ---
	t := trie.New()
	t.Insert("go")
	t.Insert("goal")
	fmt.Println("Trie Search(go):", t.Search("go"))        // true
	fmt.Println("Trie Search(god):", t.Search("god"))      // false
	fmt.Println("Trie StartsWith(go):", t.StartsWith("go")) // true
}
```

Go is great for backend systems but lacks built-in data structures for DSA/competitive programming.
This library fills the gaps and makes Go feel closer to Python STL + C++ STL for problem-solving.


---
