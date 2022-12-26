package splaytree

import "fmt"

type SplayTree struct {
	root *SplayNode
}

type SplayNode struct {
	// Whatever type you need for the data
	Data   string
	left   *SplayNode
	right  *SplayNode
	parent *SplayNode
}

// TODO
// Make this whatever you want it to be!
// type DataItem struct {
// 	Data     string
// 	attr      string
// 	otherattr int
// }

func (node *SplayNode) searchTree(key string) *SplayNode {
	if node == nil || node.Data == key {
		return node
	}
	if key < node.Data {
		return node.left.searchTree(key)
	} else {
		return node.right.searchTree(key)
	}
}

func (tree *SplayTree) deleteNode(node *SplayNode, key string) {
	var x, t, s *SplayNode
	for node != nil {
		if node.Data == key {
			x = node
		}
		if node.Data < key {
			node = node.right
		} else {
			node = node.left
		}
	}
	// No such node; return silently
	if x == nil {
		return
	}

	tree.splay(x)

	// Split
	if x.right != nil {
		t = x.right
		t.parent = nil
	} else {
		t = nil
	}
	s = x
	s.right = nil
	x = nil

	// Join
	if s.left != nil {
		s.left.parent = nil
	}

	tree.root = tree.join(s.left, t)
	s = nil // TODO what does this do?
}

func (tree *SplayTree) leftRotate(node *SplayNode) {
	y := node.right
	node.right = y.left
	if y.left != nil {
		y.left.parent = node
	}

	y.parent = node.parent
	if node.parent == nil {
		tree.root = y
	} else if node == node.parent.left {
		node.parent.left = y
	} else {
		node.parent.right = y
	}
	y.left = node
	node.parent = y
}

func (tree *SplayTree) rightRotate(node *SplayNode) {
	y := node.left
	node.left = y.right
	if y.right != nil {
		y.right.parent = node
	}

	y.parent = node.parent
	if node.parent == nil {
		tree.root = y
	} else if node == node.parent.right {
		node.parent.right = y
	} else {
		node.parent.left = y
	}
	y.right = node
	node.parent = y
}

// Move node to the root of the tree
// TODO use switch statement instead of else-elseif-else
func (tree *SplayTree) splay(node *SplayNode) {
	for node.parent != nil {

		// final rotation
		if node.parent.parent == nil {
			if node == node.parent.left {
				// zig
				tree.rightRotate(node.parent)
			} else {
				// zag
				tree.leftRotate(node.parent)
			}
		} else if node == node.parent.left && node.parent == node.parent.parent.left {
			// zig-zig
			tree.rightRotate(node.parent.parent)
			tree.rightRotate(node.parent)
		} else if node == node.parent.right && node.parent == node.parent.parent.right {
			// zag-zag
			tree.leftRotate(node.parent.parent)
			tree.leftRotate(node.parent)
		} else if node == node.parent.right && node.parent == node.parent.parent.left {
			tree.leftRotate(node.parent)
			tree.rightRotate(node.parent)
		} else {
			// zag-zig
			tree.rightRotate(node.parent)
			tree.leftRotate(node.parent)
		}
	}
}

func (tree *SplayTree) join(s *SplayNode, t *SplayNode) *SplayNode {
	if s == nil {
		return t
	}
	if t == nil {
		return s
	}

	x := s.Max()
	tree.splay(x)
	x.right = t
	t.parent = x
	return x
}

func (node *SplayNode) inorder() {
	if node != nil {
		node.left.inorder()
		fmt.Println(node.Data)
		node.right.inorder()
	}
}

// search tree and return node with key
// splays the found node
func (tree *SplayTree) Search(key string) *SplayNode {
	node := tree.root.searchTree(key)
	if node != nil {
		tree.splay(node)
	}
	return node
}

// Find and return the node with the minimum value
func (node *SplayNode) Min() *SplayNode {
	for node.left != nil {
		node = node.left
	}
	return node
}

// Find and return the node with the maximum value
func (node *SplayNode) Max() *SplayNode {
	for node.right != nil {
		node = node.right
	}
	return node
}

func (node *SplayNode) Successor() *SplayNode {
	if node.right != nil {
		return node.right.Min()
	}
	// Leftmost ancestor of this node
	p := node.parent
	for p != nil && node == p.right {
		node = p
		p = p.parent
	}
	return p
}

func (node *SplayNode) Predecessor() *SplayNode {
	if node.left != nil {
		return node.left.Min()
	}
	p := node.parent
	for p != nil && node == p.left {
		node = p
	}
	p = p.parent

	return p
}

func (tree *SplayTree) Insert(key string) {
	newNode := &SplayNode{
		Data: key,
	}
	var y *SplayNode
	var x = tree.root

	for x != nil {
		y = x
		if newNode.Data < x.Data {
			x = x.left
		} else {
			x = x.right
		}
	}

	newNode.parent = y
	if y == nil {
		tree.root = newNode
	} else if newNode.Data < y.Data {
		y.left = newNode
	} else {
		y.right = newNode
	}
	// splay the Node
	tree.splay(newNode)
}
