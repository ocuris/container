package heap

// Ordered defines types that support the built-in comparison operators.
// The tilde (~) allows types with underlying type int or float64.
type Ordered interface {
	~int | ~float64
}

// pq is a generic heap structure that can operate as either a min-heap or a max-heap.
// The 'min' field determines its behavior:
//   - true: min-heap (the smallest element is at the top)
//   - false: max-heap (the largest element is at the top)
type pq[T Ordered] struct {
	data []T
	cap  int  // maximum capacity; if non-zero, the heap is bounded
	min  bool // if true: min-heap; if false: max-heap
}

const defaultCap = 0

// New creates a new default min-heap.
func New[T Ordered]() *pq[T] {
	return &pq[T]{
		data: make([]T, 0),
		cap:  0,
		min:  true,
	}
}

// NewMin creates a new min-heap with an optional capacity.
func NewMin[T Ordered](size ...int) *pq[T] {
	actualCap := defaultCap
	if len(size) > 0 {
		if size[0] <= 0 {
			panic("heap: invalid capacity")
		}
		actualCap = size[0]
	}
	return &pq[T]{
		data: make([]T, 0, actualCap),
		cap:  actualCap,
		min:  true,
	}
}

// NewMax creates a new max-heap with an optional capacity.
func NewMax[T Ordered](size ...int) *pq[T] {
	actualCap := defaultCap
	if len(size) > 0 {
		if size[0] <= 0 {
			panic("heap: invalid capacity")
		}
		actualCap = size[0]
	}
	return &pq[T]{
		data: make([]T, 0, actualCap),
		cap:  actualCap,
		min:  false,
	}
}

// compare is an internal helper that returns true if a has higher priority than b.
// For min-heaps, lower values have higher priority (a < b).
// For max-heaps, higher values have higher priority (a > b).
func (h *pq[T]) compare(a, b T) bool {
	if h.min {
		return a < b
	}
	return a > b
}

// Push inserts an element into the heap.
// If the heap is bounded (cap != 0) and full, it removes the top element first.
func (h *pq[T]) Push(x T) {
	if h.cap != 0 && len(h.data) == h.cap {
		h.Pop() // Bounded behavior: remove top element to make space.
	}
	h.data = append(h.data, x)
	h.bubbleUp(len(h.data) - 1)
}

// Pop removes and returns the top element from the heap.
// For a min-heap, this is the smallest element; for a max-heap, the largest.
func (h *pq[T]) Pop() T {
	if len(h.data) == 0 {
		panic("heap: pop from empty heap")
	}
	top := h.data[0]
	last := len(h.data) - 1
	// Move the last element to the root and then shorten the slice.
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	h.bubbleDown(0)
	return top
}

// Len returns the current number of elements in the heap.
func (h *pq[T]) Len() int {
	return len(h.data)
}

// bubbleUp restores the heap property from index i upward.
func (h *pq[T]) bubbleUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if !h.compare(h.data[i], h.data[parent]) {
			break
		}
		h.data[i], h.data[parent] = h.data[parent], h.data[i]
		i = parent
	}
}

// bubbleDown restores the heap property from index i downward.
func (h *pq[T]) bubbleDown(i int) {
	n := len(h.data)
	for {
		best := i
		left := 2*i + 1
		right := 2*i + 2
		if left < n && h.compare(h.data[left], h.data[best]) {
			best = left
		}
		if right < n && h.compare(h.data[right], h.data[best]) {
			best = right
		}
		if best == i {
			break
		}
		h.data[i], h.data[best] = h.data[best], h.data[i]
		i = best
	}
}
