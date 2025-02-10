package fundimentals

// a linked list is a recursive data structure that is either empty (nil) and
// a node having a generic item and a reference to a node having a generic item
// and a reference to a linked list

// a linked list represents a sequence of items

type Node[T any] struct {
	item T
	next *Node[T]
}
