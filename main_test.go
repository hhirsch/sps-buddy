package main

import (
	"testing"
)

func resetState() {
	parserIsInsideVariableBlock = false
}

func TestDetectsWhenInVariableBlock(test *testing.T) {
	test.Cleanup(resetState)
	handleLine("errorActive1 : Bool;")
	if parserIsInsideVariableBlock {
		test.Error("mixed camel case must be valid")
	}
}

func TestMixedCamelCaseVariableIsValid(test *testing.T) {
	test.Cleanup(resetState)
	parserIsInsideVariableBlock = true
	handleLine("errorActive1 : Bool;")
	if isErrorDetected {
		test.Error("mixed camel case must be valid")
	}
}

func TestMixedCamelCaseFunctionIsValid(test *testing.T) {
	test.Cleanup(resetState)
	parserIsInsideVariableBlock = true
	handleLine("almError {InstructionName := 'Program_Alarm'; LibVersion := '1.0'} : Program_Alarm;")
	if isErrorDetected {
		test.Error("mixed camel case must be valid")
	}
}

func TestVarInputTriggersInsideVariableBlock(test *testing.T) {
	test.Cleanup(resetState)
	if parserIsInsideVariableBlock {
		test.Error("initial value of inside variable block should be false")
	}
	handleLine("VAR_INPUT")
	if !parserIsInsideVariableBlock {
		test.Error("did not detect that VAR_INPUT started")
	}
}
