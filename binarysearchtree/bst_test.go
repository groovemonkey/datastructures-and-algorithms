package binarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertAndSearchAndContains(t *testing.T) {
	tcs := []struct {
		name          string
		insert        []int
		expectPresent []int
		expectAbsent  []int
	}{
		{
			name:          "no insert",
			insert:        []int{},
			expectPresent: []int{},
		},
		{
			name:          "one insert",
			insert:        []int{5},
			expectPresent: []int{5},
			expectAbsent:  []int{1},
		},
		{
			name:          "a few inserts",
			insert:        []int{7, 2301231, 1},
			expectPresent: []int{2301231},
			expectAbsent:  []int{5, 43},
		},
	}

	for _, tc := range tcs {
		tree := BinarySearchTree{}
		for _, v := range tc.insert {
			tree.Insert(v)
		}

		// Test Search() and Contains() -- same underlying logic
		for _, v := range tc.expectPresent {
			found := tree.Search(v)
			assert.Equal(t, found.Value, v, tc.name)
			assert.True(t, tree.Contains(v), tc.name)
		}

		for _, v := range tc.expectAbsent {
			found := tree.Search(v)
			assert.Nil(t, found, tc.name)
		}
	}
}

func TestDelete(t *testing.T) {
	tree := BinarySearchTree{}
	tree.Insert(5)

	// find and delete that node
	five := tree.Search(5)
	tree.Delete(five)

	found := tree.Search(5)
	assert.Nil(t, found, "delete an element from the tree")
}

func TestTreeMinMax(t *testing.T) {
	tree := BinarySearchTree{}
	tree.Insert(5)
	assert.Equal(t, 5, tree.Min().Value)
	assert.Equal(t, 5, tree.Max().Value)

	tree.Insert(2)
	assert.Equal(t, 2, tree.Min().Value)
	assert.Equal(t, 5, tree.Max().Value)

	tree.Insert(7)
	assert.Equal(t, 7, tree.Max().Value)
}
