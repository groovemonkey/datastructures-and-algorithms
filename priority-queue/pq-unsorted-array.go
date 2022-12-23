package priorityqueue

type PriorityQueueUnsortedArray struct {
	Items  []int
	MinIdx int
}

func (pq *PriorityQueueUnsortedArray) Insert(val int) {
	pq.Items = append(pq.Items, val)
	// Update minimum index
	if val < pq.Items[pq.MinIdx] {
		pq.MinIdx = len(pq.Items) - 1
	}
}

func (pq *PriorityQueueUnsortedArray) FindMin() int {
	return pq.Items[pq.MinIdx]
}

func (pq *PriorityQueueUnsortedArray) DeleteMin() {
	pq.Items = append(pq.Items[:pq.MinIdx], pq.Items[pq.MinIdx+1:]...)

	// Find new minimum element
	minIdx := 0
	for i := range pq.Items {
		if pq.Items[i] < pq.Items[minIdx] {
			minIdx = i
		}
	}
	pq.MinIdx = minIdx
}

// This is just a convenience wrapper around FindMin() && DeleteMin()
func (pq *PriorityQueueUnsortedArray) Pop() int {
	val := pq.FindMin()
	pq.DeleteMin()
	return val
}
