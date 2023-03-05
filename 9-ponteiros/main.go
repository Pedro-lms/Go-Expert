package main

func main() {
	//Memória -> Endereço valor

	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := a
	println(b)
	//fmt.Printf("o ponteiro %s representa o valor na memória de %d", ponteiro, a)
}
