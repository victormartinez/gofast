package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     string
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func structs() {
	fmt.Println("This is the structs section.")
	execstruct()
}

func execstruct() {
	victor := Cliente{
		Nome:  "Victor",
		Idade: 32,
		Ativo: true,
		Endereco: Endereco{
			Logradouro: "Rua",
			Numero:     "306",
			Cidade:     "Salvador",
			Estado:     "Bahia",
		},
	}
	victor.Endereco.Logradouro = "RUA JOÃO BIÃO DE CERQUEIRA"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t Endereco: %s\n", victor.Nome, victor.Idade, victor.Ativo, victor.Endereco)

	victor.Desativar()
	fmt.Printf("Nome: %s, Ativo: %t\n", victor.Nome, victor.Ativo)
}

func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Println("Desativando...")
}
