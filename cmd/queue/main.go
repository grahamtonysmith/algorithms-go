package main

import (
	"log"

	"github.com/grahamtonysmith/algorithms/pkg/fundimentals"
)

func main() {
	queue := fundimentals.NewQueue[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	if queue.IsEmpty() {
		log.Fatalln("queue is empty")
	}

	if queue.Size() != 5 {
		log.Fatalln("queue size is not 5")
	}

	if queue.Dequeue() != 1 {
		log.Fatalln("first in was not 1")
	}

	for i, number := range queue.All() {
		log.Println("i", i, "number", number)
	}
}
