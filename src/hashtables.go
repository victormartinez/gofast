package main

import "fmt"

func hashtables() {
	fmt.Println("This is the hashtable section.")
	printHash()
}

func printHash() {
	salarios := map[string]int{
		"Victor": 25000, "João": 34000,
	}

	fmt.Println(salarios["Victor"])
	delete(salarios, "Victor")
	fmt.Println(salarios["Victor"])
	fmt.Println(salarios["João"])

	salarios["Carlos"] = 30401

	// novosSalarios := make(map[string]int)
	// novosSalarios := map[string]int{}

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}
}
