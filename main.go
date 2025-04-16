package main

import (
	"fmt"

	"github.com/ocuris/container/heap"
)

func main() {

	pq := heap.NewMax[int](3)
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
}
