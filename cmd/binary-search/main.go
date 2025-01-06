package main

import (
	"log"

	"github.com/grahamtonysmith/algorithms/pkg/fundimentals"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	key := 6
	val := fundimentals.BinarySrarch(arr, key)

	if val != -1 {
		log.Printf("found: %d\n", val)
	} else {
		log.Println("key not found")
	}
}
