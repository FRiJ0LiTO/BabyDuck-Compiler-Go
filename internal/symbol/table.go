package symbol

import (
	"BabyDuck/internal/memory"
	"fmt"
)

// Variable represents a variable in the symbol table with its type and value.
// Fields:
//   - variableType (string): The data type of the variable.
type Variable struct {
	VariableType     memory.DataType
	VirtualDirection int
}

// VariableTable is a map of variable names to their corresponding Variable objects.
// It serves as the symbol table for a specific scope.
type VariableTable map[string]Variable

// AddVariable adds a new variable to the current scope's variable table.
//
// Parameters:
//   - name (string): The name of the variable to be added.
//   - varType (string): The data type of the variable.
//
// Returns:
//   - error: An error if the variable is already declared in the current scope or if no active scope exists.
func (fd *FunctionDirectory) AddVariable(name string, dataType memory.DataType,
	virtualAddress int, scope string) error {

	varTable := fd.FunctionsDirectory[scope].VariableTable

	// Check if the variable already exists in the current scope.
	if _, exists := varTable[name]; exists {
		return fmt.Errorf("error: variable '%s' already declared in scope '%s'", name, scope)
	}

	// Add the new variable to the variable table.
	varTable[name] = Variable{
		VariableType:     dataType,
		VirtualDirection: virtualAddress,
	}

	return nil
}

// ValidateVariableExists checks if a variable exists in the current scope hierarchy.
//
// Parameters:
//   - scopes ([]string): A list of scopes to search for the variable.
//   - name (string): The name of the variable to validate.
//
// Returns:
//   - error: An error if the variable is not found in any accessible scope.
func (fd *FunctionDirectory) ValidateVariableExists(scopes []string, name string) (Variable, bool) {
	for i := len(scopes) - 1; i >= 0; i-- {
		scope := scopes[i]
		variable, exists := fd.LookupVariable(scope, name)
		if exists {
			return variable, exists
		}
	}
	return Variable{}, false

	//return fmt.Errorf("error: undefined variable '%s'", name)
}

// LookupVariable searches for a variable in a specific scope.
//
// Parameters:
//   - scope (string): The name of the scope to search in.
//   - name (string): The name of the variable to look up.
//
// Returns:
//   - string: The data type of the variable if found.
//   - int: Virtual address of the variable if found.
//   - bool: A boolean indicating whether the variable was found in the specified scope.
func (fd *FunctionDirectory) LookupVariable(scope string, name string) (Variable, bool) {
	function, scopeExists := fd.FunctionsDirectory[scope]
	if !scopeExists {
		return Variable{}, false
	}

	variable, varExists := function.VariableTable[name]
	return variable, varExists
}

// LookupVariableByVirtualAddress searches for a variable by its virtual address in a specific scope.
// Parameters:
//   - scopes ([]string): A list of scope names to search for the variable.
//   - targetAddress (int): The virtual memory address to search for.
//
// Returns:
//   - memory.DataType: The data type of the variable if found.
//   - bool: A boolean indicating whether a variable with the given virtual address was found.
func (fd *FunctionDirectory) LookupVariableByVirtualAddress(scopes []string, targetAddress int) (memory.DataType, bool) {

	// Check if the target address exists in the constants table.
	if constant, exists := fd.Constants[targetAddress]; exists {
		return constant.DataType, true
	}

	// Iterate through the scopes in reverse order to search for the variable.
	for i := len(scopes) - 1; i >= 0; i-- {
		scope := scopes[i]

		// Retrieve the variable table for the current scope.
		variableTable, exists := fd.FunctionsDirectory[scope]
		if exists != false {
			// Iterate through the variables in the table to find a match by virtual address.
			for _, variableInfo := range variableTable.VariableTable {
				if variableInfo.VirtualDirection == targetAddress {
					return variableInfo.VariableType, true
				}
			}
		}
	}

	return "bool", false
}
