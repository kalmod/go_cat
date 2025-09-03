package main

import (
	"flag"
	"fmt"
	"os"
)

type opts struct {
	numbFlag *bool
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) < 1 {
		// No args are Equivalent to "-"
		arguments = append(arguments, "-")
	}

	opt := opts{}

	opt.numbFlag = flag.Bool("number", false, "number all output lines")
	flag.BoolVar(opt.numbFlag, "n", false, "number all output lines (short arg)")

	flag.Parse()

	for _, arg := range arguments {
		err := readInput(arg)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

	}
	os.Exit(0)
}
