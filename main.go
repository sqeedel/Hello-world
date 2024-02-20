package main

import "fmt"

func main() {
	fmt.Println(" Hello " + " Akezhonus ")
	fmt.Println(1 + 100)
	fmt.Println("true")
	//declartion
	var x string

	//initialization
	x = "Elnur"

	fmt.Println(x)

	//declaration and initialization
	y := "Kostya"
	fmt.Println(y)

	var z string = "Ivan"
	fmt.Println(z)

	//boolean (true/false)
	var a bool = true
	//integer
	var b int = 1
	//floating point number
	var c float32 = 3.14
	var d string = "HI!"

	fmt.Println("boolean: ", a)
	fmt.Println("integer: ", b)
	fmt.Println("float: ", c)
	fmt.Println("string: ", d)

	//constanta
	const pi = 3.14

	fmt.Println(pi)

	// +
	// -
	// *
	// /
	// %
	r := 10
	u := 31
	result := u % r
	println(result)

}
