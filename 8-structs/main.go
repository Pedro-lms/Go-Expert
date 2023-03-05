package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Client struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco //Linked Struct
	Endereco
}

func (c Client) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
	fmt.Println()
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

func main() {

	pedro := Client{
		Nome:  "Pedro Sadit",
		Idade: 27,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", pedro.Nome, pedro.Idade, pedro.Ativo)
	fmt.Println()
	Desativacao(pedro)
	pedro.Address.Cidade = "São Paulo"
	fmt.Println(pedro.Nome, pedro.Ativo, pedro.Address.Cidade)
}
