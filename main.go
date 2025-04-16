package main

import (
	"fmt"
	"strings"

	"github.com/ocuris/container/heap"
)

func main() {

	pq := heap.NewOrdered[int]()
	pq.Push(5)
	pq.Push(3)
	pq.Push(8)
	pq.Push(1)
	pq.Push(2)
	pq.Push(4)
	pq.Push(7)
	pq.Push(6)
	pq.Push(0)
	pq.Push(9)

	fmt.Println("Priority Queue:")
	for pq.Len() > 0 {
		fmt.Println(pq.Pop())
	}

	lengthComparator := func(a, b string) int {
		return len(a) - len(b) // Shorter strings have higher priority
	}

	pq1 := heap.New(lengthComparator)
	pq1.Push("go")
	pq1.Push("rust")
	pq1.Push("python")

	for pq1.Len() > 0 {
		val, _ := pq1.Pop()
		fmt.Println(val) // Output: go, rust, python
	}

	type Task struct {
		Priority int
		Name     string
	}

	pq2 := heap.New(func(a, b Task) int {
		if a.Priority != b.Priority {
			return a.Priority - b.Priority // Higher priority first
		}
		return strings.Compare(a.Name, b.Name)
	})
	pq2.Push(Task{2, "Write docs"})
	pq2.Push(Task{1, "Fix bug"})
	pq2.Push(Task{2, "Review code"})

	for pq2.Len() > 0 {
		task, _ := pq2.Pop()
		fmt.Printf("%d: %s\n", task.Priority, task.Name)
	}

}
