package fundimentals

import (
	"iter"
)

// resizing array implementation of a pushdown stack
// can not use nil assignment to avoid loitering
type StrackResizingArray[T any] struct {
	items []T
}

func (r *StrackResizingArray[T]) IsEmpty() bool {
	return len(r.items) == 0
}

func (r *StrackResizingArray[T]) Size() int {
	return len(r.items)
}

func (r *StrackResizingArray[T]) Push(item T) {
	if len(r.items) == cap(r.items) {
		temp := make([]T, len(r.items), cap(r.items)*2)

		copy(temp, r.items)

		r.items = temp
	}

	r.items = append(r.items, item)
}

func (r *StrackResizingArray[T]) Pop() T {
	item := r.items[len(r.items)-1]
	r.items = r.items[0 : len(r.items)-1]

	if len(r.items) == cap(r.items)/4 {
		temp := make([]T, len(r.items), cap(r.items)/2)

		copy(temp, r.items)

		r.items = temp
	}

	return item
}

func (r *StrackResizingArray[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := len(r.items); i > 0; i-- {
			ok := yield(r.items[i-1])
			if !ok {
				return
			}
		}
	}
}

func (r *StrackResizingArray[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i := len(r.items); i > 0; i-- {
			ok := yield(len(r.items)-i, r.items[i-1])
			if !ok {
				return
			}
		}
	}
}

func NewStrackResizingArray[T any]() *StrackResizingArray[T] {
	return &StrackResizingArray[T]{
		items: make([]T, 0, 1), // (data, len, capacity)
	}
}
