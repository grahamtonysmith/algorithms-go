package fundimentals

import (
	"iter"
)

type Bag[T any] struct {
	head *Node[T]
	n    int
}

func (b *Bag[T]) IsEmpty() bool {
	return b.n == 0
}

func (b *Bag[T]) Size() int {
	return b.n
}

func (b *Bag[T]) Add(item T) {
	node := &Node[T]{
		item: item,
		next: b.head,
	}

	b.head = node
	b.n += 1
}

func (b *Bag[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		node := b.head

		for node != nil {
			ok := yield(node.item)
			if !ok {
				return
			}

			node = node.next
		}
	}
}

func (b *Bag[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		node := b.head

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

func NewBag[T any]() *Bag[T] {
	return &Bag[T]{
		head: nil,
		n:    0,
	}
}
