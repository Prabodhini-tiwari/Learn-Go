package main

import "fmt"

// Correct way to write a function
// func main() {
// 	fmt.Println("This is the correct way to write a function")
// }

// Wrong way
// func main() {
// 	fmt.Println("This is the incorrect way to write a function")
// }

// function with input parameter
// func add(a int, b int) int {

// 	return a + b
// }

// another way to write input parameter
// func add(a, b int) int {

// 	return a + b
// }

// function with input parameter and output parameter with type
func add(a, b int) (result int) {

	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func main() {
	fmt.Println("We are Learning Function in Golang")
	ans := add(3, 6)

	fmt.Println("Add two number: ", ans)

	fmt.Println("__________________")

	answer := multiply(4, 6)
	fmt.Println("Multiplication of two numbers are: ", answer)
}
