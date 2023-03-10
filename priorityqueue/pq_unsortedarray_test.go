package priorityqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertAndFindMin(t *testing.T) {
	tcs := []struct {
		name      string
		insert    []int
		expectMin int
	}{
		{
			name:      "one insert",
			insert:    []int{5},
			expectMin: 5,
		},
		{
			name:      "a few inserts",
			insert:    []int{7, 2301231, 1},
			expectMin: 1,
		},
	}

	for _, tc := range tcs {
		pq := PriorityQueueUnsortedArray{}
		for _, v := range tc.insert {
			pq.Insert(v)
		}
		assert.Equal(t, tc.expectMin, pq.FindMin(), tc.name)
	}
}

func TestDeleteMin(t *testing.T) {
	pq := PriorityQueueUnsortedArray{}
	pq.Insert(5)
	pq.Insert(10)
	pq.DeleteMin()
	assert.Equal(t, 10, pq.FindMin(), "delete minimum element")
}
