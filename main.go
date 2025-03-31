package main

import (
	"bufio"
	"fmt"
	"github.com/hhirsch/sps-buddy/internal/models/style"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type Block int64

const (
	Variable Block = iota
	Constant
	None
)

var arguments []string = os.Args
var isErrorDetected bool = false
var variableBlockStartStrings = []string{"VAR_INPUT", "VAR", "VAR_OUTPUT", "VAR_TEMP", "VAR_GLOBAL"}
var currentBlock Block = None

func checkVariableStyle(symbol string) {
	if style.IsMixedCamelCase(symbol) {
		fmt.Printf("Success: Variable %s is camel case.\n", symbol)
	} else {
		fmt.Fprintf(os.Stderr, "Error: Variable %s is not mixed camel case.\n", symbol)
		isErrorDetected = true
	}
}

func checkConstantStyle(symbol string) {
	if style.IsMixedCapitalSnakeCase(symbol) {
		fmt.Printf("Success: Constant %s is mixed capital snake case.\n", symbol)
	} else {
		fmt.Fprintf(os.Stderr, "Error: Constant %s is not mixed capital snake case.\n", symbol)
		isErrorDetected = true
	}
}

func handleLine(input string) {
	if strings.TrimSpace(input) == "VAR_GLOBAL CONSTANT" {
		currentBlock = Constant
		return
	}
	inputWithoutWhiteSpaces := strings.ReplaceAll(input, " ", "")
	if currentBlock == Variable {
		if strings.Contains(inputWithoutWhiteSpaces, "{") {
			parts := strings.Split(inputWithoutWhiteSpaces, "{")
			checkVariableStyle(parts[0])
			return
		}
		if strings.Contains(inputWithoutWhiteSpaces, ":") {
			parts := strings.Split(inputWithoutWhiteSpaces, ":")
			checkVariableStyle(parts[0])
			return
		}
	}

	if currentBlock == Constant {
		if strings.Contains(inputWithoutWhiteSpaces, ":") {
			parts := strings.Split(inputWithoutWhiteSpaces, ":")
			checkConstantStyle(parts[0])
			return
		}
	}

	if slices.Contains(variableBlockStartStrings, inputWithoutWhiteSpaces) {
		if currentBlock == None {
			currentBlock = Variable
			return
		} else {
			fmt.Fprintf(os.Stderr, "Error: Found beginning of new variable block before old one has ended: %v\n", input)
			isErrorDetected = true
		}
	}

	if inputWithoutWhiteSpaces == "END_VAR" {
		currentBlock = None
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
