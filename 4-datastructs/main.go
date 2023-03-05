package main

import "fmt"

//Maps = hashtables
func main() {
	salarios := map[string]int{"Pedro": 10000, "Claudia": 2500}
	fmt.Println(salarios)
	fmt.Println(salarios["Claudia"])
}
