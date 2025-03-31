package main

import (
	"testing"
)

func resetState() {
	parserIsInsideVariableBlock = false
	currentBlock = None
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
	currentBlock = Variable
	handleLine("errorActive1 : Bool;")
	if isErrorDetected {
		test.Error("mixed camel case must be valid")
	}
}

func TestMixedCamelCaseFunctionIsValid(test *testing.T) {
	test.Cleanup(resetState)
	parserIsInsideVariableBlock = true
	currentBlock = Variable
	handleLine("almError {InstructionName := 'Program_Alarm'; LibVersion := '1.0'} : Program_Alarm;")
	if isErrorDetected {
		test.Error("mixed camel case must be valid")
	}
}

func TestVariableBlockInitiallyNotDetected(test *testing.T) {
	test.Cleanup(resetState)
	if parserIsInsideVariableBlock || currentBlock == Constant || currentBlock == Variable {
		test.Error("initial value of inside variable block should be false")
	}
}

func TestSnakeCaseConstantsAreValid(test *testing.T) {
	test.Cleanup(resetState)
	handleLine("VAR_GLOBAL CONSTANT")
	if currentBlock != Constant {
		test.Error("did not detect that VAR_GLOBAL CONSTANT started")
	}
}

func TestDoubleOpeningVariableBlockLeadsToError(test *testing.T) {
	test.Cleanup(resetState)
	handleLine("VAR_GLOBAL")
	handleLine("VAR_GLOBAL")
	if isErrorDetected != true {
		test.Error("double variable block should lead to an error")
	}
}

func TestVarInputTriggersInsideVariableBlock(test *testing.T) {
	test.Cleanup(resetState)
	handleLine("VAR_INPUT")
	if currentBlock != Variable {
		test.Error("did not detect that VAR_INPUT started")
	}
}
