package main

import "fmt"

// use isso com moderação
type x interface{}

func execEmptyInterface() {
	fmt.Println("This is the empty interface section.")
	execEmpty()
}

func execEmpty() {
	var x interface{} = 10
	var y interface{} = "Hello, World!"
	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é '%v'\n", t, t)
}
