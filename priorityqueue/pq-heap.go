package priorityqueue

//
//            1
//      2          3
//   4    5     6      7
// 8 9  10 11 12 13 14 15

type Heap struct {
	data []HeapItem
}

type HeapItem struct {
	Key  int
	data any
}

func (h *Heap) Insert(item HeapItem) {
	h.data = append(h.data, item)
	i := len(h.data) - 1
	h.maxHeapifyUp(i)
}

func (h *Heap) maxHeapifyUp(i int) {
	// If we're at the root, stop
	if i == 0 {
		return
	}

	item := h.data[i]
	parentIdx := parent(i)

	if item.Key > h.data[parentIdx].Key {
		// swap parent and child
		h.swap(parentIdx, i)
		h.maxHeapifyUp(parentIdx)
	}
}

func (h *Heap) maxIndex(i, j int) int {
	if h.data[i].Key > h.data[j].Key {
		return i
	}
	return j
}

func (h *Heap) maxHeapifyDown(i int) {
	// if i is a leaf, we're done!
	// get child indexes
	childIdx1 := i * 2
	childIdx2 := (i * 2) + 1

	// are either of them out of bounds?
	lastIndex := len(h.data) - 1
	if childIdx2 > lastIndex {
		if childIdx1 > lastIndex {
			// no valid children
			return
		}
		// child1 is valid
		h.swap(i, childIdx1)
		return
	}

	// both children are valid; swap with the bigger one
	maxChildIdx := h.maxIndex(childIdx1, childIdx2)
	h.swap(i, maxChildIdx)
	h.maxHeapifyDown(maxChildIdx)
}

// Delete max is the cool "trick" of how heapsort works on an input array --
// because DeleteMax moves the biggest item to the end of the array, you can sort IN PLACE
// by consuming an unsorted array with repeated inserts, and then shrinking the heap with every
// deleted Max item -- at the end the underlying array is sorted, and the heap is size 0 again.
// TODO I'm not sure how to implement this in Go, actually. Decrement slice.Length manually?
func (h *Heap) DeleteMax() {
	// swap first and last item
	h.swap(0, len(h.data)-1)

	// delete last item
	h.data = h.data[:len(h.data)-1]

	// restore heap property, if necessary
	h.maxHeapifyDown(0)
}

func parent(i int) int {
	return i / 2
}

func (h *Heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
