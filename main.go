package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var arguments []string = os.Args
var parserIsInsideVariableBlock bool = false
var isErrorDetected bool = false

func isCamelCaseAllowNumbers(input string) bool {
	expression := regexp.MustCompile(`^[a-z]+([A-Z][a-z0-9]*)*$`)
	return expression.MatchString(input)
}

func isCamelCase(input string) bool {
	expression := regexp.MustCompile(`^[a-z]+([A-Z][a-z]*)*$`)
	return expression.MatchString(input)
}

func handleLine(input string) {
	inputWithoutWhiteSpaces := strings.ReplaceAll(input, " ", "")
	if parserIsInsideVariableBlock {
		if strings.Contains(inputWithoutWhiteSpaces, "{") {
			parts := strings.Split(inputWithoutWhiteSpaces, "{")
			if isCamelCaseAllowNumbers(parts[0]) {
				fmt.Printf("Success: Variable %s is camel case.\n", parts[0])
			} else {
				fmt.Printf("Error: Variable %s is not camel case.\n", parts[0])
				isErrorDetected = true
			}
			return
		}
		if strings.Contains(inputWithoutWhiteSpaces, ":") {
			parts := strings.Split(inputWithoutWhiteSpaces, ":")
			if isCamelCaseAllowNumbers(parts[0]) {
				fmt.Printf("Success: Variable %s is camel case.\n", parts[0])
			} else {
				fmt.Printf("Error: Variable %s is not camel case.\n", parts[0])
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
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		handleLine(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %s.\n", err.Error())
		isErrorDetected = true
	}

	file.Close()
	if isErrorDetected == true {
		os.Exit(1)
	}
	os.Exit(0)
}
