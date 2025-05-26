package virtualMachine

import (
	"BabyDuck/internal/memory"
	"BabyDuck/internal/symbol"
	"fmt"
	"strconv"
)

// ResourceType defines the different types of memory segments
type ResourceType string

const (
	GlobalInt   ResourceType = "Globalint"
	GlobalFloat ResourceType = "Globalfloat"
	LocalInt    ResourceType = "Localint"
	LocalFloat  ResourceType = "Localfloat"
	TempInt     ResourceType = "Tempint"
	TempFloat   ResourceType = "Tempfloat"
	TempBool    ResourceType = "Tempbool"
)

// DynamicResources manages memory slices dynamically based on configuration
type DynamicResources struct {
	slices map[ResourceType]interface{}
}

// NewDynamicResources creates a new dynamic resources manager
func NewDynamicResources(resources map[string]int) *DynamicResources {
	dr := &DynamicResources{
		slices: make(map[ResourceType]interface{}),
	}

	// Map memory.DataType to ResourceType and create slices dynamically
	typeMapping := map[string]ResourceType{
		"Globalint":   GlobalInt,
		"Globalfloat": GlobalFloat,
		"Localint":    LocalInt,
		"Localfloat":  LocalFloat,
		"Tempint":     TempInt,
		"Tempfloat":   TempFloat,
		"Tempbool":    TempBool,
	}

	for dataType, count := range resources {
		if resourceType, exists := typeMapping[dataType]; exists {
			dr.createSliceForType(resourceType, count)
		}
	}

	return dr
}

// createSliceForType creates a slice of the appropriate type with given capacity
func (dr *DynamicResources) createSliceForType(resourceType ResourceType, capacity int) {
	switch resourceType {
	case GlobalInt, LocalInt, TempInt:
		dr.slices[resourceType] = make([]int, 0, capacity)
	case GlobalFloat, LocalFloat, TempFloat:
		dr.slices[resourceType] = make([]float64, 0, capacity)
	case TempBool:
		dr.slices[resourceType] = make([]bool, 0, capacity)
	}
}

// GetInts Type-safe getters
func (dr *DynamicResources) GetInts(resourceType ResourceType) []int {
	if slice, exists := dr.slices[resourceType]; exists {
		if intSlice, ok := slice.([]int); ok {
			return intSlice
		}
	}
	return nil
}

func (dr *DynamicResources) GetFloats(resourceType ResourceType) []float64 {
	if slice, exists := dr.slices[resourceType]; exists {
		if floatSlice, ok := slice.([]float64); ok {
			return floatSlice
		}
	}
	return nil
}

func (dr *DynamicResources) GetBools(resourceType ResourceType) []bool {
	if slice, exists := dr.slices[resourceType]; exists {
		if boolSlice, ok := slice.([]bool); ok {
			return boolSlice
		}
	}
	return nil
}

// Memory represents the virtual machine's memory management system
type Memory struct {
	VariableStorage *DynamicResources
	constantsTable  map[int]symbol.Constant
	memoryConfig    memory.Configuration
}

// AddressInfo contains information about a memory address
type AddressInfo struct {
	ResourceType ResourceType
	Index        int
	IsValid      bool
}

// NewMemory creates and initializes a new memory management instance
func NewMemory(resources map[string]int, config memory.Configuration,
	constants map[int]symbol.Constant) *Memory {
	return &Memory{
		VariableStorage: NewDynamicResources(resources),
		constantsTable:  constants,
		memoryConfig:    config,
	}
}

// getAddressInfo determines which resource type and index an address corresponds to
func (m *Memory) getAddressInfo(address int) AddressInfo {
	config := m.memoryConfig

	switch {
	case address >= config.GlobalIntStart && address <= config.GlobalIntEnd:
		return AddressInfo{
			ResourceType: GlobalInt,
			Index:        address - config.GlobalIntStart,
			IsValid:      true,
		}
	case address >= config.LocalIntStart && address <= config.LocalIntEnd:
		return AddressInfo{
			ResourceType: LocalInt,
			Index:        address - config.LocalIntStart,
			IsValid:      true,
		}
	case address >= config.LocalFloatStart && address <= config.LocalFloatEnd:
		return AddressInfo{
			ResourceType: LocalFloat,
			Index:        address - config.LocalFloatStart,
			IsValid:      true,
		}
	case address >= config.GlobalFloatStart && address <= config.GlobalFloatEnd:
		return AddressInfo{
			ResourceType: GlobalFloat,
			Index:        address - config.GlobalFloatStart,
			IsValid:      true,
		}
	case address >= config.TempIntStart && address <= config.TempIntEnd:
		return AddressInfo{
			ResourceType: TempInt,
			Index:        address - config.TempIntStart,
			IsValid:      true,
		}
	case address >= config.TempFloatStart && address <= config.TempFloatEnd:
		return AddressInfo{
			ResourceType: TempFloat,
			Index:        address - config.TempFloatStart,
			IsValid:      true,
		}
	case address >= config.TempBoolStart && address <= config.TempBoolEnd:
		return AddressInfo{
			ResourceType: TempBool,
			Index:        address - config.TempBoolStart,
			IsValid:      true,
		}
	case address >= config.ConstIntStart && address <= config.ConstIntEnd:
		return AddressInfo{
			IsValid: true,
		}
	case address >= config.ConstFloatStart && address <= config.ConstFloatEnd:
		return AddressInfo{
			IsValid: true,
		}
	case address >= config.ConstStringStart && address <= config.ConstStringEnd:
		return AddressInfo{
			IsValid: true,
		}
	default:
		return AddressInfo{IsValid: false}
	}
}

// Set stores a value at the specified memory address
func (m *Memory) Set(address int, value interface{}) error {
	addressInfo := m.getAddressInfo(address)
	if !addressInfo.IsValid {
		return fmt.Errorf("invalid memory address: %d - address outside valid ranges", address)
	}

	switch v := value.(type) {
	case int:
		return m.setIntValue(addressInfo.ResourceType, addressInfo.Index, v)
	case float64:
		return m.setFloatValue(addressInfo.ResourceType, addressInfo.Index, v)
	case bool:
		return m.setBoolValue(addressInfo.ResourceType, addressInfo.Index, v)
	default:
		return fmt.Errorf("unsupported value type: %T", value)
	}
}

// Helper methods for setting different types
func (m *Memory) setIntValue(resourceType ResourceType, index int, value int) error {
	if slice, exists := m.VariableStorage.slices[resourceType]; exists {
		if intSlice, ok := slice.([]int); ok {
			// Extend slice if necessary
			for len(intSlice) <= index {
				intSlice = append(intSlice, 0)
			}
			intSlice[index] = value
			m.VariableStorage.slices[resourceType] = intSlice
			return nil
		}
		return fmt.Errorf("resource type %s is not a float slice", resourceType)
	}
	return fmt.Errorf("resource type %s does not exist", resourceType)
}

func (m *Memory) setFloatValue(resourceType ResourceType, index int, value float64) error {
	if slice, exists := m.VariableStorage.slices[resourceType]; exists {
		if floatSlice, ok := slice.([]float64); ok {
			// Extend slice if necessary
			for len(floatSlice) <= index {
				floatSlice = append(floatSlice, 0.0)
			}
			floatSlice[index] = value
			m.VariableStorage.slices[resourceType] = floatSlice
			return nil
		}
		return fmt.Errorf("resource type %s is not a float slice", resourceType)
	}
	return fmt.Errorf("resource type %s does not exist", resourceType)
}

func (m *Memory) setBoolValue(resourceType ResourceType, index int, value bool) error {
	if slice, exists := m.VariableStorage.slices[resourceType]; exists {
		if boolSlice, ok := slice.([]bool); ok {
			// Extend slice if necessary
			for len(boolSlice) <= index {
				boolSlice = append(boolSlice, false)
			}
			boolSlice[index] = value
			m.VariableStorage.slices[resourceType] = boolSlice
			return nil
		}
		return fmt.Errorf("resource type %s is not a bool slice", resourceType)
	}
	return fmt.Errorf("resource type %s does not exist", resourceType)
}

// Get retrieves a value from the specified memory address
func (m *Memory) Get(address int) (interface{}, error) {
	addressInfo := m.getAddressInfo(address)
	if !addressInfo.IsValid {
		return nil, fmt.Errorf("invalid memory address: %d - address outside valid ranges", address)
	}

	// Check constants table first
	if constantValue, isConstant := m.getConstantValue(address); isConstant {
		return constantValue, nil
	}

	// Get value from variable storage
	return m.getVariableValue(addressInfo)
}

// getVariableValue retrieves a value from variable storage based on address info
func (m *Memory) getVariableValue(addressInfo AddressInfo) (interface{}, error) {
	switch addressInfo.ResourceType {
	case GlobalInt, LocalInt, TempInt:
		if slice := m.VariableStorage.GetInts(addressInfo.ResourceType); slice != nil && addressInfo.Index < len(slice) {
			return slice[addressInfo.Index], nil
		}
		return 0, nil // Default value

	case GlobalFloat, LocalFloat, TempFloat:
		if slice := m.VariableStorage.GetFloats(addressInfo.ResourceType); slice != nil && addressInfo.Index < len(slice) {
			return slice[addressInfo.Index], nil
		}
		return 0.0, nil // Default value

	case TempBool:
		if slice := m.VariableStorage.GetBools(addressInfo.ResourceType); slice != nil && addressInfo.Index < len(slice) {
			return slice[addressInfo.Index], nil
		}
		return false, nil // Default value

	default:
		return nil, fmt.Errorf("unsupported resource type: %s", addressInfo.ResourceType)
	}
}

// isValidAddress validates whether an address falls within any valid memory range
func (m *Memory) isValidAddress(address int) bool {
	return m.getAddressInfo(address).IsValid
}

// getConstantValue retrieves and parses a constant value if it exists at the given address
func (m *Memory) getConstantValue(address int) (interface{}, bool) {
	if constantInfo, exists := m.constantsTable[address]; exists {
		parsedValue, err := m.parseConstantByType(constantInfo.Value, constantInfo.DataType)
		if err != nil {
			fmt.Printf("Warning: Failed to parse constant at address %d: %v\n", address, err)
			return constantInfo.Value, true
		}
		return parsedValue, true
	}
	return nil, false
}

// parseConstantByType converts a string constant to its appropriate runtime type
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
		return constantString, nil
	}
}
