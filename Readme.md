
# Ocuris Container 📦

A lightweight, idiomatic Go library offering core data structures inspired by Python's built-in collections. Designed for maximum control and learnability, this library is built without relying on Go's standard `container/heap` or `container/list`, enabling clear, readable, and customizable code.

---

## ✨ Features

- ✅ **MinHeap / MaxHeap**
- ✅ **Priority Queue**
- ✅ **Stack**
- ✅ **Queue (FIFO)**
- ✅ **Linked List**
- ✅ **Deque (Double-ended Queue)**
- ✅ **HashSet (like Python `set`)**
- ✅ **Graph (Adjacency List)**
- ✅ **Tree (Binary Trees)**
- ✅ **Tree Traversals**
- ✅ **Utility Functions** (Sorting, Searching, etc.)
- ✅ **While-loop Style Iteration** for containers

---

## 📦 Installation

To install the Ocuris Container library, run the following command:

```bash
go get github.com/ocuris/container
```

---

## 📚 Usage Examples

### 🔺 MinHeap / MaxHeap

The `MinHeap` and `MaxHeap` allow you to maintain a priority queue where elements are stored based on their natural order (smallest or largest).

```go
import "github.com/ocuris/container/heap"

// Create a MinHeap
pq := heap.MinHeap{}

// Push items into the heap
pq.Push(heap.Item{Value: 20})
pq.Push(heap.Item{Value: 15})
pq.Push(heap.Item{Value: 30})

// Peek at the smallest item without removing it
item := pq.Peek()
fmt.Println("Peek:", item.Value) // => 15

// Pop the smallest item (and remove it)
item = pq.Pop()
fmt.Println("Pop:", item.Value) // => 15

// Peek again after popping
item = pq.Peek()
fmt.Println("Peek after Pop:", item.Value) // => 20
```

### 🚦 Priority Queue

A priority queue can be implemented using either `MinHeap` or `MaxHeap` for custom behavior. Add elements to the queue, and they will be ordered by priority (based on their value).

### 📚 Stack

The stack implements the typical LIFO (Last In, First Out) behavior.

```go
import "github.com/ocuris/container/stack"

s := stack.New()
s.Push("go")
s.Push("lang")
fmt.Println(s.Pop()) // => lang
```

### 📦 Queue

The queue follows FIFO (First In, First Out) order.

```go
import "github.com/ocuris/container/queue"

q := queue.New()
q.Enqueue("first")
q.Enqueue("second")
fmt.Println(q.Dequeue()) // => first
```

### 🔁 Linked List

The linked list supports common operations such as append, prepend, and delete.

```go
import "github.com/ocuris/container/list"

ll := list.New[int]()
ll.Append(10)
ll.Append(20)
ll.Prepend(5)
ll.Delete(10)
```

### 🔁 Deque

The deque supports operations at both ends (push/pop from both front and back).

```go
import "github.com/ocuris/container/deque"

dq := deque.New[int]()
dq.PushFront(1)
dq.PushBack(2)
fmt.Println(dq.PopFront()) // => 1
```

### 🌳 Tree (Binary Tree)

A simple binary tree structure to represent hierarchical data.

```go
import "github.com/ocuris/container/tree"

root := tree.NewNode(1)
root.Left = tree.NewNode(2)
root.Right = tree.NewNode(3)
```

### 🧩 HashSet

The hash set allows you to store unique elements with efficient look-up operations.

```go
import "github.com/ocuris/container/hashset"

set := hashset.New[int]()
set.Add(1)
set.Add(2)
fmt.Println(set.Contains(2)) // true
fmt.Println(set.Contains(5)) // false
```

### 🔗 Graph

The graph supports basic adjacency list operations for creating and traversing graphs.

```go
import "github.com/ocuris/container/graph"

g := graph.New()
g.AddEdge("A", "B")
g.AddEdge("A", "C")
fmt.Println(g.Neighbors("A")) // => [B C]
```

### 🔁 While-Loop Style Iteration

Inspired by Python's `for` loop, this library supports a `while`-like iteration style to consume elements from containers.

```go
// Mimics Python-style while loops for consuming elements
for !q.IsEmpty() {
    val := q.Dequeue()
    fmt.Println(val)
}

for !s.IsEmpty() {
    val := s.Pop()
    fmt.Println(val)
}

for !pq.IsEmpty() {
    val := pq.Pop()
    fmt.Println(val)
}
```

---

## 🛠️ Goals

- Provide Python-like data structures for rapid prototyping in Go
- Focus on zero dependencies and native slices for full control
- Offer clean, readable code that is easy to learn and extend

---

## ✅ Coming Soon

- Trie Data Structure
- LRU Cache
- Thread-safe wrappers for concurrency
- Additional Sorting and Searching algorithms

---

## 🤝 Contributing

We welcome contributions! Feel free to open issues or pull requests to enhance the library. Contributions can include:

- Bug fixes
- New data structures
- Performance improvements
- Documentation enhancements

---

## 📄 License

MIT © [Ocuris](https://github.com/ocuris/container/blob/main/LICENSE)

---
