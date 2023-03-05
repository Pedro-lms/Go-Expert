package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json: "numero"`
	Saldo  int `json: "saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 10000}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"numero":2,"saldo":5000}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}
	println(contaX.Numero)
	println(contaX.Saldo)
}
