package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/grahamtonysmith/algorithms/pkg/sorting"
)

func main() {
	fmt.Printf("String to sort: ")

	br := bufio.NewReader(os.Stdin)

	str, err := br.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	data := sorting.StringSlice(strings.Split(strings.TrimSpace(str), ""))

	sorting.MergeSort(data)

	fmt.Println("Sorted string:", strings.Join(data, ""))
}
