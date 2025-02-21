package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var arguments []string = os.Args
var parserIsInsideVariableBlock bool = false

func handleLine(input string) {
	inputWithoutWhiteSpaces := strings.ReplaceAll(input, " ", "")
	if parserIsInsideVariableBlock {

	}

	if inputWithoutWhiteSpaces == "VAR_INPUT" ||
		inputWithoutWhiteSpaces == "VAR" ||
		inputWithoutWhiteSpaces == "VAR_TEMP" {
		parserIsInsideVariableBlock = true
		return
	}

	if inputWithoutWhiteSpaces == "END_VAR" {
		parserIsInsideVariableBlock = false
	}
	return
}

func main() {
	if len(arguments) < 2 {
		fmt.Printf("Command needs an argument.\n")
		os.Exit(1)
	}
	var fileName string = arguments[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %s.\n", err.Error())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		handleLine(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %s.\n", err.Error())
	}

	file.Close()
	os.Exit(0)
}
