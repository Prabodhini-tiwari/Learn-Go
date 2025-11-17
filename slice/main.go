// In Go lang, slice is a flexible and dynamic data structure that provides a more poowerful alternative to arrays.
// Slices have dynamic size, and their length can be changed during the program's execution.

package main

import "fmt"

func main() {
	fmt.Println("We are learning Slice in Golang")

	// numbers := []int{2, 4, 5}

	// numbers = append(numbers, 6, 7, 9, 10)
	// fmt.Printf("type of numbers: %T\n", numbers)

	// fmt.Println("The length of numbers: ", len(numbers))
	// fmt.Println("capacity of numbers: ", cap(numbers))

	// fmt.Println("elements in numbers: ", numbers)

	// make function

	stock := make([]int, 0, 5)

	stock = append(stock, 4, 5)
	stock = append(stock, 6, 7, 8)
	stock = append(stock, 12)

	fmt.Println("Stocks are: ", stock)
	fmt.Println("Stock's length:", len(stock))
	fmt.Println("Stock's capacity:", cap(stock))

}
