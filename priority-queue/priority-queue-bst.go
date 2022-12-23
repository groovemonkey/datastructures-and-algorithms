package priorityqueue

import (
	"github.com/groovemonkey/datastructures-and-algorithms/binarysearchtree"
)

type PriorityQueueBinarySearchTree struct {
	Items  binarysearchtree.BinarySearchTree
	MinVal int
}

func (pq *PriorityQueueBinarySearchTree) Insert(val int) {
	pq.Items.Insert(val)

	// Update minimum index
	if val < pq.MinVal {
		pq.MinVal = val
	}
}

func (pq *PriorityQueueBinarySearchTree) FindMin() int {
	return pq.MinVal
}

func (pq *PriorityQueueBinarySearchTree) DeleteMin() {
	pq.Items.Delete(pq.MinVal)

	// Find new minimum value
	pq.MinVal = pq.Items.root.FindMin().Value
}

// This is just a convenience wrapper around FindMin() && DeleteMin()
func (pq *PriorityQueueBinarySearchTree) Pop() int {
	val := pq.MinVal
	pq.DeleteMin()
	return val
}
