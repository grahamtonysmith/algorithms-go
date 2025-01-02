package fundimentals

import (
	"iter"
)

type Bag[T any] struct {
	items []T
}

func (b *Bag[T]) Add(item T) {
	b.items = append(b.items, item)
}

func (b *Bag[T]) IsEmpty() bool {
	return len(b.items) == 0
}

func (b *Bag[T]) Size() int {
	return len(b.items)
}

func (b *Bag[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range b.items {
			ok := yield(item)
			if !ok {
				return
			}
		}
	}
}

func (b *Bag[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, item := range b.items {
			ok := yield(i, item)
			if !ok {
				return
			}
		}
	}
}

func NewBag[T any]() *Bag[T] {
	return &Bag[T]{
		items: make([]T, 0),
	}
}
