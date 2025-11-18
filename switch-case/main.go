package main

import "fmt"

func main() {
	fmt.Println("We are learning switch-case in golang")

	// Example 1: Basic switch statement
	day := 8

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")

	default:
		fmt.Println("Invalid day")
	}

	// Example 2: Switch statement with multiple values
	months := "June"

	switch months {
	case "December", "January", "February":
		fmt.Println("Winter season")

	case "March", "April", "May", "June":
		fmt.Println("Spring")

	default:
		fmt.Println("Other Season")

	}

	// Example 3: Switch statement with conditions
	temperature := 15

	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 10:
		fmt.Println("cold")
	case temperature >= 10 && temperature < 20:
		fmt.Println("cool")
	case temperature >= 20 && temperature < 30:
		fmt.Println("warm")
	default:
		fmt.Println("Hot")
	}
}
