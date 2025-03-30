package main

import (
	"testing"
)

func TestMixedCamelCaseVariableIsValid(test *testing.T) {
	parserIsInsideVariableBlock = true
	handleLine("errorActive1 : Bool;")
	if isErrorDetected {
		test.Error("mixed camel case must be valid")
	}
}
