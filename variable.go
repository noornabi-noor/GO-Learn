package main

import "fmt"

func main () {
	// var name string = "Noornabi"
	name := "Noornabi"
	num := 10

	fmt.Println("Hello, ", name)
	fmt.Println("The number is: ", num)

	// Group variable declaration
	var (
		firstName string = "Noor"
		lastName string = "Nabi"
		age int = 23
	)

	fmt.Println("First Name: ", firstName)
	fmt.Println("Last Name: ", lastName)
	fmt.Println("Age: ", age)

	// multiple variable declaration
	var a, b, c int = 1, 2, 3
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)

	// short variable declaration with multiple variables
	x, y := 4, 5
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)

	// Constants
	const pi = 3.14
	fmt.Println("Value of pi: ", pi)
}