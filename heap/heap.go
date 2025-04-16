package heap

// Comparator defines a function type that compares two elements.
// It should return:
//   - negative if a < b
//   - zero if a == b
//   - positive if a > b
type Comparator[T any] func(a, b T) int

// PriorityQueue is a generic heap that can be configured with a custom comparator.
type PriorityQueue[T any] struct {
	data       []T
	comparator Comparator[T]
	capacity   int // maximum capacity; if <= 0, the queue is unbounded
}

// New creates a new PriorityQueue with the given comparator.
func New[T any](comparator Comparator[T]) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		data:       make([]T, 0),
		comparator: comparator,
		capacity:   0,
	}
}

// NewWithCapacity creates a new bounded PriorityQueue with the given comparator and capacity.
func NewWithCapacity[T any](comparator Comparator[T], capacity int) *PriorityQueue[T] {
	if capacity <= 0 {
		panic("heap: capacity must be positive")
	}
	return &PriorityQueue[T]{
		data:       make([]T, 0, capacity),
		comparator: comparator,
		capacity:   capacity,
	}
}

// NewOrdered creates a new min-heap for ordered types (supports <, > operators).
func NewOrdered[T Ordered]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		data: make([]T, 0),
		comparator: func(a, b T) int {
			if a < b {
				return -1
			}
			if a > b {
				return 1
			}
			return 0
		},
	}
}

// NewOrderedMax creates a new max-heap for ordered types.
func NewOrderedMax[T Ordered]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		data: make([]T, 0),
		comparator: func(a, b T) int {
			if a > b {
				return -1
			}
			if a < b {
				return 1
			}
			return 0
		},
	}
}

// Ordered defines types that support the built-in comparison operators.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Push inserts an element into the queue.
// If the queue is bounded and full, it removes the lowest priority element first.
func (pq *PriorityQueue[T]) Push(x T) bool {
	if pq.capacity > 0 && len(pq.data) >= pq.capacity {
		// For bounded queue, check if new element has higher priority than current min
		if pq.comparator(x, pq.data[0]) <= 0 {
			return false // New element has lower priority, don't add it
		}
		pq.Pop() // Remove lowest priority element to make space
	}
	pq.data = append(pq.data, x)
	pq.bubbleUp(len(pq.data) - 1)
	return true
}

// Pop removes and returns the highest priority element from the queue.
func (pq *PriorityQueue[T]) Pop() (T, bool) {
	if len(pq.data) == 0 {
		var zero T
		return zero, false
	}
	top := pq.data[0]
	last := len(pq.data) - 1
	pq.data[0] = pq.data[last]
	pq.data = pq.data[:last]
	pq.bubbleDown(0)
	return top, true
}

// Peek returns the highest priority element without removing it.
func (pq *PriorityQueue[T]) Peek() (T, bool) {
	if len(pq.data) == 0 {
		var zero T
		return zero, false
	}
	return pq.data[0], true
}

// Len returns the number of elements in the queue.
func (pq *PriorityQueue[T]) Len() int {
	return len(pq.data)
}

// IsEmpty returns true if the queue is empty.
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return len(pq.data) == 0
}

// Clear removes all elements from the queue.
func (pq *PriorityQueue[T]) Clear() {
	pq.data = pq.data[:0]
}

// Contains checks if the queue contains the given element.
// Note: This is an O(n) operation.
func (pq *PriorityQueue[T]) Contains(x T, equals func(a, b T) bool) bool {
	for _, item := range pq.data {
		if equals(item, x) {
			return true
		}
	}
	return false
}

// Remove removes the first occurrence of the specified element from the queue.
// Returns true if the element was found and removed.
// Note: This is an O(n) operation.
func (pq *PriorityQueue[T]) Remove(x T, equals func(a, b T) bool) bool {
	for i, item := range pq.data {
		if equals(item, x) {
			pq.removeAt(i)
			return true
		}
	}
	return false
}

// removeAt removes the element at the specified position.
func (pq *PriorityQueue[T]) removeAt(i int) {
	n := len(pq.data) - 1
	if n != i {
		pq.data[i] = pq.data[n]
		pq.data = pq.data[:n]
		if !pq.bubbleDown(i) {
			pq.bubbleUp(i)
		}
	} else {
		pq.data = pq.data[:n]
	}
}

// bubbleUp restores the heap property from index i upward.
func (pq *PriorityQueue[T]) bubbleUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if pq.comparator(pq.data[i], pq.data[parent]) >= 0 {
			break
		}
		pq.data[i], pq.data[parent] = pq.data[parent], pq.data[i]
		i = parent
	}
}

// bubbleDown restores the heap property from index i downward.
// Returns true if any swaps occurred.
func (pq *PriorityQueue[T]) bubbleDown(i int) bool {
	n := len(pq.data)
	swapped := false
	for {
		smallest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < n && pq.comparator(pq.data[left], pq.data[smallest]) < 0 {
			smallest = left
		}
		if right < n && pq.comparator(pq.data[right], pq.data[smallest]) < 0 {
			smallest = right
		}
		if smallest == i {
			break
		}
		pq.data[i], pq.data[smallest] = pq.data[smallest], pq.data[i]
		i = smallest
		swapped = true
	}
	return swapped
}
