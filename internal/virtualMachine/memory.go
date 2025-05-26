package virtualMachine

import (
	"BabyDuck/internal/memory"
	"BabyDuck/internal/symbol"
	"fmt"
	"strconv"
)

// TypedValue represents a value with its associated data type.
// This structure is used for type checking and validation during execution.
type TypedValue struct {
	Value    interface{}     // The actual stored value
	DataType memory.DataType // The type of the stored value (int, float, etc.)
}

// Memory represents the virtual machine's memory management system.
// It handles storage, retrieval, and validation of variables and constants
// across different memory segments (global, local, temporal, constant).
type Memory struct {
	variableStorage map[int]interface{}     // Runtime storage for variables and temporaries
	constantsTable  map[int]symbol.Constant // Read-only constants with their metadata
	memoryConfig    memory.Configuration    // Memory layout configuration (address ranges)
}

// NewMemory creates and initializes a new memory management instance.
//
// Parameters:
//   - config: Memory configuration defining address ranges for different variable types
//   - constants: Pre-defined constants table from the compilation phase
//
// Returns:
//   - *Memory: Pointer to the newly created memory manager
func NewMemory(config memory.Configuration, constants map[int]symbol.Constant) *Memory {
	return &Memory{
		variableStorage: make(map[int]interface{}),
		constantsTable:  constants,
		memoryConfig:    config,
	}
}

// Set stores a value at the specified memory address.
// Validates the address before storing to prevent memory corruption.
//
// Parameters:
//   - address: Virtual memory address where the value should be stored
//   - value: The value to store (can be any type: int, float64, bool, etc.)
//
// Returns:
//   - error: nil if successful, error if address is invalid
//
// Errors:
//   - Returns error if the address is outside valid memory ranges
func (m *Memory) Set(address int, value interface{}) error {
	if !m.isValidAddress(address) {
		return fmt.Errorf("invalid memory address: %d - address outside valid ranges", address)
	}

	m.variableStorage[address] = value
	return nil
}

// Get retrieves a value from the specified memory address.
// Checks constants table first, then variable storage, returning default values if uninitialized.
//
// Parameters:
//   - address: Virtual memory address to read from
//
// Returns:
//   - interface{}: The value stored at the address, or default value if uninitialized
//   - error: nil if successful, error if address is invalid
//
// Behavior:
//   - First checks if address contains a constant value
//   - Then checks variable storage for runtime values
//   - Returns type-appropriate default value if address is uninitialized
//   - Validates address ranges before any access
func (m *Memory) Get(address int) (interface{}, error) {
	if !m.isValidAddress(address) {
		return nil, fmt.Errorf("invalid memory address: %d - address outside valid ranges", address)
	}

	// Check constants table first (constants have priority)
	constantValue, isConstant := m.getConstantValue(address)
	if isConstant {
		return constantValue, nil
	}

	// Check variable storage for runtime values
	runtimeValue, exists := m.variableStorage[address]
	if exists {
		return runtimeValue, nil
	}

	// Return default value for uninitialized variables
	return 0, nil
}

// isValidAddress validates whether an address falls within any valid memory range.
//
// Parameters:
//   - address: Virtual memory address to validate
//
// Returns:
//   - bool: true if address is within valid ranges, false otherwise
//
// Validation Rules:
//   - Address must be within global, local, temporal, or constant ranges
//   - Uses memory configuration to determine valid bounds
func (m *Memory) isValidAddress(address int) bool {
	config := m.memoryConfig

	// Check if address falls within any valid memory segment
	return address >= config.GlobalIntStart && address <= config.ConstStringEnd
}

// getConstantValue retrieves and parses a constant value if it exists at the given address.
//
// Parameters:
//   - address: Virtual memory address to check for constants
//
// Returns:
//   - interface{}: The parsed constant value if found
//   - bool: true if a constant was found at this address, false otherwise
//
// Behavior:
//   - Searches constants table for matching virtual address
//   - Parses string representation to appropriate type (int, float, string)
//   - Returns nil, false if no constant found at address
func (m *Memory) getConstantValue(address int) (interface{}, bool) {
	for constantAddress, constantInfo := range m.constantsTable {
		if constantAddress == address {
			parsedValue, err := m.parseConstantByType(constantInfo.Value, constantInfo.DataType)
			if err != nil {
				// Log error but don't crash - return string value as fallback
				fmt.Printf("Warning: Failed to parse constant at address %d: %v\n", address, err)
				return constantInfo.Value, true
			}
			return parsedValue, true
		}
	}
	return nil, false
}

// parseConstantByType converts a string constant to its appropriate runtime type.
//
// Parameters:
//   - constantString: String representation of the constant value
//   - dataType: Expected data type for the constant
//
// Returns:
//   - interface{}: The parsed value in the correct type
//   - error: nil if parsing successful, error if conversion fails
//
// Supported Types:
//   - memory.Integer: Converts to int using strconv.Atoi
//   - memory.Float: Converts to float64 using strconv.ParseFloat
//   - Default: Returns original string value
func (m *Memory) parseConstantByType(constantString string, dataType memory.DataType) (interface{}, error) {
	switch dataType {
	case memory.Integer:
		intValue, err := strconv.Atoi(constantString)
		if err != nil {
			return nil, fmt.Errorf("failed to parse integer constant '%s': %w", constantString, err)
		}
		return intValue, nil

	case memory.Float:
		floatValue, err := strconv.ParseFloat(constantString, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse float constant '%s': %w", constantString, err)
		}
		return floatValue, nil

	default:
		// String literals and other types are returned as-is
		return constantString, nil
	}
}
