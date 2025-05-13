package symbol

import (
	"BabyDuck/internal/memory"
	"fmt"
)

// Variable represents a variable in the symbol table with its type and value.
// Fields:
//   - variableType (string): The data type of the variable.
//   - VariableValue (interface{}): The value of the variable, which can be of any type.
type Variable struct {
	variableType     memory.DataType
	virtualDirection int
	VariableValue    interface{}
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
func (fd *FunctionDirectory) AddVariable(name string, dataType memory.DataType, virtualAddress int) error {
	if fd.CurrentScope.Height == 0 {
		return fmt.Errorf("error: no active scope defined")
	}

	// Get the current scope from the stack.
	context := fd.CurrentScope.Peek()
	varTable := fd.Directory[context.(string)]

	// Check if the variable already exists in the current scope.
	if _, exists := varTable[name]; exists {
		return fmt.Errorf("error: variable '%s' already declared in scope '%s'", name, context)
	}

	// Add the new variable to the variable table.
	varTable[name] = Variable{
		variableType:     dataType,
		virtualDirection: virtualAddress,
		VariableValue:    nil,
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
func (fd *FunctionDirectory) ValidateVariableExists(scopes []string, name string) error {
	for i := len(scopes) - 1; i >= 0; i-- {
		scope := scopes[i]
		_, exists := fd.LookupVariable(scope, name)
		if exists {
			return nil
		}
	}

	return fmt.Errorf("error: undefined variable '%s'", name)
}

// LookupVariable searches for a variable in a specific scope.
//
// Parameters:
//   - scope (string): The name of the scope to search in.
//   - name (string): The name of the variable to look up.
//
// Returns:
//   - string: The data type of the variable if found.
//   - bool: A boolean indicating whether the variable was found in the specified scope.
func (fd *FunctionDirectory) LookupVariable(scope string, name string) (memory.DataType, bool) {
	varTable, scopeExists := fd.Directory[scope]
	if !scopeExists {
		return "", false
	}

	variable, varExists := varTable[name]
	return variable.variableType, varExists
}
