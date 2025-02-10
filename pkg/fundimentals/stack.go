package fundimentals

import "iter"

// LIFO
type Stack[T any] struct {
	head *Node[T]
	n    int
}

func (s *Stack[T]) IsEmpty() bool {
	return s.n == 0
}

func (s *Stack[T]) Size() int {
	return s.n
}

func (s *Stack[T]) Push(item T) {
	node := &Node[T]{
		item: item,
		next: s.head,
	}

	s.head = node
	s.n += 1
}

func (s *Stack[T]) Pop() T {
	item := s.head.item

	s.head = s.head.next
	s.n -= 1

	return item
}

func (s *Stack[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		node := s.head

		for node != nil {
			ok := yield(node.item)
			if !ok {
				return
			}

			node = node.next
		}
	}
}

func (s *Stack[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		node := s.head

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

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		head: nil,
		n:    0,
	}
}
