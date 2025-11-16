package main

import "fmt"

func main() {
	fmt.Println("We are learning array")

	//Array declaration
	var name [3]string
	name[0] = "Sonal"
	name[1] = "chhotu"
	name[2] = "Vedik"

	fmt.Println("name of the person is: ", name)
	fmt.Println("name of the second person is: ", name[1])

	//Array declaration and intialization
	var number = [4]int{3, 4, 6, 7}
	fmt.Println("Print all the numbers: ", number)
	fmt.Println("Print number at 1st and 2nd index:", number[1], number[2])

	//Length of an array
	fmt.Println("Length of numbers array is: ", len(number))

	//if no values intiated, it sets as default zero for int type, Empty ("") for String, false for bool, pointers or complex type it is nil
	// var price [3]int
	// fmt.Println("Default price is: ", price)

	// price[1] = 45
	// fmt.Println("price of 1st index is initated: ", price[1])

	// fmt.Println("Updated price list is: ", price)

	// var price [3]string
	// fmt.Println("Price is", price)

	var price [3]string
	fmt.Println("Price is", price)

	fmt.Printf("Value of Price is: %q ", price)
}
