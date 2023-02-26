// Closures
package main

//funções variáveis
import "fmt"

func main() {

	total := func() int {
		return sum(1, 3, 45, 6, 34, 654, 7465, 543, 555) * 2
	}()
	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
