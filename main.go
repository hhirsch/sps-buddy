package main

import (
	"bufio"
	"fmt"
	"github.com/hhirsch/sps-buddy/internal/models/style"
	"os"
	"path/filepath"
	"strings"
)

var arguments []string = os.Args
var parserIsInsideVariableBlock bool = false
var isErrorDetected bool = false

func handleLine(input string) {
	inputWithoutWhiteSpaces := strings.ReplaceAll(input, " ", "")
	if parserIsInsideVariableBlock {
		if strings.Contains(inputWithoutWhiteSpaces, "{") {
			parts := strings.Split(inputWithoutWhiteSpaces, "{")
			if style.IsMixedCamelCase(parts[0]) {
				fmt.Printf("Success: Variable %s is camel case.\n", parts[0])
			} else {
				fmt.Fprintf(os.Stderr, "Error: Variable %s is not camel case.\n", parts[0])
				isErrorDetected = true
			}
			return
		}
		if strings.Contains(inputWithoutWhiteSpaces, ":") {
			parts := strings.Split(inputWithoutWhiteSpaces, ":")
			if style.IsMixedCamelCase(parts[0]) {
				fmt.Printf("Success: Variable %s is camel case.\n", parts[0])
			} else {
				fmt.Fprintf(os.Stderr, "Error: Variable %s is not camel case.\n", parts[0])
				isErrorDetected = true
			}
			return
		}
	}

	if inputWithoutWhiteSpaces == "VAR_INPUT" ||
		inputWithoutWhiteSpaces == "VAR" ||
		inputWithoutWhiteSpaces == "VAR_OUTPUT" ||
		inputWithoutWhiteSpaces == "VAR_TEMP" {
		parserIsInsideVariableBlock = true
		return
	}

	if inputWithoutWhiteSpaces == "END_VAR" {
		parserIsInsideVariableBlock = false
	}
}

func processFile(fileName string) int {
	fmt.Printf("\nReading file: %s\n", fileName)
	file, err := os.Open(fileName)
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("error closing file %v", err)
		}
	}()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %s.\n", err.Error())
		return 1
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		handleLine(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning file: %s.\n", err.Error())
		isErrorDetected = true
	}

	if isErrorDetected {
		return 1
	}
	return 0
}

func main() {
	if len(arguments) < 2 {
		fmt.Printf("Command needs an argument.\n")
		os.Exit(1)
	}

	if strings.Contains(arguments[1], ".scl") {
		var fileName = arguments[1]
		os.Exit(processFile(fileName))
	}

	if strings.Contains(arguments[1], "--batch") {
		var errorCounter int
		err := filepath.WalkDir("./", func(path string, dirEntry os.DirEntry, err error) error {
			if !dirEntry.IsDir() && filepath.Ext(dirEntry.Name()) == ".scl" {
				errorCounter += processFile(path)
			}

			return nil
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while trying to scan the directory: %v\n", err)
			os.Exit(1)
		}
		if errorCounter > 0 {
			fmt.Fprintf(os.Stderr, "Coding standards not met.\n")
			os.Exit(1)
		} else {
			fmt.Printf("No coding standards violation detected.\n")
			os.Exit(0)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error in batch processing:  %v\n", err)
			os.Exit(1)
		}

	}

	fmt.Printf("Unexpected parameters.\n")
	os.Exit(1)

}
