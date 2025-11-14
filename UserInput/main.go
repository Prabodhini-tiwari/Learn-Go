package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("let's Understand How to take input from a user")
	fmt.Println("________________________________________________")

	// var name string

	// fmt.Scan(&name) // Scan() only read till the whiteSpace and ampersand(&) to pass the variable's address
	// fmt.Println("Hello Mr.", name)

	reader := bufio.NewReader(os.Stdin)

	name, _ := reader.ReadString('\n') // Read until newline

	fmt.Println("Hello", name)
}
