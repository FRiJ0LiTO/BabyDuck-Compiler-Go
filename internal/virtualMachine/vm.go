package virtualMachine

import (
	"BabyDuck/internal/memory"
	"BabyDuck/internal/semantic"
	"BabyDuck/internal/symbol"
	"BabyDuck/structures/queue"
	"BabyDuck/structures/stack"
	"fmt"
	"strconv"
	"strings"
)

// ExecutionContext represents the execution state for a function call.
// It maintains the context needed to properly handle function calls and returns,
// including the return address and memory state management.
type ExecutionContext struct {
	FunctionName   string  // Name of the function being called
	ReturnAddress  int     // Address to return to after function execution
	PreviousMemory *Memory // Memory that was active before the function call
	ParameterCount int     // Number of parameters expected by the function
}

// ParameterInfo simplifies parameter handling by encapsulating
// both the value and its original memory address.
type ParameterInfo struct {
	Value   any // The actual parameter value
	Address int // Original memory address of the parameter
}

// VirtualMachine represents the execution environment for BabyDuck programs.
// It manages memory, execution flow, and provides the runtime environment
// for executing intermediate code quadruples.
type VirtualMachine struct {
	GlobalMemory       *Memory                  // Global memory management system
	CurrentMemory      *Memory                  // Current active memory (global or local)
	FunctionsDirectory symbol.FunctionDirectory // Symbol table for variables and functions
	quadruples         []semantic.Quadruple     // List of intermediate code instructions
	executionStack     *stack.Stack             // Stack for function calls and scope management
	parameterQueue     *queue.Queue             // Queue for function parameters
	printable          []string
	programCounter     int // Current instruction pointer
}

// NewVirtualMachine creates and initializes a new virtual machine instance.
// It sets up the global memory based on the program's resource requirements
// and initializes all necessary components for program execution.
//
// Parameters:
//   - functionsDirectory: Symbol table containing function and variable information
//   - quadruples: Array of intermediate code instructions to execute
//
// Returns:
//   - *VirtualMachine: Initialized virtual machine ready for execution
func NewVirtualMachine(functionsDirectory symbol.FunctionDirectory, quadruples []semantic.Quadruple) *VirtualMachine {
	globalMemory := NewMemory(functionsDirectory.FunctionsDirectory["program"].Resources,
		memory.DefaultMemoryConfig, functionsDirectory.Constants)

	vm := &VirtualMachine{
		GlobalMemory:       globalMemory,
		CurrentMemory:      globalMemory, // Start with global memory as current
		FunctionsDirectory: functionsDirectory,
		quadruples:         quadruples,
		executionStack:     stack.New(),
		parameterQueue:     queue.New(),
		printable:          make([]string, 0),
		programCounter:     0,
	}

	return vm
}

// Execute runs the virtual machine, processing quadruples until completion.
// This is the main execution loop that dispatches instructions to their
// respective handlers and manages the program counter.
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

		case memory.ALLPRINT:
			vm.flushPrintBuffer()

		case memory.GOTOF:
			shouldIncrementPC = vm.executeConditionalJump(currentInstruction)

		case memory.GOTO:
			shouldIncrementPC = vm.executeUnconditionalJump(currentInstruction)

		case memory.ERA:
			vm.executeERA(currentInstruction)

		case memory.PARAM:
			vm.executeParam(currentInstruction)

		case memory.GOSUB:
			shouldIncrementPC = vm.executeGosub(currentInstruction)

		case memory.ENDFUNC:
			shouldIncrementPC = vm.executeEndFunc()

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
// This instruction marks the beginning of program execution and sets the
// program counter to the main function's starting address.
//
// Parameters:
//   - instruction: The PROGRAM quadruple containing the main function address
//
// Returns:
//   - bool: false to prevent automatic program counter increment
func (vm *VirtualMachine) handleProgramStart(instruction semantic.Quadruple) bool {
	mainFunctionAddress := instruction.Result.(int)
	vm.programCounter = mainFunctionAddress
	return false
}

// handleProgramEnd displays completion message and program statistics.
// This method is called when the END instruction is encountered,
// indicating successful program termination.
func (vm *VirtualMachine) handleProgramEnd() {
	fmt.Println("✓ BabyDuck program executed successfully.")
	fmt.Printf("  Instructions processed: %d\n", vm.programCounter+1)
}

// executeUnconditionalJump performs an unconditional jump to the specified address.
// This implements the GOTO instruction, which unconditionally transfers
// control to the target address.
//
// Parameters:
//   - instruction: The GOTO quadruple containing the target address
//
// Returns:
//   - bool: false to prevent automatic program counter increment
func (vm *VirtualMachine) executeUnconditionalJump(instruction semantic.Quadruple) bool {
	targetAddress := instruction.Result.(int)
	vm.programCounter = targetAddress
	return false
}

// executeConditionalJump performs conditional jump based on boolean condition.
// This implements the GOTOF (Go To if False) instruction, which jumps to
// the target address only if the condition evaluates to false.
//
// Parameters:
//   - instruction: The GOTOF quadruple containing condition and target addresses
//
// Returns:
//   - bool: false if jump occurs, true if execution continues normally
func (vm *VirtualMachine) executeConditionalJump(instruction semantic.Quadruple) bool {
	conditionAddress := instruction.LeftOperand.(int)
	conditionValue, err := vm.getCurrentMemoryValue(conditionAddress)
	if err != nil {
		panic(fmt.Sprintf("Error getting condition for jump: %v", err))
	}

	conditionBool, ok := conditionValue.(bool)
	if !ok {
		panic(fmt.Sprintf("Expected bool for conditional jump, got %T", conditionValue))
	}

	if !conditionBool {
		targetAddress := instruction.Result.(int)
		vm.programCounter = targetAddress
		return false
	}

	return true
}

// executeERA handles the ERA (Activation Record) instruction for function calls.
// ERA prepares the execution context for a function call by creating
// the necessary context information and pushing it onto the execution stack.
//
// Parameters:
//   - instruction: The ERA quadruple containing the function name
func (vm *VirtualMachine) executeERA(instruction semantic.Quadruple) {
	functionName := instruction.LeftOperand.(string)

	functionInfo := vm.FunctionsDirectory.FunctionsDirectory[functionName]

	context := &ExecutionContext{
		FunctionName:   functionName,
		ReturnAddress:  0,
		PreviousMemory: nil, // Will be set in GOSUB
		ParameterCount: len(functionInfo.Parameters),
	}

	vm.executionStack.Push(context)
}

// executeParam handles parameter passing for function calls.
// This instruction collects parameters from the calling context
// and pushes them onto the parameter stack for later use.
//
// Parameters:
//   - instruction: The PARAM quadruple containing the parameter address
func (vm *VirtualMachine) executeParam(instruction semantic.Quadruple) {
	parameterAddress := instruction.LeftOperand.(int)

	parameterValue, err := vm.CurrentMemory.Get(parameterAddress)
	if err != nil {
		panic(fmt.Sprintf("Error getting parameter value: %v", err))
	}

	vm.parameterQueue.Enqueue(parameterValue)
}

// executeGosub handles function calls (GOSUB instruction).
// This instruction performs the actual function call by setting up
// local memory, copying parameters, and transferring control to the function.
//
// Parameters:
//   - instruction: The GOSUB quadruple containing the function name
//
// Returns:
//   - bool: false to prevent automatic program counter increment
func (vm *VirtualMachine) executeGosub(instruction semantic.Quadruple) bool {
	functionName := instruction.LeftOperand.(string)

	if vm.executionStack.Height == 0 {
		panic("No execution context available for GOSUB")
	}

	context := vm.executionStack.Peek().(*ExecutionContext)

	if context.FunctionName != functionName {
		panic(fmt.Sprintf("Function name mismatch: expected %s, got %s", context.FunctionName, functionName))
	}

	functionInfo := vm.FunctionsDirectory.FunctionsDirectory[functionName]

	localMemory := NewMemory(functionInfo.Resources, memory.DefaultMemoryConfig, vm.FunctionsDirectory.Constants)

	// Save current memory (to restore upon return)
	context.PreviousMemory = vm.CurrentMemory
	context.ReturnAddress = vm.programCounter

	vm.CurrentMemory = localMemory

	vm.copyParametersToLocalMemory(functionInfo, context)

	vm.programCounter = functionInfo.Position

	return false // Don't increment PC since we're jumping
}

// copyParametersToLocalMemory copies parameters from the parameter stack to local variable memory.
// This method handles the parameter passing mechanism by extracting parameters
// from the stack and placing them in the appropriate local memory locations.
//
// Parameters:
//   - functionInfo: Information about the function including parameter types
//   - context: Execution context containing parameter count information
func (vm *VirtualMachine) copyParametersToLocalMemory(functionInfo *symbol.FunctionInformation, context *ExecutionContext) {
	if context.ParameterCount == 0 {
		return
	}

	parameters := make([]any, context.ParameterCount)

	currentNode := vm.parameterQueue.Start
	paramIndex := 0
	for currentNode != nil {
		parameters[paramIndex] = vm.parameterQueue.Dequeue()
		currentNode = currentNode.Next
		paramIndex++
	}

	var intCounter = 0
	var floatCounter = 0

	// Copy each parameter to its corresponding address in local memory
	for i, parameterType := range functionInfo.Parameters {
		if i >= len(parameters) {
			panic(fmt.Sprintf("Parameter type mismatch: expected %d parameters, got %d", len(functionInfo.Parameters), len(parameters)))
		}

		paramValue := parameters[i]
		var targetAddress int

		switch parameterType {
		case "int":
			targetAddress = vm.GlobalMemory.memoryConfig.LocalIntStart + intCounter
			intCounter++
		case "float":
			targetAddress = vm.GlobalMemory.memoryConfig.LocalFloatStart + floatCounter
			floatCounter++
		default:
			panic(fmt.Sprintf("Unsupported parameter type: %s", parameterType))
		}

		err := vm.CurrentMemory.Set(targetAddress, paramValue)
		if err != nil {
			panic(fmt.Sprintf("Error setting parameter %d in local memory: %v", i, err))
		}
	}
	vm.parameterQueue = queue.New()
}

// executeEndFunc handles function return (ENDFUNC instruction).
// This instruction manages the return from a function call by restoring
// the previous memory context and returning control to the caller.
//
// Returns:
//   - bool: false to prevent automatic program counter increment
func (vm *VirtualMachine) executeEndFunc() bool {
	if vm.executionStack.Height == 0 {
		panic("No execution context to return from")
	}

	context := vm.executionStack.Pop().(*ExecutionContext)

	vm.CurrentMemory = context.PreviousMemory

	vm.programCounter = context.ReturnAddress

	return true
}

// executeNegation performs unary negation operation on a value.
// This instruction implements the unary minus operator,
// negating numeric values (both integers and floats).
//
// Parameters:
//   - instruction: The NEG quadruple containing operand and result addresses
func (vm *VirtualMachine) executeNegation(instruction semantic.Quadruple) {
	operandAddress := instruction.LeftOperand.(int)
	resultAddress := instruction.Result.(int)

	operandValue, err := vm.getCurrentMemoryValue(operandAddress)
	if err != nil {
		panic(fmt.Sprintf("Error getting operand for negation: %v", err))
	}

	negatedResult := vm.negateValue(operandValue)

	if err := vm.setCurrentMemoryValue(resultAddress, negatedResult); err != nil {
		panic(fmt.Sprintf("Error setting negation result: %v", err))
	}
}

// executeArithmetic performs binary arithmetic operations (ADD, SUB, MUL, DIV).
// This method handles all binary arithmetic operations, supporting both
// integer and floating-point arithmetic with appropriate type coercion.
//
// Parameters:
//   - instruction: The arithmetic quadruple containing operands and result addresses
func (vm *VirtualMachine) executeArithmetic(instruction semantic.Quadruple) {
	leftOperandAddress := instruction.LeftOperand.(int)
	rightOperandAddress := instruction.RightOperand.(int)
	resultAddress := instruction.Result.(int)

	leftValue, err := vm.getCurrentMemoryValue(leftOperandAddress)
	if err != nil {
		panic(fmt.Sprintf("Error getting left operand: %v", err))
	}

	rightValue, err := vm.getCurrentMemoryValue(rightOperandAddress)
	if err != nil {
		panic(fmt.Sprintf("Error getting right operand: %v", err))
	}

	operator := instruction.Operator.(memory.Operator)

	result, err := vm.performArithmetic(leftValue, rightValue, operator)
	if err != nil {
		panic(fmt.Sprintf("Arithmetic operation failed: %v", err))
	}

	if err := vm.setCurrentMemoryValue(resultAddress, result); err != nil {
		panic(fmt.Sprintf("Error setting arithmetic result: %v", err))
	}
}

// executeComparison performs relational comparison operations (>, <, !=).
// This method handles all relational comparison operations, supporting
// comparisons between numeric types with appropriate type coercion.
//
// Parameters:
//   - instruction: The comparison quadruple containing operands and result addresses
func (vm *VirtualMachine) executeComparison(instruction semantic.Quadruple) {
	leftValue, err := vm.getCurrentMemoryValue(instruction.LeftOperand.(int))
	if err != nil {
		panic(fmt.Sprintf("Error getting left operand for comparison: %v", err))
	}

	rightValue, err := vm.getCurrentMemoryValue(instruction.RightOperand.(int))
	if err != nil {
		panic(fmt.Sprintf("Error getting right operand for comparison: %v", err))
	}

	resultAddress := instruction.Result.(int)

	comparisonResult, err := vm.performComparison(leftValue, rightValue, instruction.Operator.(memory.Operator))
	if err != nil {
		panic(fmt.Sprintf("Comparison operation failed: %v", err))
	}

	if err := vm.setCurrentMemoryValue(resultAddress, comparisonResult); err != nil {
		panic(fmt.Sprintf("Error setting comparison result: %v", err))
	}
}

// executeAssignment performs variable assignment operation.
// This instruction copies a value from one memory location to another,
// implementing the assignment operator in the BabyDuck language.
//
// Parameters:
//   - instruction: The ASSIGN quadruple containing source and destination addresses
func (vm *VirtualMachine) executeAssignment(instruction semantic.Quadruple) {
	sourceAddress := instruction.LeftOperand.(int)
	destinationAddress := instruction.Result.(int)

	sourceValue, err := vm.getCurrentMemoryValue(sourceAddress)
	if err != nil {
		panic(fmt.Sprintf("Error getting source value for assignment: %v", err))
	}

	if err := vm.setCurrentMemoryValue(destinationAddress, sourceValue); err != nil {
		panic(fmt.Sprintf("Error setting assignment value: %v", err))
	}
}

// executePrint outputs a value to the console.
// This instruction implements the print statement in BabyDuck,
// displaying the value at the specified memory address.
//
// Parameters:
//   - instruction: The PRINT quadruple containing the value address to print
func (vm *VirtualMachine) executePrint(instruction semantic.Quadruple) {
	valueAddress := instruction.Result.(int)
	valueToPrint, err := vm.getCurrentMemoryValue(valueAddress)
	if err != nil {
		panic(fmt.Sprintf("Error getting value to print: %v", err))
	}
	switch valueToPrint.(type) {
	case int:
		vm.printable = append(vm.printable, strconv.Itoa(valueToPrint.(int)))
	case float64:
		vm.printable = append(vm.printable, strconv.FormatFloat(valueToPrint.(float64), 'f', -1, 64))
	case bool:
		vm.printable = append(vm.printable, strconv.FormatBool(valueToPrint.(bool)))
	default:
		vm.printable = append(vm.printable, strings.ReplaceAll(valueToPrint.(string), `"`, ""))
	}
}

// flushPrintBuffer outputs all accumulated printable values and clears the buffer.
// This method joins all pending print values with spaces and outputs them as a single line,
// then resets the buffer for future print operations.
func (vm *VirtualMachine) flushPrintBuffer() {
	fmt.Println(strings.Join(vm.printable, " "))
	vm.printable = make([]string, 0)
}

// getCurrentMemoryValue gets a value from the appropriate memory (global or local).
// This method abstracts memory access by determining whether to access
// global or current (local) memory based on the address.
//
// Parameters:
//   - address: Memory address to read from
//
// Returns:
//   - any: Value stored at the address
//   - error: Error if memory access fails
func (vm *VirtualMachine) getCurrentMemoryValue(address int) (any, error) {
	// Check if it's a global address first
	if vm.isGlobalAddress(address) {
		return vm.GlobalMemory.Get(address)
	}
	// Otherwise use current memory (which might be local during function execution)
	return vm.CurrentMemory.Get(address)
}

// setCurrentMemoryValue sets a value in the appropriate memory (global or local).
// This method abstracts memory writes by determining whether to write to
// global or current (local) memory based on the address.
//
// Parameters:
//   - address: Memory address to write to
//   - value: Value to store at the address
//
// Returns:
//   - error: Error if memory write fails
func (vm *VirtualMachine) setCurrentMemoryValue(address int, value any) error {
	// Check if it's a global address first
	if vm.isGlobalAddress(address) {
		return vm.GlobalMemory.Set(address, value)
	}
	// Otherwise use current memory (which might be local during function execution)
	return vm.CurrentMemory.Set(address, value)
}

// isGlobalAddress determines if an address belongs to global memory.
// This method checks if the given address falls within any of the
// global memory ranges (variables or constants).
//
// Parameters:
//   - address: Memory address to check
//
// Returns:
//   - bool: true if address is in global memory, false otherwise
func (vm *VirtualMachine) isGlobalAddress(address int) bool {
	config := memory.DefaultMemoryConfig
	return (address >= config.GlobalIntStart && address <= config.GlobalIntEnd) ||
		(address >= config.GlobalFloatStart && address <= config.GlobalFloatEnd) ||
		(address >= config.GlobalBoolStart && address <= config.GlobalBoolEnd) ||
		(address >= config.ConstIntStart && address <= config.ConstStringEnd)
}

// performArithmetic dispatches arithmetic operations based on operand types.
// This method handles type checking and coercion for arithmetic operations,
// supporting both integer and floating-point arithmetic.
//
// Parameters:
//   - left: Left operand value
//   - right: Right operand value
//   - operator: Arithmetic operator to apply
//
// Returns:
//   - any: Result of the arithmetic operation
//   - error: Error if operation fails or types are incompatible
func (vm *VirtualMachine) performArithmetic(left, right any, operator memory.Operator) (any, error) {
	if leftInt, leftOk := left.(int); leftOk {
		if rightInt, rightOk := right.(int); rightOk {
			return vm.performIntArithmetic(leftInt, rightInt, operator)
		}
		if rightFloat, rightOk := right.(float64); rightOk {
			return vm.performFloatArithmetic(float64(leftInt), rightFloat, operator)
		}
	}

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

// performIntArithmetic performs arithmetic operations on integer operands.
// This method handles all arithmetic operations for integer values,
// with special handling for division to return appropriate types.
//
// Parameters:
//   - left: Left integer operand
//   - right: Right integer operand
//   - operator: Arithmetic operator to apply
//
// Returns:
//   - any: Result (int for exact division, float64 for non-exact)
//   - error: Error if operation fails (e.g., division by zero)
func (vm *VirtualMachine) performIntArithmetic(left, right int, operator memory.Operator) (any, error) {
	switch operator {
	case memory.ADD:
		return left + right, nil
	case memory.SUB:
		return left - right, nil
	case memory.MUL:
		return left * right, nil
	case memory.DIV:
		if right == 0 {
			return nil, fmt.Errorf("division by zero")
		}
		return left / right, nil
	default:
		return nil, fmt.Errorf("unsupported operator for integers: %v", operator)
	}
}

// performFloatArithmetic performs arithmetic operations on floating-point operands.
// This method handles all arithmetic operations for floating-point values.
//
// Parameters:
//   - left: Left floating-point operand
//   - right: Right floating-point operand
//   - operator: Arithmetic operator to apply
//
// Returns:
//   - any: Result of the floating-point operation
//   - error: Error if operation fails (e.g., division by zero)
func (vm *VirtualMachine) performFloatArithmetic(left, right float64, operator memory.Operator) (any, error) {
	switch operator {
	case memory.ADD:
		return left + right, nil
	case memory.SUB:
		return left - right, nil
	case memory.MUL:
		return left * right, nil
	case memory.DIV:
		if right == 0.0 {
			return nil, fmt.Errorf("division by zero")
		}
		return left / right, nil
	default:
		return nil, fmt.Errorf("unsupported operator for floats: %v", operator)
	}
}

// performComparison dispatches comparison operations based on operand types.
// This method handles type checking and coercion for comparison operations,
// supporting comparisons between numeric types.
//
// Parameters:
//   - left: Left operand value
//   - right: Right operand value
//   - operator: Comparison operator to apply
//
// Returns:
//   - bool: Result of the comparison operation
//   - error: Error if types are incompatible for comparison
func (vm *VirtualMachine) performComparison(left, right any, operator memory.Operator) (bool, error) {
	if leftInt, leftOk := left.(int); leftOk {
		if rightInt, rightOk := right.(int); rightOk {
			return vm.compareInts(leftInt, rightInt, operator), nil
		}
		if rightFloat, rightOk := right.(float64); rightOk {
			return vm.compareFloats(float64(leftInt), rightFloat, operator), nil
		}
	}

	if leftFloat, leftOk := left.(float64); leftOk {
		if rightFloat, rightOk := right.(float64); rightOk {
			return vm.compareFloats(leftFloat, rightFloat, operator), nil
		}
		if rightInt, rightOk := right.(int); rightOk {
			return vm.compareFloats(leftFloat, float64(rightInt), operator), nil
		}
	}

	return false, fmt.Errorf("unsupported operand types for comparison: %T and %T", left, right)
}

// compareInts performs comparison operations on integer operands.
// This method handles all supported comparison operations for integer values.
//
// Parameters:
//   - left: Left integer operand
//   - right: Right integer operand
//   - operator: Comparison operator to apply
//
// Returns:
//   - bool: Result of the integer comparison
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

// compareFloats performs comparison operations on floating-point operands.
// This method handles all supported comparison operations for floating-point values.
//
// Parameters:
//   - left: Left floating-point operand
//   - right: Right floating-point operand
//   - operator: Comparison operator to apply
//
// Returns:
//   - bool: Result of the floating-point comparison
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

// negateValue performs unary negation on numeric values.
// This utility function handles negation for both integer and floating-point values.
//
// Parameters:
//   - value: The value to negate (must be int or float64)
//
// Returns:
//   - any: The negated value with the same type as input
func (vm *VirtualMachine) negateValue(value any) any {
	switch typedValue := value.(type) {
	case int:
		return -typedValue
	case float64:
		return -typedValue
	default:
		panic(fmt.Sprintf("Cannot negate value of type %T", typedValue))
	}
}
