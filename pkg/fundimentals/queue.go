package fundimentals

import "iter"

// FIFO
type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	n    int
}

func (q *Queue[T]) IsEmpty() bool {
	return q.n == 0
}

func (q *Queue[T]) Size() int {
	return q.n
}

func (q *Queue[T]) Enqueue(item T) {
	node := &Node[T]{
		item: item,
		next: nil,
	}

	if q.IsEmpty() {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}

	q.n += 1
}

func (q *Queue[T]) Dequeue() T {
	node := q.head

	q.head = q.head.next
	q.n -= 1

	return node.item
}

func (q *Queue[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		node := q.head

		for node != nil {
			ok := yield(node.item)
			if !ok {
				return
			}

			node = node.next
		}
	}
}

func (q *Queue[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		node := q.head

		for node != nil {
			ok := yield(i, node.item)
			if !ok {
				return
			}

			node = node.next
			i += 1
		}
	}
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		head: nil,
		tail: nil,
		n:    0,
	}
}
