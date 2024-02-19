package main

import "fmt"

// Interfaces em GO só permitem passar métodos!
type Pessoa interface {
	Desativar()
	Ativar()
}

type Cidadao struct {
	Nome  string
	Idade int
	Ativo bool
}

func (c Cidadao) Desativar() {
	c.Ativo = false
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", c.Nome, c.Idade, c.Ativo)
}

func (c Cidadao) Ativar() {
	c.Ativo = true
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func execInterface() {
	fmt.Println("This is the interface section.")
	inter()
}

func inter() {
	cidadao := Cidadao{
		Nome:  "Victor",
		Idade: 32,
		Ativo: true,
	}
	// cidadao.Desativar()
	// fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", cidadao.Nome, cidadao.Idade, cidadao.Ativo)

	// cidadao.Ativar()
	// fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", cidadao.Nome, cidadao.Idade, cidadao.Ativo)
	Desativacao(cidadao)

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", cidadao.Nome, cidadao.Idade, cidadao.Ativo)
}
