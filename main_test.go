package main

import (
	"testing"
)

func resetState() {
	currentBlock = None
}

func TestDetectsWhenInVariableBlock(test *testing.T) {
	test.Cleanup(resetState)
	handleLine("errorActive1 : Bool;")
	if currentBlock == Variable {
		test.Error("mixed camel case must be valid")
	}
}

func TestMixedCamelCaseVariableIsValid(test *testing.T) {
	test.Cleanup(resetState)
	currentBlock = Variable
	handleLine("errorActive1 : Bool;")
	if isErrorDetected {
		test.Error("mixed camel case must be valid")
	}
}

func TestMixedCamelCaseFunctionIsValid(test *testing.T) {
	test.Cleanup(resetState)
	currentBlock = Variable
	handleLine("almError {InstructionName := 'Program_Alarm'; LibVersion := '1.0'} : Program_Alarm;")
	if isErrorDetected {
		test.Error("mixed camel case must be valid")
	}
}

func TestVariableBlockInitiallyNotDetected(test *testing.T) {
	test.Cleanup(resetState)
	if currentBlock == Constant || currentBlock == Variable {
		test.Error("initial value of inside variable block should be false")
	}
}

func TestConstantBlockRecognized(test *testing.T) {
	test.Cleanup(resetState)
	handleLine("VAR_GLOBAL CONSTANT")
	if currentBlock != Constant {
		test.Error("did not detect that VAR_GLOBAL CONSTANT started")
	}
}

func TestSnakeCaseConstantsAreValid(test *testing.T) {
	test.Cleanup(resetState)
	handleLine("VAR_GLOBAL CONSTANT")
	handleLine("      MAX_HEIGHT      : INT := 100;")
	if isErrorDetected == true {
		test.Error("constants should be in mixed capital snake case")
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
