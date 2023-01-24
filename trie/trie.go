package trie

import "fmt"

type Trie struct {
	root *TrieNode
}

// func NewTrie() *Trie {
// 	return &Trie{
// 		root: NewTrieNode(),
// 	}
// }

type TrieNode struct {
	terminal bool // is this the end of a string?
	children map[rune]*TrieNode
}

// func NewTrieNode() *TrieNode {
// 	return &TrieNode{
// 		children: make(map[rune]*TrieNode),
// 	}
// }

func (trie *Trie) Insert(value string) bool {
	var createdNewNodes bool

	if trie.root == nil {
		trie.root = &TrieNode{
			children: make(map[rune]*TrieNode),
		}
	}

	// Start traversing at root
	node := trie.root
	lastIdx := len(value) - 1

	for i, r := range value {
		// rune doesn't exist yet
		if node.children[r] == nil {
			node.children[r] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
			createdNewNodes = true
		}
		node = node.children[r]
		if i == lastIdx && !node.terminal {
			node.terminal = true
		}
	}

	return createdNewNodes
}

// TODO this is not terminal-node aware
// (i.e. if CATTAILS has been inserted, Contains(CAT) is true even though the T node has terminal = false)
func (trie *Trie) Contains(value string) bool {
	if trie.root == nil {
		return false
	}
	node := trie.root
	for _, r := range value {
		if node.children[r] == nil {
			return false
		} else {
			node = node.children[r]
		}
	}
	// we've gone through the whole string
	return true
}

func depthFirst(node *TrieNode) {
	for r, n := range node.children {
		fmt.Println(string(r))
		if n.terminal {
			fmt.Println("END OF A WORD")
		}
		depthFirst(n)
	}
}

func (trie *Trie) DepthFirst() {
	if trie.root == nil {
		return
	}
	depthFirst(trie.root)
}
