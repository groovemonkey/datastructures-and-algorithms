package queue

import "testing"

func TestArrayQueueInitSize(t *testing.T) {
	q := NewArrayQueue[string](2)

	err := q.Enqueue("hello")
	if err != nil {
		t.Logf("error enqueuing: %s", err)
	}

	err = q.Enqueue("world")
	if err != nil {
		t.Logf("error enqueuing: %s", err)
	}

	hello, err := q.Dequeue()
	if err != nil {
		t.Logf("Dequeue produced error: %s", err)
		t.Fail()
	}
	if hello != "hello" {
		t.Logf("first Dequeue() did not produce first value: hello was %s", hello)
		t.Fail()
	}

	world, err := q.Dequeue()
	if err != nil {
		t.Logf("hello is not set: %s", hello)
		t.Fail()
	}
	if world != "world" {
		t.Logf("world is not set: %s", world)
		t.Fail()
	}
}

func TestArrayQueueZeroLen(t *testing.T) {
	q := NewArrayQueue[int](0)

	err := q.Enqueue(5)
	if err == nil {
		t.Log("expected error enqueuing to zero-length queue")
		t.Fail()
	}

	v, err := q.Dequeue()
	if err == nil {
		t.Log("expected error dequeuing from zero-length queue")
		t.Fail()
	}
	if v != 0 {
		t.Log("expected default value from failed Dequeue")
		t.Fail()
	}
}

func TestOverflow(t *testing.T) {
	q := NewArrayQueue[int](1)
	q.Enqueue(0)
	err := q.Enqueue(100)
	if err == nil {
		t.Log("expected error enqueuing to full queue")
		t.Fail()
	}
}

func TestUnderflow(t *testing.T) {
	q := NewArrayQueue[int](1)
	q.Enqueue(0)
	_, _ = q.Dequeue()
	_, err := q.Dequeue()
	if err == nil {
		t.Log("expected error dequeueing from empty queue")
		t.Fail()
	}
}
