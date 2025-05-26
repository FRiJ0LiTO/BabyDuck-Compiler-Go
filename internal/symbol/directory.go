package symbol

import (
	"BabyDuck/internal/memory"
	"BabyDuck/structures/stack"
	"fmt"
)

type FunctionInformation struct {
	VariableTable VariableTable
	Parameters    []memory.DataType
	Resources     map[string]int
}

// FunctionDirectory manages the symbol tables for different scopes in a program.
// It provides functionality to create and manage scopes and their associated variable tables.
type FunctionDirectory struct {
	FunctionsDirectory map[string]*FunctionInformation // Maps scope names to their variable tables.
	Constants          map[int]Constant                //Maps constants.
	CurrentScope       *stack.Stack                    // Stack of active scopes, with the innermost scope at the end.
}

// NewFunctionDirectory creates and initializes a new function directory with a global "program" scope.
//
// Returns:
//
//	*FunctionDirectory: A pointer to the newly created function directory.
func NewFunctionDirectory() *FunctionDirectory {
	directory := &FunctionDirectory{
		FunctionsDirectory: make(map[string]*FunctionInformation),
		Constants:          make(map[int]Constant),
		CurrentScope:       stack.New(),
	}

	// Initialize the global "program" scope with an empty variable table.
	directory.FunctionsDirectory["program"] = &FunctionInformation{
		VariableTable: make(VariableTable),
		Parameters:    make([]memory.DataType, 0),
		Resources:     make(map[string]int),
	}
	directory.CurrentScope.Push("program")
	return directory
}

// AddFunction creates a new scope with the given function name.
//
// Parameters:
//
//	functionName (string): The name of the function to be added as a new scope.
//
// Returns:
//
//	error: An error if a function with the same name already exists in the directory, otherwise nil.
func (fd *FunctionDirectory) AddFunction(functionName string) error {
	// Check if the function name already exists in the directory.
	if _, exists := fd.FunctionsDirectory[functionName]; exists {
		return fmt.Errorf("error: function '%s' already declared in current scope", functionName)
	}

	// Create a new variable table for the function scope.
	fd.FunctionsDirectory[functionName] = &FunctionInformation{
		VariableTable: make(VariableTable),
		Parameters:    make([]memory.DataType, 0),
		Resources:     make(map[string]int),
	}

	return nil
}

func (fd *FunctionDirectory) AddResource(functionName string, variableType string, count int) {
	if fd.FunctionsDirectory[functionName].Resources == nil {
		fd.FunctionsDirectory[functionName].Resources = make(map[string]int)
	}
	prevCounter := fd.FunctionsDirectory[functionName].Resources[variableType]
	fd.FunctionsDirectory[functionName].Resources[variableType] = prevCounter + count
}
