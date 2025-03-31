package main

import (
	"testing"
)

func TestDetectsWhenInVariableBlock(test *testing.T) {
	handleLine("errorActive1 : Bool;")
	if parserIsInsideVariableBlock {
		test.Error("mixed camel case must be valid")
	}
}

func TestMixedCamelCaseVariableIsValid(test *testing.T) {
	parserIsInsideVariableBlock = true
	handleLine("errorActive1 : Bool;")
	if isErrorDetected {
		test.Error("mixed camel case must be valid")
	}
}

func TestMixedCamelCaseFunctionIsValid(test *testing.T) {
	parserIsInsideVariableBlock = true
	handleLine("almError {InstructionName := 'Program_Alarm'; LibVersion := '1.0'} : Program_Alarm;")
	if isErrorDetected {
		test.Error("mixed camel case must be valid")
	}
}
