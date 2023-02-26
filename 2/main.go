package main

import "fmt"

const a = "Ola"

type ID int

var (
	b bool = true
	c int
	d string
	e float64 = 7.7
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo de E é %T", e)
	println()
	a := "X"
	println(f)
	println(e)
	println(d)
	println(c)
	println(b)
	println(a)
}

func x() {

}
