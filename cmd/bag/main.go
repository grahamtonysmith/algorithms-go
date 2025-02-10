package main

import (
	"log"

	"github.com/grahamtonysmith/algorithms/pkg/fundimentals"
)

func main() {
	bag := fundimentals.NewBag[int]()

	bag.Add(1)
	bag.Add(2)
	bag.Add(3)

	if bag.IsEmpty() {
		log.Fatalln("bag is empty")
	}

	if bag.Size() != 3 {
		log.Fatalln("bag size is not 3")
	}

	for i, n := range bag.All() {
		log.Printf("i=%d n=%d\n", i, n)
	}
}
