package heap

import (
	"sort"
)

// Item represents a single element in the priority queue
type Item struct {
	Value    any
	Priority int
}

// PriorityQueue represents a slice-based max-heap priority queue
type PriorityQueue []Item

// Push adds an item to the priority queue and maintains heap property
func (pq *PriorityQueue) Push(item Item) {
	*pq = append(*pq, item)
	pq.heapifyUp(len(*pq) - 1)
}

// Pop removes and returns the highest priority item
func (pq *PriorityQueue) Pop() Item {
	n := len(*pq)
	if n == 0 {
		return Item{}
	}
	top := (*pq)[0]
	(*pq)[0] = (*pq)[n-1]
	*pq = (*pq)[:n-1]
	pq.heapifyDown(0)
	return top
}

// Peek returns the highest priority item without removing it
func (pq PriorityQueue) Peek() Item {
	if len(pq) == 0 {
		return Item{}
	}
	return pq[0]
}

// Len returns the number of items in the queue
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// heapifyUp maintains the heap property after Push
func (pq PriorityQueue) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if pq[index].Priority <= pq[parent].Priority {
			break
		}
		pq[index], pq[parent] = pq[parent], pq[index]
		index = parent
	}
}

// heapifyDown maintains the heap property after Pop
func (pq PriorityQueue) heapifyDown(index int) {
	last := len(pq) - 1
	for {
		left := 2*index + 1
		right := 2*index + 2
		largest := index

		if left <= last && pq[left].Priority > pq[largest].Priority {
			largest = left
		}
		if right <= last && pq[right].Priority > pq[largest].Priority {
			largest = right
		}
		if largest == index {
			break
		}
		pq[index], pq[largest] = pq[largest], pq[index]
		index = largest
	}
}

// Sort sorts the priority queue based on priority (highest first)
func (pq *PriorityQueue) Sort() {
	sort.SliceStable(*pq, func(i, j int) bool {
		return (*pq)[i].Priority > (*pq)[j].Priority
	})
}

// Update modifies an item and re-heapifies
func (pq *PriorityQueue) Update(index int, newItem Item) {
	if index < 0 || index >= len(*pq) {
		return
	}
	oldPriority := (*pq)[index].Priority
	(*pq)[index] = newItem
	if newItem.Priority > oldPriority {
		pq.heapifyUp(index)
	} else {
		pq.heapifyDown(index)
	}
}

// Remove deletes an item at a given index
func (pq *PriorityQueue) Remove(index int) {
	n := len(*pq)
	if index < 0 || index >= n {
		return
	}
	(*pq)[index] = (*pq)[n-1]
	*pq = (*pq)[:n-1]
	pq.heapifyDown(index)
}
