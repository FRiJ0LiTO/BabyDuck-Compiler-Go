package virtualMachine

import (
	"BabyDuck/internal/memory"
	"BabyDuck/internal/semantic"
	"BabyDuck/internal/symbol"
	"BabyDuck/structures/stack"
	"fmt"
	"os"
)

// VirtualMachine represents the execution environment for BabyDuck programs.
// It manages memory, executes quadruples, and maintains execution state.
type VirtualMachine struct {
	Memory             *Memory                  // Memory management system
	FunctionsDirectory symbol.FunctionDirectory // Symbol table for variables and functions
	quadruples         []semantic.Quadruple     // List of intermediate code instructions
	executionStack     *stack.Stack             // Stack for function calls and scope management
	programCounter     int                      // Current instruction pointer
}

// NewVirtualMachine creates and initializes a new virtual machine instance.
//
// Parameters:
//   - functionsDirectory: Symbol table containing all declared functions and variables
//   - quadruples: Array of intermediate code instructions to execute
//
// Returns:
//   - *VirtualMachine: Pointer to the newly created virtual machine
func NewVirtualMachine(functionsDirectory symbol.FunctionDirectory, quadruples []semantic.Quadruple) *VirtualMachine {
	return &VirtualMachine{
		Memory: NewMemory(functionsDirectory.FunctionsDirectory["program"].Resources,
			memory.DefaultMemoryConfig, functionsDirectory.Constants),
		FunctionsDirectory: functionsDirectory,
		quadruples:         quadruples,
		executionStack:     stack.New(),
		programCounter:     0,
	}
}

// Execute runs the virtual machine, processing quadruples until completion.
// This is the main execution loop that interprets and executes each instruction.
func (vm *VirtualMachine) Execute() {
	for vm.programCounter < len(vm.quadruples) {
		currentInstruction := vm.quadruples[vm.programCounter]
		shouldIncrementPC := true

		switch currentInstruction.Operator.(memory.Operator) {
		case memory.PROGRAM:
			shouldIncrementPC = vm.handleProgramStart(currentInstruction)

		case memory.NEG:
			vm.executeNegation(currentInstruction)

		case memory.GREATER, memory.LESS, memory.NOTEQUAL:
			vm.executeComparison(currentInstruction)

		case memory.ADD, memory.SUB, memory.MUL, memory.DIV:
			vm.executeArithmetic(currentInstruction)

		case memory.ASSIGN:
			vm.executeAssignment(currentInstruction)

		case memory.PRINT:
			vm.executePrint(currentInstruction)

		case memory.GOTOF:
			shouldIncrementPC = vm.executeConditionalJump(currentInstruction)

		case memory.GOTO:
			shouldIncrementPC = vm.executeUnconditionalJump(currentInstruction)

		case memory.END:
			vm.handleProgramEnd()
			return

		default:
			fmt.Printf("Warning: Unhandled operator %v at PC %d\n", currentInstruction.Operator, vm.programCounter)
		}

		if shouldIncrementPC {
			vm.programCounter++
		}
	}
}

// handleProgramStart processes the PROGRAM instruction to jump to main function.
//
// Parameters:
//   - instruction: The PROGRAM quadruple containing the main function address
//
// Returns:
//   - bool: false to prevent normal PC increment (since we're jumping)
func (vm *VirtualMachine) handleProgramStart(instruction semantic.Quadruple) bool {
	mainFunctionAddress := instruction.Result.(int)
	vm.programCounter = mainFunctionAddress
	return false
}

// handleProgramEnd displays completion message and program statistics.
func (vm *VirtualMachine) handleProgramEnd() {
	fmt.Println("✓ BabyDuck program executed successfully.")
	fmt.Printf("  Instructions processed: %d\n", vm.programCounter+1)
	vm.Memory.VariableStorage.slices = make(map[ResourceType]interface{})
}

// executeUnconditionalJump performs an unconditional jump to the specified address.
//
// Parameters:
//   - instruction: The GOTO quadruple containing the target address
//
// Returns:
//   - bool: false to prevent normal PC increment
func (vm *VirtualMachine) executeUnconditionalJump(instruction semantic.Quadruple) bool {
	targetAddress := instruction.Result.(int)
	vm.programCounter = targetAddress
	return false
}

// executeNegation performs unary negation operation on a value.
//
// Parameters:
//   - instruction: The NEG quadruple containing operand and result addresses
func (vm *VirtualMachine) executeNegation(instruction semantic.Quadruple) {
	operandAddress := instruction.LeftOperand.(int)
	resultAddress := instruction.Result.(int)

	operandValue, err := vm.Memory.Get(operandAddress)
	if err != nil {
		panic(fmt.Sprintf("Error getting operand for negation: %v", err))
	}

	negatedResult := vm.negateValue(operandValue)

	if err := vm.Memory.Set(resultAddress, negatedResult); err != nil {
		panic(fmt.Sprintf("Error setting negation result: %v", err))
	}
}

// executeArithmetic performs binary arithmetic operations (ADD, SUB, MUL, DIV).
// Now with proper type handling that preserves types when possible.
//
// Parameters:
//   - instruction: The arithmetic quadruple containing operands and result addresses
func (vm *VirtualMachine) executeArithmetic(instruction semantic.Quadruple) {
	leftOperandAddress := instruction.LeftOperand.(int)
	rightOperandAddress := instruction.RightOperand.(int)
	resultAddress := instruction.Result.(int)

	leftValue, err := vm.Memory.Get(leftOperandAddress)
	if err != nil {
		panic(fmt.Sprintf("Error getting left operand: %v", err))
	}

	rightValue, err := vm.Memory.Get(rightOperandAddress)
	if err != nil {
		panic(fmt.Sprintf("Error getting right operand: %v", err))
	}

	operator := instruction.Operator.(memory.Operator)

	// Use type-safe arithmetic operation
	result, err := vm.performArithmetic(leftValue, rightValue, operator)
	if err != nil {
		panic(fmt.Sprintf("Arithmetic operation failed: %v", err))
	}

	if err := vm.Memory.Set(resultAddress, result); err != nil {
		panic(fmt.Sprintf("Error setting arithmetic result: %v", err))
	}
}

// performArithmetic handles arithmetic operations while preserving types when appropriate
func (vm *VirtualMachine) performArithmetic(left, right interface{}, operator memory.Operator) (interface{}, error) {
	// Handle int + int = int operations
	if leftInt, leftOk := left.(int); leftOk {
		if rightInt, rightOk := right.(int); rightOk {
			return vm.performIntArithmetic(leftInt, rightInt, operator)
		}
		// int + float = float
		if rightFloat, rightOk := right.(float64); rightOk {
			return vm.performFloatArithmetic(float64(leftInt), rightFloat, operator)
		}
	}

	// Handle float + anything = float operations
	if leftFloat, leftOk := left.(float64); leftOk {
		if rightFloat, rightOk := right.(float64); rightOk {
			return vm.performFloatArithmetic(leftFloat, rightFloat, operator)
		}
		if rightInt, rightOk := right.(int); rightOk {
			return vm.performFloatArithmetic(leftFloat, float64(rightInt), operator)
		}
	}

	return nil, fmt.Errorf("unsupported operand types for arithmetic: %T and %T", left, right)
}

// performIntArithmetic handles arithmetic operations between two integers
func (vm *VirtualMachine) performIntArithmetic(left, right int, operator memory.Operator) (interface{}, error) {
	switch operator {
	case memory.ADD:
		result := left + right
		return result, nil
	case memory.SUB:
		result := left - right
		return result, nil
	case memory.MUL:
		result := left * right
		return result, nil
	case memory.DIV:
		if right == 0 {
			return nil, fmt.Errorf("division by zero")
		}
		// Check if division is exact to decide between int or float result
		if left%right == 0 {
			result := left / right
			return result, nil
		} else {
			err := "type error: cannot assign value of type float to variable of type int"
			fmt.Printf("Arithmetic operation failed: %v\n", err)
			os.Exit(1)
			return nil, nil
		}
	default:
		return nil, fmt.Errorf("unsupported operator for integers: %v", operator)
	}
}

// performFloatArithmetic handles arithmetic operations between floating-point numbers
func (vm *VirtualMachine) performFloatArithmetic(left, right float64, operator memory.Operator) (interface{}, error) {
	switch operator {
	case memory.ADD:
		result := left + right
		return result, nil
	case memory.SUB:
		result := left - right
		return result, nil
	case memory.MUL:
		result := left * right
		return result, nil
	case memory.DIV:
		if right == 0.0 {
			return nil, fmt.Errorf("division by zero")
		}
		result := left / right
		return result, nil
	default:
		return nil, fmt.Errorf("unsupported operator for floats: %v", operator)
	}
}

// executeComparison performs relational comparison operations (>, <, !=).
// Improved to handle type-safe comparisons
//
// Parameters:
//   - instruction: The comparison quadruple containing operands and result addresses
func (vm *VirtualMachine) executeComparison(instruction semantic.Quadruple) {
	leftValue, err := vm.Memory.Get(instruction.LeftOperand.(int))
	if err != nil {
		panic(fmt.Sprintf("Error getting left operand for comparison: %v", err))
	}

	rightValue, err := vm.Memory.Get(instruction.RightOperand.(int))
	if err != nil {
		panic(fmt.Sprintf("Error getting right operand for comparison: %v", err))
	}

	resultAddress := instruction.Result.(int)

	comparisonResult, err := vm.performComparison(leftValue, rightValue, instruction.Operator.(memory.Operator))
	if err != nil {
		panic(fmt.Sprintf("Comparison operation failed: %v", err))
	}

	if err := vm.Memory.Set(resultAddress, comparisonResult); err != nil {
		panic(fmt.Sprintf("Error setting comparison result: %v", err))
	}
}

// performComparison handles comparison operations with proper type handling
func (vm *VirtualMachine) performComparison(left, right interface{}, operator memory.Operator) (bool, error) {
	// Handle int-int comparisons
	if leftInt, leftOk := left.(int); leftOk {
		if rightInt, rightOk := right.(int); rightOk {
			return vm.compareInts(leftInt, rightInt, operator), nil
		}
		// int-float comparison: promote int to float
		if rightFloat, rightOk := right.(float64); rightOk {
			return vm.compareFloats(float64(leftInt), rightFloat, operator), nil
		}
	}

	// Handle float comparisons
	if leftFloat, leftOk := left.(float64); leftOk {
		if rightFloat, rightOk := right.(float64); rightOk {
			return vm.compareFloats(leftFloat, rightFloat, operator), nil
		}
		// float-int comparison: promote int to float
		if rightInt, rightOk := right.(int); rightOk {
			return vm.compareFloats(leftFloat, float64(rightInt), operator), nil
		}
	}

	return false, fmt.Errorf("unsupported operand types for comparison: %T and %T", left, right)
}

// compareInts performs comparison operations between integers
func (vm *VirtualMachine) compareInts(left, right int, operator memory.Operator) bool {
	switch operator {
	case memory.GREATER:
		return left > right
	case memory.LESS:
		return left < right
	case memory.NOTEQUAL:
		return left != right
	default:
		panic(fmt.Sprintf("Unhandled comparison operator: %v", operator))
	}
}

// compareFloats performs comparison operations between floating-point numbers
func (vm *VirtualMachine) compareFloats(left, right float64, operator memory.Operator) bool {
	switch operator {
	case memory.GREATER:
		return left > right
	case memory.LESS:
		return left < right
	case memory.NOTEQUAL:
		return left != right
	default:
		panic(fmt.Sprintf("Unhandled comparison operator: %v", operator))
	}
}

// executeConditionalJump performs conditional jump based on boolean condition.
//
// Parameters:
//   - instruction: The GOTOF quadruple containing condition and target addresses
//
// Returns:
//   - bool: true if condition is met (continue normal execution), false if jumped
func (vm *VirtualMachine) executeConditionalJump(instruction semantic.Quadruple) bool {
	conditionAddress := instruction.LeftOperand.(int)
	conditionValue, err := vm.Memory.Get(conditionAddress)
	if err != nil {
		panic(fmt.Sprintf("Error getting condition for jump: %v", err))
	}

	// Ensure we have a boolean value
	conditionBool, ok := conditionValue.(bool)
	if !ok {
		panic(fmt.Sprintf("Expected bool for conditional jump, got %T", conditionValue))
	}

	// GOTOF: Jump if condition is false
	if !conditionBool {
		targetAddress := instruction.Result.(int)
		vm.programCounter = targetAddress
		return false // Don't increment PC
	}

	return true // Continue normal execution
}

// executeAssignment performs variable assignment operation.
//
// Parameters:
//   - instruction: The ASSIGN quadruple containing source and destination addresses
func (vm *VirtualMachine) executeAssignment(instruction semantic.Quadruple) {
	sourceAddress := instruction.LeftOperand.(int)
	destinationAddress := instruction.Result.(int)

	sourceValue, err := vm.Memory.Get(sourceAddress)
	if err != nil {
		panic(fmt.Sprintf("Error getting source value for assignment: %v", err))
	}

	if err := vm.Memory.Set(destinationAddress, sourceValue); err != nil {
		panic(fmt.Sprintf("Error setting assignment value: %v", err))
	}
}

// executePrint outputs a value to the console.
//
// Parameters:
//   - instruction: The PRINT quadruple containing the address of value to print
func (vm *VirtualMachine) executePrint(instruction semantic.Quadruple) {
	valueAddress := instruction.Result.(int)
	valueToPrint, err := vm.Memory.Get(valueAddress)
	if err != nil {
		panic(fmt.Sprintf("Error getting value to print: %v", err))
	}

	fmt.Printf("Output: %v\n", valueToPrint)
}

// negateValue performs unary negation on numeric values while preserving type.
//
// Parameters:
//   - value: The value to negate (int, float64, etc.)
//
// Returns:
//   - interface{}: The negated value with the same type as input
//
// Panics:
//   - If the value is not a numeric type that supports negation
func (vm *VirtualMachine) negateValue(value interface{}) interface{} {
	switch typedValue := value.(type) {
	case int:
		return -typedValue
	case float64:
		return -typedValue
	default:
		panic(fmt.Sprintf("Cannot negate value of type %T", typedValue))
	}
}
