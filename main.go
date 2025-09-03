package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) < 1 {
		// Equivalent to "-"
		arguments = append(arguments, "-")
	}
	for _, arg := range arguments {
		err := readInput(arg)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

	}
	os.Exit(0)
}
