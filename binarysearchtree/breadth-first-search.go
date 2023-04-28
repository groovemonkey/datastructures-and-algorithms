package binarysearchtree

import "fmt"

func (n *BSTNode) BreadthFirstSearch(root *BSTNode) *BSTNode {
	// TODO use a deque for this -- write one and then import it in this package!
	queue := make([]*BSTNode, 0)

	if root != nil {
		queue.append(root)
	}

	level := 0

	for {
		if !len(queue) > 0 {
			break
		}
		fmt.Println("bfs level: ", level)
		for i := 0; i < len(queue); i++ {
			var current *BSTNode
			current = queue.popleft()
			// print current node's value
			fmt.Println("current:", current.value)

			if current.left != nil {
				queue.append(current.left)
			}
			if current.right != nil {
				queue.append(current.right)
			}
		}
		level += 1
	}
	// what should this actually return?
	return &BSTNode{}
}
