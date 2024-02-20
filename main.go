package main

import (
	"fmt"
	"gofast/custommath/custommath"
)

func main() {
	// No need to import because intro also declares itself belonging to main package
	fmt.Println("-----------------------------")
	intro()
	fmt.Println("-----------------------------")
	arrays()
	fmt.Println("-----------------------------")
	hashtables()
	fmt.Println("-----------------------------")
	functions()
	fmt.Println("-----------------------------")
	structs()
	fmt.Println("-----------------------------")
	execInterface()
	fmt.Println("-----------------------------")
	execPointers()
	fmt.Println("-----------------------------")
	execEmptyInterface()
	fmt.Println("-----------------------------")
	execAssertions()
	fmt.Println("-----------------------------")
	execGenerics()
	fmt.Println("-----------------------------")

	result := custommath.SomaCustomMath(10, 10)
	println(result)
}
