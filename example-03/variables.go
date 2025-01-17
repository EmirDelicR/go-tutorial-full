package main

import (
	"fmt"
)

/**
There’s more to learn about declaration and assignments. For now, remember that you’ll use:

var NAME TYPE when declaring a variable to its zero value,

NAME := VALUE when declaring and assigning a value,

and NAME = VALUE when assigning to a previously declared variable.
**/

func main() {
	power := getPower()
	fmt.Printf("It's over %d\n", power)
	power = 9001
	fmt.Printf("It's over %d\n", power)
	printDetails()
	// need to declare this values
	name, pow := powerByName("Vegeta")
	// If you need just one you can do
	// _, pow := powerByName("Vegeta")
	fmt.Printf("%s's power is over %d\n", name, pow)
}

func getPower() int {
	// Declare variable
	// option 1: var power int = 9000
	// option 2: power: = 9000
	power := 1000
	// update value
	power = 9000
	return power
}

func printDetails() {
	// Declare multiple variables in same line
	name, power := "Goku", 1000
	fmt.Printf("%s's power is over %d\n", name, power)
}

// Function can return multiple values
func powerByName(name string) (string, int) {
	if name == "Goku" || name == "goku" {
		return "Goku", 1000
	}

	return name, 0
}
