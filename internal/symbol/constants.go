package symbol

import (
	"BabyDuck/internal/memory"
	"fmt"
)

type Constant struct {
	Value    string
	DataType memory.DataType
}

// AddConstant adds a new constant to the function directory.
//
// Parameters:
//   - constant (string): The name of the constant to be added.
//   - virtualAddress (int): The virtual memory address associated with the constant.
//
// Returns:
//   - (error): An error if there is no active scope defined, otherwise nil.
//
// Behavior:
//   - If there is no active scope (i.e., CurrentScope.Height is 0), an error is returned.
//   - If the constant already exists in the Constants map, the function does nothing and returns nil.
//   - Otherwise, the constant is added to the Constants map with the provided virtual address.
func (fd *FunctionDirectory) AddConstant(constantValue string, constantType memory.DataType, virtualAddress int) error {
	if fd.CurrentScope.Height == 0 {
		return fmt.Errorf("error: no active scope defined")
	}

	existingAddress, _ := fd.LookupConstant(constantValue)
	if existingAddress != -1 {
		return nil
	}

	fd.Constants[virtualAddress] = Constant{
		Value:    constantValue,
		DataType: constantType,
	}
	return nil
}

// LookupConstant retrieves the virtual memory address of a constant from the function directory.
//
// Parameters:
//   - constantValue (string): The value of the constant to look up.
//
// Returns:
//   - (int): The virtual memory address of the constant if found, or -1 if the constant does not exist.
//
// Behavior:
//   - If the constant is found in the Constants map, its virtual memory address is returned.
//   - If the constant is not found, the function returns -1.
func (fd *FunctionDirectory) LookupConstant(constantValue string) (int, memory.DataType) {
	for virtualAddress, constantInfo := range fd.Constants {
		if constantValue == constantInfo.Value {
			return virtualAddress, constantInfo.DataType
		}
	}

	return -1, ""
}
