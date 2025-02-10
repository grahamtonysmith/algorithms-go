package main

import (
	"log"

	"github.com/grahamtonysmith/algorithms/pkg/fundimentals"
)

func main() {
	stack := fundimentals.NewStrackResizingArray[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	if stack.IsEmpty() {
		log.Fatalln("stack is empty")
	}

	if stack.Size() != 5 {
		log.Fatalln("stack size is not 5")
	}

	if stack.Pop() != 5 {
		log.Fatalln("last in was not 5")
	}

	for i, number := range stack.All() {
		log.Println("i", i, "number", number)
	}

	log.Println(stack)
}
