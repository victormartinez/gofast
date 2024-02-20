package main

import (
	"errors"
	"fmt"
)

func functions() {
	fmt.Println("This is the functions section.")
	fmt.Println(sum(1, 2))
	fmt.Println(sum2(1, 2))

	valor, err := sum3(1, 2)
	fmt.Println(valor, err)

	valor2, err2 := sum4(1, 2)
	fmt.Println(valor2, err2)

	fmt.Println(sumMany(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(closure(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func sum4(a, b int) (int, error) {
	// Go não tem exceptions
	return a + b, errors.New("Error")
}

func sum3(a, b int) (int, error) {
	// Go não tem exceptions
	return a + b, nil
}

func sum2(a, b int) (int, bool) {
	return a + b, true
}

func sum(a int, b int) int {
	return a + b
}

func sumMany(numeros ...int) int { // função variática
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

// funções anonimas: closures
func closure(values ...int) int {
	multiplier := 2

	totalSoma := func() int {
		total := 0
		for _, num := range values {
			total += num * multiplier
		}
		return total
	}()

	return totalSoma
}
