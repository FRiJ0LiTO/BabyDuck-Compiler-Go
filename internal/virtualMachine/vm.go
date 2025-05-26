package virtualMachine

import (
	"BabyDuck/internal/memory"
	"BabyDuck/internal/semantic"
	"BabyDuck/internal/symbol"
	"BabyDuck/structures/stack"
	"fmt"
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
		Memory:             NewMemory(memory.DefaultMemoryConfig, functionsDirectory.Constants),
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

	operandValue, _ := vm.Memory.Get(operandAddress)
	negatedResult := vm.negateValue(operandValue)

	_ = vm.Memory.Set(resultAddress, negatedResult)
}

// executeArithmetic performs binary arithmetic operations (ADD, SUB, MUL, DIV).
//
// Parameters:
//   - instruction: The arithmetic quadruple containing operands and result addresses
func (vm *VirtualMachine) executeArithmetic(instruction semantic.Quadruple) {
	leftOperandAddress := instruction.LeftOperand.(int)
	rightOperandAddress := instruction.RightOperand.(int)
	resultAddress := instruction.Result.(int)

	leftValue, _ := vm.Memory.Get(leftOperandAddress)
	rightValue, _ := vm.Memory.Get(rightOperandAddress)

	leftAsFloat := vm.toFloat64(leftValue)
	rightAsFloat := vm.toFloat64(rightValue)

	var calculationResult float64
	switch instruction.Operator.(memory.Operator) {
	case memory.ADD:
		calculationResult = leftAsFloat + rightAsFloat
	case memory.SUB:
		calculationResult = leftAsFloat - rightAsFloat
	case memory.MUL:
		calculationResult = leftAsFloat * rightAsFloat
	case memory.DIV:
		calculationResult = leftAsFloat / rightAsFloat
	default:
		panic(fmt.Sprintf("Unhandled arithmetic operator: %v", instruction.Operator))
	}

	_ = vm.Memory.Set(resultAddress, calculationResult)
}

// executeComparison performs relational comparison operations (>, <, !=).
//
// Parameters:
//   - instruction: The comparison quadruple containing operands and result addresses
func (vm *VirtualMachine) executeComparison(instruction semantic.Quadruple) {
	leftValue, _ := vm.Memory.Get(instruction.LeftOperand.(int))
	rightValue, _ := vm.Memory.Get(instruction.RightOperand.(int))
	resultAddress := instruction.Result.(int)

	leftAsFloat := vm.toFloat64(leftValue)
	rightAsFloat := vm.toFloat64(rightValue)

	var comparisonResult bool
	switch instruction.Operator.(memory.Operator) {
	case memory.GREATER:
		comparisonResult = leftAsFloat > rightAsFloat
	case memory.LESS:
		comparisonResult = leftAsFloat < rightAsFloat
	case memory.NOTEQUAL:
		comparisonResult = leftAsFloat != rightAsFloat
	default:
		panic(fmt.Sprintf("Unhandled comparison operator: %v", instruction.Operator))
	}

	_ = vm.Memory.Set(resultAddress, comparisonResult)
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
	conditionValue, _ := vm.Memory.Get(conditionAddress)

	// GOTOF: Jump if condition is false
	if conditionValue != true {
		targetAddress := instruction.Result.(int)
		vm.programCounter = targetAddress
	}

	return conditionValue.(bool)
}

// executeAssignment performs variable assignment operation.
//
// Parameters:
//   - instruction: The ASSIGN quadruple containing source and destination addresses
func (vm *VirtualMachine) executeAssignment(instruction semantic.Quadruple) {
	sourceAddress := instruction.LeftOperand.(int)
	destinationAddress := instruction.Result.(int)

	sourceValue, _ := vm.Memory.Get(sourceAddress)

	_ = vm.Memory.Set(destinationAddress, sourceValue)
}

// executePrint outputs a value to the console.
//
// Parameters:
//   - instruction: The PRINT quadruple containing the address of value to print
func (vm *VirtualMachine) executePrint(instruction semantic.Quadruple) {
	valueAddress := instruction.Result.(int)
	valueToPrint, _ := vm.Memory.Get(valueAddress)
	fmt.Println(valueToPrint)
}

// toFloat64 converts various numeric types to float64 for uniform arithmetic operations.
//
// Parameters:
//   - value: The value to convert (int, float64, etc.)
//
// Returns:
//   - float64: The converted floating-point value
//
// Panics:
//   - If the value cannot be converted to a numeric type
func (vm *VirtualMachine) toFloat64(value interface{}) float64 {
	switch typedValue := value.(type) {
	case int:
		return float64(typedValue)
	case float64:
		return typedValue
	default:
		panic(fmt.Sprintf("Cannot convert type %T to numeric value", typedValue))
	}
}

// negateValue performs unary negation on numeric values.
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
