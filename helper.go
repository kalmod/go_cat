package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func readInput(opt *opts, inputString string) error {
	fileCheck := isFile(inputString)
	var scanner *bufio.Reader
	if fileCheck {
		file, err := openFile(inputString)
		if err != nil {
			return fmt.Errorf("ERROR:Reading file::%s", err)
		}
		defer file.Close()
		scanner = bufio.NewReader(file)
	} else if !fileCheck && inputString != "-" {
		return fmt.Errorf("%s: No such file or directory", inputString)
	} else {
		scanner = bufio.NewReader(os.Stdin)
	}

	for {
		line, err := scanner.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("ERROR::Reading File::%s", err)
		}
		formatPrinterr := formatPrint(opt, line)
		if formatPrinterr != nil {
			return fmt.Errorf("ERROR::Printing line::%s", formatPrinterr)
		}
	}
}

func isFile(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}

func openFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		return nil, fmt.Errorf("ERROR::%s", err.Error())
	}
	return file, err
}

func formatPrint(opt *opts, line []byte) error {
	var sb strings.Builder
	if *opt.numbFlag || *opt.numbNBFlag {
		var err error
		if *opt.numbNBFlag && len(line) == 1 && line[0] == byte('\n') {
			_, err = sb.WriteString(fmt.Sprintf("%6s ", " "))
		} else {
			opt.lineNumber++
			_, err = sb.WriteString(fmt.Sprintf("%6d ", opt.lineNumber))

		}
		if err != nil {
			return err
		}

	}
	_, err := sb.Write(line)
	if err != nil {
		return err
	}
	fmt.Print(sb.String())
	return nil
}
