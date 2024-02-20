package main

import "fmt"

func execAssertions() {
	fmt.Println("This is the assertions section.")
	assertions()
}

func assertions() {
	var minhaVariavel interface{} = "Victor Martinez"

	println(minhaVariavel)
	println(minhaVariavel.(string))

	res, ok := minhaVariavel.(int)
	println(res)
	println(ok)

}
