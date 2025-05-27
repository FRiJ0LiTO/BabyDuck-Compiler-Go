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
	Position      int
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

// HasFunction checks if a function with the given name exists in the function directory.
//
// Parameters:
//   - functionName (string): The name of the function to check.
//
// Returns:
//   - bool: True if the function exists in the directory, false otherwise.
func (fd *FunctionDirectory) HasFunction(functionName string) bool {
	_, exists := fd.FunctionsDirectory[functionName]
	return exists
}

// AddResource adds or updates the resource count for a specific variable type in a given function.
//
// Parameters:
//   - functionName (string): The name of the function whose resources are being updated.
//   - variableType (string): The type of the variable for which the resource count is being added or updated.
//   - count (int): The number of resources to add to the existing count.
//
// Behavior:
//   - If the `Resources` map for the specified function is nil, it initializes the map.
//   - Retrieves the current count for the specified variable type.
//   - Updates the count by adding the provided `count` value to the existing count.
func (fd *FunctionDirectory) AddResource(functionName string, variableType string, count int) {
	if fd.FunctionsDirectory[functionName].Resources == nil {
		fd.FunctionsDirectory[functionName].Resources = make(map[string]int)
	}
	prevCounter := fd.FunctionsDirectory[functionName].Resources[variableType]
	fd.FunctionsDirectory[functionName].Resources[variableType] = prevCounter + count
}
