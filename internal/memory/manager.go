package memory

import "fmt"

type VarType string

type DataType string

const (
	Global   VarType = "Global"
	Local    VarType = "Local"
	Temporal VarType = "Temporal"
	Constant VarType = "Constant"
)

const (
	Integer     DataType = "int"
	Float       DataType = "float"
	Boolean     DataType = "bool"
	String      DataType = "string"
	TempInt     DataType = "TempInt"
	TempFloat   DataType = "TempFloat"
	TempBoolean DataType = "TempBoolean"
)

// Range represents a memory range with a start, end, and the next available address.
// It is used to manage memory allocation for different variable types and data types.
type Range struct {
	Start int // The starting address of the memory range.
	End   int // The ending address of the memory range.
	Next  int // The next available address within the memory range.
}

// NewMemoryRange creates and initializes a new memory range.
//
// Parameters:
//   - start (int): The starting address of the memory range.
//   - end (int): The ending address of the memory range.
//
// Returns:
//   - (*Range): A pointer to the newly created Range instance.
//
// The `Next` field is initialized to the same value as `Start`.
func NewMemoryRange(start, end int) *Range {
	return &Range{
		Start: start,
		End:   end,
		Next:  start,
	}
}

// Manager manages memory ranges for different variable types and data types.
// It provides functionality to allocate and reset memory addresses.
type Manager struct {
	Ranges map[VarType]map[DataType]*Range
}

// NewMemoryManager initializes a new memory manager with memory ranges for different variable types and data types.
//
// Parameters:
//   - conf (Configuration): A configuration object containing the start and end addresses for each memory range.
//
// Returns:
//   - (*Manager): A pointer to the newly created memory manager.
//
// The function creates a `Manager` instance and initializes memory ranges for the following variable types:
//   - Global: Memory ranges for global variables.
//   - Local: Memory ranges for local variables.
//   - Temporal: Memory ranges for temporary variables.
//   - Constant: Memory ranges for constants.
//
// Each variable type is further divided into data types (Integer, Float, Boolean), and their respective memory ranges
// are initialized using the `NewMemoryRange` function.
func NewMemoryManager(conf Configuration) *Manager {
	m := &Manager{
		Ranges: make(map[VarType]map[DataType]*Range),
	}

	m.Ranges[Global] = make(map[DataType]*Range)
	m.Ranges[Local] = make(map[DataType]*Range)
	m.Ranges[Temporal] = make(map[DataType]*Range)
	m.Ranges[Constant] = make(map[DataType]*Range)

	m.Ranges[Global][Integer] = NewMemoryRange(conf.GlobalIntStart, conf.GlobalIntEnd)
	m.Ranges[Global][Float] = NewMemoryRange(conf.GlobalFloatStart, conf.GlobalFloatEnd)
	m.Ranges[Global][Boolean] = NewMemoryRange(conf.GlobalBoolStart, conf.GlobalBoolEnd)

	m.Ranges[Local][Integer] = NewMemoryRange(conf.LocalIntStart, conf.LocalIntEnd)
	m.Ranges[Local][Float] = NewMemoryRange(conf.LocalFloatStart, conf.LocalFloatEnd)
	m.Ranges[Local][Boolean] = NewMemoryRange(conf.LocalBoolStart, conf.LocalBoolEnd)

	m.Ranges[Temporal][Integer] = NewMemoryRange(conf.TempIntStart, conf.TempIntEnd)
	m.Ranges[Temporal][Float] = NewMemoryRange(conf.TempFloatStart, conf.TempFloatEnd)
	m.Ranges[Temporal][Boolean] = NewMemoryRange(conf.TempBoolStart, conf.TempBoolEnd)

	m.Ranges[Constant][Integer] = NewMemoryRange(conf.ConstIntStart, conf.ConstIntEnd)
	m.Ranges[Constant][Float] = NewMemoryRange(conf.ConstFloatStart, conf.ConstFloatEnd)
	m.Ranges[Constant][String] = NewMemoryRange(conf.ConstStringStart, conf.ConstStringEnd)

	return m
}

// GetAddress retrieves the next available memory address for a given variable type and data type.
//
// Parameters:
//   - varType (VarType): The type of variable (e.g., Global, Local, Temporal, Constant).
//   - dataType (DataType): The data type of the variable (e.g., Integer, Float, Boolean).
//
// Returns:
//   - (int): The next available memory address for the specified variable and data type.
//   - (error): An error if the memory range does not exist or if there is no available memory.
//
// Errors:
//   - Returns an error if the memory range for the specified varType and dataType does not exist.
//   - Returns an error if the memory range is exhausted (i.e., no more addresses are available).
func (m *Manager) GetAddress(varType VarType, dataType DataType) (int, error) {
	memRange, exists := m.Ranges[varType][dataType]
	if !exists {
		return -1, fmt.Errorf("rango de memoria no encontrado para tipo %s y dato %s", varType, dataType)
	}

	// Check if there is availability
	if memRange.Next > memRange.End {
		return -1, fmt.Errorf("error: out of memory for varType %s and dataType %s", varType, dataType)
	}

	addr := memRange.Next
	memRange.Next++
	return addr, nil
}

// ResetLocal resets the memory ranges for local and temporal variables.
// It sets the `Next` field of each memory range back to its `Start` value,
// effectively clearing the memory allocation for these variable types.
func (m *Manager) ResetLocal() {
	for dataType, memRange := range m.Ranges[Local] {
		m.Ranges[Local][dataType].Next = memRange.Start
	}

	for dataType, memRange := range m.Ranges[Temporal] {
		m.Ranges[Temporal][dataType].Next = memRange.Start
	}
}
