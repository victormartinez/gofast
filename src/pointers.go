package main

import "fmt"

type Client struct {
	name string
}

func NewClient() *Client {
	return &Client{name: "Client"}
}

func (c *Client) walk() {
	c.name = "Victor Martinez"
	fmt.Printf("Client %v walked", c.name)
	fmt.Println()
}

func execPointers() {
	fmt.Println("This is the pointers section.")
	pointers()
}

func pointers() {
	/* Ponteiros são muito utilizados em GO pois em diversos momentos não queremos
	que os valores das variáveis estejam em escopo local. Muitas vezes queremos
	declarar variáveis que podem ser alteradas em diversas partes do código pois
	o valor manipulado está diretamente associado à localização na memória.
	*/
	a := 10
	fmt.Println("Valor de a =", a)
	fmt.Println("Endereço de memória de a =", &a)
	fmt.Println()

	var b *int
	fmt.Println("b guarda o ponteiro de memória de a")
	b = &a

	fmt.Println("Valor de b =", b)
	fmt.Println("Endereço de memória de b =", &b)
	fmt.Println("Valor do Endereço de memória de b =", *b)
	fmt.Println()

	fmt.Println("alterando b para apontar para o valor 30")
	*b = 30

	fmt.Println("Valor de b =", b)
	fmt.Println("Valor de *b =", *b)
	fmt.Println("Endereço de memória de b =", &b)

	fmt.Println("Valor de a =", a)

	teste()
	structsAndPointers()
}

func teste() {
	/*Quando passamos as variáveis como parâmetro, via de regras ocorre a passagem por valor
	(uma cópia do que existe na memória).
	*/
	var1 := 10
	var2 := 20
	fmt.Println("somaValores", somaValores(var1, var2))
	fmt.Println("var1", var1)

	fmt.Println("somaValoresPonteiro", somaValoresPonteiro(&var1, &var2))
	fmt.Println("var1", var1)
}

func somaValores(a, b int) int {
	a = 30
	return a + b
}

func somaValoresPonteiro(a, b *int) int {
	*a = 30
	return *a + *b
}

func structsAndPointers() {
	client := Client{
		name: "Victor",
	}
	client.walk()

	fmt.Printf("Struct %v", client.name)
	fmt.Println()

	newClient := NewClient()
	fmt.Printf("Struct %v", newClient.name)
	fmt.Println()
	newClient.name = "Victor"
	fmt.Printf("Struct %v", newClient.name)
	fmt.Println()
}
