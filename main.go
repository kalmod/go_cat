package main

import (
	"flag"
	"fmt"
	"os"
)

type opts struct {
	numbFlag   *bool
	numbNBFlag *bool
	lineNumber int
}

func main() {
	var opt opts

	opt.numbFlag = flag.Bool("number", false, "number all output lines")
	opt.numbNBFlag = flag.Bool("number-nonblank", false, "number non-blank lines")
	flag.BoolVar(opt.numbFlag, "n", false, "number all output lines (short arg)")
	flag.BoolVar(opt.numbNBFlag, "b", false, "number non-blank lines")

	flag.Parse()

	arguments := flag.Args()
	if len(arguments) < 1 {
		// No args are Equivalent to "-"
		arguments = append(arguments, "-")
	}

	for _, arg := range arguments {
		err := readInput(&opt, arg)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

	}
	os.Exit(0)
}
