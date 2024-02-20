package main

import "fmt"

func intro() {
	fmt.Println("This is the intro section.")
	variables()
}

func variables() {
	const a = "Hello, World!"

	type ID int // Custom Type
	var (
		b bool    // false
		c int     // 0
		d string  //
		e float64 // +0.000000e+000
		f ID      = 1
	)

	fmt.Println("Constant =", a)
	fmt.Println("Default bool =", b)
	fmt.Println("Default int =", c)
	fmt.Println("Default string =", d)
	fmt.Println("Default float64 =", e)
	fmt.Println("Default ID(custom) =", f)

	g := "X" // no need to explicitly type because of the assignment
	println("Value of g =", g)
}
