package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readInput(inputString string) error {
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
		fmt.Print(string(line))
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
