package main

//funções variáveis
import "fmt"

func main() {
	fmt.Println(sum(1, 3, 45, 6, 34, 654, 7465, 543, 555))
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
