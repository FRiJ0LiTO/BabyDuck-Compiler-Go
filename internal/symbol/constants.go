package symbol

import (
	"fmt"
)

type Constant struct {
	virtualDirection int
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
func (fd *FunctionDirectory) AddConstant(constant string, virtualAddress int) error {
	if fd.CurrentScope.Height == 0 {
		return fmt.Errorf("error: no active scope defined")
	}

	if _, exists := fd.Constants[constant]; exists {
		return nil
	}

	fd.Constants[constant] = Constant{
		virtualDirection: virtualAddress,
	}
	return nil
}
