package main

import "fmt"

func main() {
	fmt.Println("Hi I am learning Go ")

	age := 28
	name := "Striver"
	height := 5.736

	// Using Println() to print output

	// fmt.Println("Age is: ", age)
	// fmt.Println("name is:", name)
	// fmt.Println("height is:", height)

	// Using Printf() to format the code  , %d for integer, %s for string, %f for float type
	fmt.Printf("Age is: %d\n", age)
	fmt.Printf("Name is: %s\n", name)
	fmt.Printf("height is: %.1f\n", height)

	fmt.Printf("Age is %d, name is %s, height is %.2f\n", age, name, height)

}
