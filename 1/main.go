package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 30

	for i, v := range meuArray {
		fmt.Printf("O valor do índice é %d e o valor é %d\n ", i, v)
	}
}
