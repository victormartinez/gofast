package main

import "fmt"

func arrays() {
	fmt.Println("This is the array section.")
	printArray()
	printSlicedArray()
}

func printArray() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println("Array length=", len(meuArray))

	for i := 0; i < len(meuArray); i++ {
		fmt.Println(meuArray[i])
	}

	for i, v := range meuArray {
		fmt.Printf("Index=%d Value=%d\n", i, v)
	}

	i := 0
	for i < 10 {
		if i > 1 && i < 3 {
			println(i)
		}
		i++
	}

	switch i {
	case 1:
		println("A")
	default:
		println("SWITCH")
	}

	// LOOP INFINITO
	// for {
	// 	print("AQ!UI")
	// }
}

func printSlicedArray() {
	slicedArray := []int{1000, 2000, 3000, 4000}
	fmt.Printf("len=%d capacity=%d %v\n", len(slicedArray), cap(slicedArray), slicedArray)
	fmt.Printf("len=%d capacity=%d %v\n", len(slicedArray[:2]), cap(slicedArray[:2]), slicedArray[:2])
	fmt.Printf("len=%d capacity=%d %v\n", len(slicedArray[2:]), cap(slicedArray[2:]), slicedArray[2:])
	fmt.Printf("len=%d capacity=%d %v\n", len(slicedArray[:]), cap(slicedArray[:]), slicedArray[:])

	// Veja o CAP = 8 apesar de só existirem 5 elementos!
	// O Go, por debaixo dos panos, dobra o tamanho do slice sempre que precisar redimensionar.
	// Se você for trabalhar com muitos dados, tente sempre inicializar com um slice próximo
	// da capacidade que você vai trabalhar.
	// Se você trabalhar com um valor muito diferente, vai ser um processo custoso para o Go ficar redimensionando.
	slicedArray = append(slicedArray, 5000)
	fmt.Printf("(ATTENTION!) len=%d capacity=%d %v\n", len(slicedArray), cap(slicedArray), slicedArray)
}
