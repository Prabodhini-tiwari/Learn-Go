package main

import "fmt"

func main() {
	fmt.Println("Let's learn if-else conditions in golang")

	x := 1

	if x > 10 {
		fmt.Println("X is greater than 10")
	} else {
		fmt.Println("X is smaller than 10")
	}

	age := 18

	if age >= 18 {
		fmt.Println("Adult")
	} else if age >= 13 {
		fmt.Println("teenagers")
	} else {
		fmt.Println("Underage")
	}

}
