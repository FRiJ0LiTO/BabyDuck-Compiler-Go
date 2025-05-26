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
	ConstStringStart   int
	ConstStringEnd     int
}

type Operator int

const (
	ADD Operator = iota + 1
	SUB
	MUL
	DIV
	NEG

	GREATER
	LESS
	NOTEQUAL

	GOTOF
	GOTO

	PRINT

	FUNC
	PARAM
	ENDFUNC
	CALL

	ASSIGN

	PROGRAM
	END
)

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
	ConstStringStart:   12000,
	ConstStringEnd:     12999,
}

func IdentifyOperator(operator string) Operator {
	switch operator {
	case "+":
		return ADD
	case "-":
		return SUB
	case "*":
		return MUL
	case "/":
		return DIV
	case "NEG":
		return NEG
	case ">":
		return GREATER
	case "<":
		return LESS
	case "!=":
		return NOTEQUAL
	case "GOTOF":
		return GOTOF
	case "GOTO":
		return GOTO
	case "PRINT":
		return PRINT
	case "FUNC":
		return FUNC
	case "PARAM":
		return PARAM
	case "ENDFUNC":
		return ENDFUNC
	case "CALL":
		return CALL
	case "=":
		return ASSIGN
	case "PROGRAM":
		return PROGRAM
	case "END":
		return END
	default:
		return -1
	}
}
