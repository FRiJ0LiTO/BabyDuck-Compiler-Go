package memory

// Configuration defines the memory address ranges for different variable types and data types.
// Each field represents the start and end of a memory range for a specific combination of variable type
// (e.g., Global, Local, Temporal, Constant) and data type (e.g., Integer, Float, Boolean).
type Configuration struct {
	GlobalIntStart   int
	GlobalIntEnd     int
	GlobalFloatStart int
	GlobalFloatEnd   int
	GlobalBoolStart  int
	GlobalBoolEnd    int
	LocalIntStart    int
	LocalIntEnd      int
	LocalFloatStart  int
	LocalFloatEnd    int
	LocalBoolStart   int
	LocalBoolEnd     int
	TempIntStart     int
	TempIntEnd       int
	TempFloatStart   int
	TempFloatEnd     int
	TempBoolStart    int
	TempBoolEnd      int
	ConstIntStart    int
	ConstIntEnd      int
	ConstFloatStart  int
	ConstFloatEnd    int
	ConstBoolStart   int
	ConstBoolEnd     int
}

// DefaultMemoryConfig provides the default memory address ranges for all variable types and data types.
// These ranges are used to initialize the memory manager with predefined values.
var DefaultMemoryConfig = Configuration{
	GlobalIntStart:   1000,
	GlobalIntEnd:     1999,
	GlobalFloatStart: 2000,
	GlobalFloatEnd:   2999,
	GlobalBoolStart:  3000,
	GlobalBoolEnd:    3999,
	LocalIntStart:    4000,
	LocalIntEnd:      4999,
	LocalFloatStart:  5000,
	LocalFloatEnd:    5999,
	LocalBoolStart:   6000,
	LocalBoolEnd:     6999,
	TempIntStart:     7000,
	TempIntEnd:       7999,
	TempFloatStart:   8000,
	TempFloatEnd:     8999,
	TempBoolStart:    9000,
	TempBoolEnd:      9999,
	ConstIntStart:    10000,
	ConstIntEnd:      10999,
	ConstFloatStart:  11000,
	ConstFloatEnd:    11999,
	ConstBoolStart:   12000,
	ConstBoolEnd:     12999,
}
