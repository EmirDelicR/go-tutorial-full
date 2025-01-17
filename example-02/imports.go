package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide one argument: example (go run . 9000)")
		os.Exit(1)
	}

	fmt.Println(os.Args)
	fmt.Println("It's over", os.Args[1])
}
