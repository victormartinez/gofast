package main

import "fmt"

type NumberType int

type Number interface {
	~int | ~float64 // ~ faz considerar que pode usar o NumberType como Number
}

func execGenerics() {
	fmt.Println("This is the generics section.")
	generics()
}

func generics() {
	m := map[string]int{"Victor": 1000, "Manuela": 2000}
	println(somaGenerics(m))

	m2 := map[string]NumberType{"Victor": 2000, "Manuela": 3000}
	println(somaGenericsWithInterface(m2))

	println(Compara(10, 10.0))
	println(Compara(10, 11.0))
}

func somaGenerics[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func somaGenericsWithInterface[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// Permite usar o T desde que ele seja comparável
func Compara[T comparable](a T, b T) bool {
	// a > b não funciona pois o comparable so sabe se são iguais
	if a == b {
		return true
	}
	return false
}
