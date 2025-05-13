package semantic

// Cube represents a semantic cube used to determine the result type of operations
// between different data types. It contains a nested map structure where the keys
// are the left operand type, right operand type, and operator, and the values are
// the resulting data type.
type Cube struct {
	semantics map[string]map[string]map[string]string // Nested map for semantic rules
}

// NewSemanticCube creates and initializes a new semantic cube with predefined
// rules for operations between "int" and "float" types. The rules specify the
// resulting type for each combination of operand types and operators.
//
// Returns:
//   Cube: The newly created semantic cube.
func NewSemanticCube() Cube {
	cube := Cube{
		semantics: map[string]map[string]map[string]string{
			"int": {
				"int": {
					"+":  "int",
					"-":  "int",
					"*":  "int",
					"/":  "int",
					"<":  "bool",
					">":  "bool",
					"!=": "bool",
				},
				"float": {
					"+":  "float",
					"-":  "float",
					"*":  "float",
					"/":  "float",
					"<":  "bool",
					">":  "bool",
					"!=": "bool",
				},
			},
			"float": {
				"int": {
					"+":  "float",
					"-":  "float",
					"*":  "float",
					"/":  "float",
					"<":  "bool",
					">":  "bool",
					"!=": "bool",
				},
				"float": {
					"+":  "float",
					"-":  "float",
					"*":  "float",
					"/":  "float",
					"<":  "bool",
					">":  "bool",
					"!=": "bool",
				},
			},
		},
	}
	return cube
}

// GetResultType retrieves the result type of operation based on the left operand type,
// right operand type, and operator. It uses the semantic cube to look up the result type.
//
// Parameters:
//   leftType (string): The type of the left operand.
//   rightType (string): The type of the right operand.
//   operator (string): The operator being applied.
//
// Returns:
//   string: The resulting type of the operation if found.
//   bool: A boolean indicating whether the operation is valid (true if valid, false otherwise).
func (sc *Cube) GetResultType(leftType string, rightType string, operator string) (string, bool) {
	if resultType, ok := sc.semantics[leftType][rightType][operator]; ok {
		return resultType, true
	}
	return "", false
}
