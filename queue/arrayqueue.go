package queue

import "fmt"

// ArrayQueue is a queue datastructure backed by a simple array (actually a Go slice)
type ArrayQueue[T interface{}] struct {
	head      int
	tail      int
	data      []T
	overflow  bool
	underflow bool
}

// NewArrayQueue returns a new queue, backed by a slice. Call with e.g. NewArrayQueue[string]()
func NewArrayQueue[T any](initialSize int) *ArrayQueue[T] {
	return &ArrayQueue[T]{
		data: make([]T, initialSize),
		// underflow is true by default; you can't Dequeue from an empty queue
		underflow: true,
	}
}

// Enqueue adds an item to the tail of the queue, then advances the tail
func (q *ArrayQueue[T]) Enqueue(val T) error {
	if q.len() == 0 {
		return fmt.Errorf("Zero-length queue: good luck with that.")
	}

	// we can't enqueue anything if our next write would be an overflow
	if q.overflow {
		return fmt.Errorf("Overflow")
	}

	// If we were about to hit an underflow on the next Dequeue, we're no longer in danger
	if q.underflow {
		q.underflow = false
	}

	nextIdx, err := q.nextTailIndex()
	if err != nil {
		return err
	}

	q.data[q.tail] = val
	q.tail = nextIdx
	return nil
}

// nextTailIndex calculates the next tail index, taking into account possible slice index overflows and slice data overflows
func (q *ArrayQueue[T]) nextTailIndex() (int, error) {
	// Are we at the end of the array? (i.e. do we need to move tail to the beginning?)
	if q.tail == q.len()-1 {
		// will we collide with the head if we move tail to the 0th element?
		if q.head == 0 {
			q.overflow = true
			return 0, nil
		}
		return 0, nil
	}

	// Are we right next to the head?
	if (q.tail + 1) == q.head {
		q.overflow = true
	}
	return q.tail + 1, nil
}

// nextHeadIndex calculates the next head index, taking into account possible slice index underflows and slice data underflows
func (q *ArrayQueue[T]) nextHeadIndex() (int, error) {
	// Are we at the end (do we need to move head to the beginning)?
	if q.head == q.len()-1 {
		if q.tail == 0 {
			q.underflow = true
		}
		return 0, nil
	}

	// Is head next to tail? i.e. would the next Dequeue underflow?
	if (q.head + 1) == q.tail {
		q.underflow = true
	}
	return q.head + 1, nil
}

func (q *ArrayQueue[T]) Dequeue() (T, error) {
	// defaultT is used as the zero-value for any type T
	var defaultT T

	// if underflow is true, the next dequeue (this one) will be an underflow
	if q.underflow {
		return defaultT, fmt.Errorf("Underflow")
	}

	// the next write is no longer an overflow
	if q.overflow {
		q.overflow = false
	}

	nextIdx, err := q.nextHeadIndex()
	if err != nil {
		return defaultT, err
	}

	val := q.data[q.head]
	q.head = nextIdx
	return val, nil
}

// len is a custom length-checker that checks capacity, not length, of a slice, so that the slice doesn't grow forever.
// this is necessary because dynamically sized arrays are not a thing in Go (size must be known at compile time)
func (q ArrayQueue[T]) len() int {
	return cap(q.data)
}
