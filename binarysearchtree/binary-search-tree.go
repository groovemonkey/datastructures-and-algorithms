package binarysearchtree

import "fmt"

type BinarySearchTree struct {
	root *BSTNode
}

func (b *BinarySearchTree) Search(value int) *BSTNode {
	if b.root == nil {
		return nil
	}
	return b.root.search(value)
}

func (b *BinarySearchTree) Contains(value int) bool {
	if b.root == nil {
		return false
	}
	// Perform a normal node search, but return a bool instead
	foundNode := b.root.search(value)
	return foundNode != nil
}

func (b *BinarySearchTree) Traverse() {
	f := func(n *BSTNode) {
		fmt.Println(n.Value)
	}
	b.root.traverse(f)
}

func (tree *BinarySearchTree) Insert(value int) *BSTNode {
	var y *BSTNode = nil
	var temp *BSTNode = tree.root

	n := &BSTNode{
		Value:  value,
		left:   nil,
		right:  nil,
		parent: nil,
	}

	// Search down the tree until we come to a leaf node
	for temp != nil {
		y = temp
		if n.Less(temp.Value) {
			temp = temp.left
		} else {
			temp = temp.right
		}
	}
	n.parent = y

	if y == nil {
		tree.root = n
	} else if n.Value < y.Value {
		y.left = n
	} else {
		y.right = n
	}

	return n
}

// Min() returns the node containing the minimum value in the entire tree
func (tree *BinarySearchTree) Min() *BSTNode {
	min := tree.root

	// Just scoot left until we're at the leftmost node
	for min.left != nil {
		min = min.left
	}
	return min
}

// Max() returns the node containing the maximum value in the entire tree
func (tree *BinarySearchTree) Max() *BSTNode {
	max := tree.root

	// Just scoot right until we're at the rightmost node
	for max.right != nil {
		max = max.right
	}
	return max
}

type BSTNode struct {
	// Whatever type you need for the value
	Value  int
	left   *BSTNode
	right  *BSTNode
	parent *BSTNode
}

// TODO unfinished -- mimic sort's functionality - .less()
// is this Node's value less than the comparison value parameter?
func (n *BSTNode) Less(val int) bool {
	return n.Value < val
}

// SubtreeMin() returns the node containing the minimum value in the subtree under n
func (n *BSTNode) SubtreeMin() *BSTNode {
	min := n

	// Just scoot left until we're at the leftmost node
	for n.left != nil {
		n = n.left
		min = n
	}
	return min
}

// SubtreeMax() returns the node containing the maximum value in the subtree under n
func (n *BSTNode) SubtreeMax() *BSTNode {
	max := n

	// Just scoot right until we're at the rightmost node
	for n.right != nil {
		n = n.right
		max = n
	}
	return max
}

// TODO how to construct search when a bool MUST be returned?
func (n *BSTNode) search(value int) *BSTNode {
	if n.Value == value {
		return n
	}
	if n.Less(value) {
		if n.right == nil {
			return nil
		} else {
			return n.right.search(value)
		}
	} else {
		if n.left == nil {
			return nil
		} else {
			return n.left.search(value)
		}
	}
}

// Recursive in-order tree traversal, applies a function
func (n *BSTNode) traverse(f func(*BSTNode)) {
	if n != nil {
		n.left.traverse(f)
		f(n)
		n.right.traverse(f)
	}
}

func (t *BinarySearchTree) transplant(n1 *BSTNode, n2 *BSTNode) {
	// is n1 the root?
	if n1.parent == nil {
		t.root = n2

		// n1 is left child
	} else if n1 == n1.parent.left {
		n1.parent.left = n2
		// n1 is right child
	} else {
		n1.parent.right = n2
	}

	if n2 != nil {
		n2.parent = n1.parent
	}
}

func (t *BinarySearchTree) Delete(n *BSTNode) {
	// Handle cases where we have nil children
	if n.left == nil {
		t.transplant(n, n.right)
	} else if n.right == nil {
		t.transplant(n, n.left)
	} else {
		y := n.right.SubtreeMin()
		if y.parent != n {
			t.transplant(n, n.right)
			y.right = n.right
			y.right.parent = y
		}
		t.transplant(n, y)
		y.left = n.left
		y.left.parent = y
	}
}
