package priorityqueue

import (
	"errors"

	"github.com/groovemonkey/datastructures-and-algorithms/binarysearchtree"
)

type PriorityQueueBinarySearchTree struct {
	Items   binarysearchtree.BinarySearchTree
	MinNode *binarysearchtree.BSTNode
}

func (pq *PriorityQueueBinarySearchTree) Insert(val int) {
	newNode := pq.Items.Insert(val)

	// Update minimum node if none exists, or if the new val is less
	if pq.MinNode == nil || val < pq.MinNode.Value {
		pq.MinNode = newNode
	}
}

func (pq *PriorityQueueBinarySearchTree) Min() (int, error) {
	if pq.MinNode != nil {
		return pq.MinNode.Value, nil
	}
	return 0, errors.New("no values in Priority Queue")
}

func (pq *PriorityQueueBinarySearchTree) DeleteMin() {
	pq.Items.Delete(pq.MinNode)

	// Find new minimum value
	pq.MinNode = pq.Items.Min()
}

// This is just a convenience wrapper around FindMin() && DeleteMin()
func (pq *PriorityQueueBinarySearchTree) Pop() int {
	val := pq.MinNode.Value
	pq.DeleteMin()
	return val
}
