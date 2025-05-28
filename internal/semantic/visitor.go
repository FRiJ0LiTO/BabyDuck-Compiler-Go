package semantic

import (
	"BabyDuck/grammar/generated"
	"BabyDuck/internal/memory"
	"BabyDuck/internal/symbol"
	"BabyDuck/structures/stack"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"os"
	"strconv"
)

type Quadruple struct {
	Operator     any
	LeftOperand  any
	RightOperand any
	Result       any
	Scope        string
}

type Visitor struct {
	generated.BaseBabyDuckVisitor
	Directory           *symbol.FunctionDirectory
	semanticCube        Cube
	MemoryManager       *memory.Manager
	typesStack          *stack.Stack
	JumpsStack          *stack.Stack
	CurrentScope        *stack.Stack // Stack of active scopes, with the innermost scope at the end.
	Quadruples          []Quadruple
	TempVariableCounter int
	debug               []bool
}

func NewVisitor(directory *symbol.FunctionDirectory, debug ...bool) *Visitor {
	return &Visitor{
		Directory:           directory,
		semanticCube:        NewSemanticCube(),
		MemoryManager:       memory.NewMemoryManager(memory.DefaultMemoryConfig),
		typesStack:          stack.New(),
		JumpsStack:          stack.New(),
		CurrentScope:        stack.New(),
		Quadruples:          []Quadruple{},
		TempVariableCounter: 0,
		debug:               debug,
	}
}

// getTemporaryDataType returns the corresponding temporary data type for a given variable type.
// This is used to track resource allocation for temporary variables in functions.
//
// Parameters:
//   - varType: The original data type of the variable ("int", "float", or "bool")
//
// Returns:
//   - memory.DataType: The corresponding temporary type ("TempInt", "TempFloat", or "TempBool")
func getTemporaryDataType(varType memory.DataType) string {
	switch varType {
	case "int":
		return "Tempint"
	case "float":
		return "Tempfloat"
	default:
		return "Tempbool"
	}
}

// newTemporaryVariable generates a new unique temporary variable name
// Each call increments the counter to ensure uniqueness
// Returns: A string representing the new temporary variable (e.g., "t0", "t1", etc.)
func (v *Visitor) newTemporaryVariable() any {
	dataType := v.typesStack.Peek().(memory.DataType)
	virtualAddress, _ := v.MemoryManager.GetAddress("Temporal", dataType)

	// Create variable name with current counter value
	temporaryVariable := "t" + strconv.Itoa(v.TempVariableCounter)
	v.TempVariableCounter++
	if len(v.debug) > 0 && v.debug[0] {

		return temporaryVariable
	}
	scope := v.CurrentScope.Peek().(string)
	_ = v.Directory.AddVariable(temporaryVariable, dataType, virtualAddress, scope)

	tempVarType := getTemporaryDataType(dataType)
	v.Directory.AddResource(scope, tempVarType, 1)

	return virtualAddress
}

// generateQuadruple creates a new quadruple and adds it to the quadruples list
// A quadruple represents an intermediate code instruction with operator and operands
// Parameters:
//   - operator: The operation to perform (e.g., "+", "-", "=", "GOTO")
//   - leftOperand: The first operand of the operation
//   - rightOperand: The second operand of the operation (maybe empty)
//   - result: The destination where the result will be stored
func (v *Visitor) generateQuadruple(operator any, leftOperand any, rightOperand any, result any) {
	quadruple := Quadruple{
		Operator:     operator,
		LeftOperand:  leftOperand,
		RightOperand: rightOperand,
		Result:       result,
		Scope:        v.CurrentScope.Peek().(string),
	}

	v.Quadruples = append(v.Quadruples, quadruple)
}

// updateQuadruple modifies an existing quadruple at the specified index
// This is typically used to complete jump instructions by setting their destination
// Parameters:
//   - index: The position of the quadruple to update
//   - result: The new result value to set in the quadruple
func (v *Visitor) updateQuadruple(index int, result int) {
	quadruple := v.Quadruples[index]
	quadruple.Result = result
	v.Quadruples[index] = quadruple
}

func (v *Visitor) PrintQuadruplesTable() {
	fmt.Println("\n=== Quadruples Table ===")
	fmt.Println("Index | Operator | Operand 1 | Operand 2 |      Result     | Scope")
	fmt.Println("-------|----------|------------|------------|--------------------|----------")

	for i, quad := range v.Quadruples {
		formatOperand := func(operand any) string {
			switch val := operand.(type) {
			case int:
				return fmt.Sprintf("%d", val)
			case string:
				return val
			case nil:
				return "nil"
			default:
				return fmt.Sprintf("%v", val)
			}
		}

		opStr := formatOperand(quad.Operator)
		leftStr := formatOperand(quad.LeftOperand)
		rightStr := formatOperand(quad.RightOperand)
		resultStr := formatOperand(quad.Result)

		fmt.Printf("%6d | %8s | %10s | %10s | %18s | %s\n",
			i, opStr, leftStr, rightStr, resultStr, quad.Scope)
	}

	fmt.Println("=============================================================")
}

// Visit is the main entry point for traversing the parse tree
// It dispatches to the appropriate visitor method based on the node type
func (v *Visitor) Visit(tree antlr.ParseTree) any {
	switch node := tree.(type) {
	case *generated.ProgramContext:
		return v.VisitProgram(node)
	case *generated.ProgramHeaderContext:
		return v.VisitProgramHeader(node)
	case *generated.ProgramBodyContext:
		return v.VisitProgramBody(node)
	case *generated.FunctionDeclarationContext:
		return v.VisitFunctionDeclaration(node)
	case *generated.FunctionBodyContext:
		return v.VisitFunctionBody(node)
	case *generated.CodeBlockContext:
		return v.VisitCodeBlock(node)
	case *generated.StatementContext:
		return v.VisitStatement(node)
	case *generated.AssignmentContext:
		return v.VisitAssignment(node)
	case *generated.ConditionalContext:
		return v.VisitConditional(node)
	case *generated.LoopContext:
		return v.VisitLoop(node)
	case *generated.FunctionCallContext:
		return v.VisitFunctionCall(node)
	case *generated.PrintStatementContext:
		return v.VisitPrintStatement(node)
	case *generated.ExpressionContext:
		return v.VisitExpression(node)
	case *generated.RelationalOperatorContext:
		return v.VisitRelationalOperator(node)
	case *generated.ArithmeticExpressionContext:
		return v.VisitArithmeticExpression(node)
	case *generated.AdditiveOperatorContext:
		return v.VisitAdditiveOperator(node)
	case *generated.TermContext:
		return v.VisitTerm(node)
	case *generated.MultiplicativeOperatorContext:
		return v.VisitMultiplicativeOperator(node)
	case *generated.FactorContext:
		return v.VisitFactor(node)
	case *generated.ParenthesizedExpressionContext:
		return v.VisitParenthesizedExpression(node)
	case *generated.ValueWithOptionalSignContext:
		return v.VisitValueWithOptionalSign(node)
	case *generated.ValueContext:
		return v.VisitValue(node)
	default:
		fmt.Printf("Unimplemented node type: %T\n", node)
		return nil
	}
}

// VisitProgram processes the top-level program structure
// Generates PROGRAM and END quadruples and visits all program sections
func (v *Visitor) VisitProgram(ctx *generated.ProgramContext) any {
	// Push program scope to the stack
	v.CurrentScope.Push("program")
	v.JumpsStack.Push(len(v.Quadruples))

	// Generate program start quadruple
	virtualAddressOpPro := memory.IdentifyOperator("PROGRAM")
	v.generateQuadruple(virtualAddressOpPro, 0, 0, 0)

	// Process program header if present
	if ctx.ProgramHeader() != nil {
		v.Visit(ctx.ProgramHeader())
	}

	// Process all function declarations
	if ctx.AllFunctionsSection() != nil {
		for _, functionSection := range ctx.AllFunctionsSection() {
			v.Visit(functionSection.FunctionDeclaration())
		}
	}

	// Process the main program body
	if ctx.ProgramBody() != nil {
		v.Visit(ctx.ProgramBody())
	}

	// Generate program end quadruple
	virtualAddressOpEnd := memory.IdentifyOperator("END")
	v.generateQuadruple(virtualAddressOpEnd, 0, 0, 0)

	return nil
}

// VisitProgramBody processes the main code block of the program
func (v *Visitor) VisitProgramBody(ctx *generated.ProgramBodyContext) any {
	if ctx.MAIN() != nil {
		v.updateQuadruple(v.JumpsStack.Pop().(int), len(v.Quadruples))
		return v.Visit(ctx.CodeBlock())
	}
	return nil
}

// VisitFunctionDeclaration processes function declarations
// Creates a new scope and generates function-related quadruples
func (v *Visitor) VisitFunctionDeclaration(ctx *generated.FunctionDeclarationContext) any {
	// Get function name and push to scope stack
	functionName := ctx.Identifier().GetText()
	v.TempVariableCounter = 0
	v.CurrentScope.Push(functionName)
	v.Directory.FunctionsDirectory[functionName].Position = len(v.Quadruples)

	// Generate function declaration quadruple
	params := ctx.AllParameter()
	newParams := make([]memory.DataType, 0, len(ctx.AllCOMMA())+1)

	for _, context := range params {
		if context.Type_() != nil {
			variableType := memory.DataType(context.Type_().GetText())
			newParams = append(newParams, variableType)
			resourceType := "Local" + string(variableType)
			v.Directory.AddResource(functionName, resourceType, 1)
		}
	}

	v.Directory.FunctionsDirectory[functionName].Parameters = newParams

	// Process function body
	return v.Visit(ctx.FunctionBody())
}

// VisitFunctionBody processes the code block within a function
// Handles scope management and generates end function quadruple
func (v *Visitor) VisitFunctionBody(ctx *generated.FunctionBodyContext) any {
	// Process function code block
	v.Visit(ctx.CodeBlock())

	// Pop function scope from stack
	v.CurrentScope.Pop()

	// Generate end function quadruple
	virtualAddressOp := memory.IdentifyOperator("ENDFUNC")
	v.generateQuadruple(virtualAddressOp, 0, 0, 0)

	if len(v.debug) > 0 && v.debug[0] {
		v.TempVariableCounter = 0
	} else {
		v.MemoryManager.ResetLocal()
	}

	return nil
}

// VisitCodeBlock processes a block of statements
// Visits each statement in the block sequentially
func (v *Visitor) VisitCodeBlock(ctx *generated.CodeBlockContext) any {
	// Process all statements in the code block
	if ctx.AllStatement() != nil {
		for _, statement := range ctx.AllStatement() {
			v.Visit(statement)
		}
	}
	return nil
}

// VisitStatement processes individual statements
// Dispatches to the appropriate visitor method based on statement type
func (v *Visitor) VisitStatement(ctx *generated.StatementContext) any {
	if ctx.Assignment() != nil {
		return v.Visit(ctx.Assignment())
	} else if ctx.Conditional() != nil {
		return v.Visit(ctx.Conditional())
	} else if ctx.Loop() != nil {
		return v.Visit(ctx.Loop())
	} else if ctx.FunctionCall() != nil {
		return v.Visit(ctx.FunctionCall())
	} else if ctx.PrintStatement() != nil {
		return v.Visit(ctx.PrintStatement())
	}
	return nil
}

// VisitAssignment processes variable assignment statements
// Generates assignment quadruple
func (v *Visitor) VisitAssignment(ctx *generated.AssignmentContext) any {
	// Evaluate the expression on the right side
	expressionResult := v.Visit(ctx.Expression())

	// Get the identifier on the left side
	variableIdentifier := ctx.Identifier().GetText()
	var result any
	if len(v.debug) > 0 && v.debug[0] {
		result = variableIdentifier
	} else {
		scopes := v.CurrentScope.ToStringSlice()
		variable, _ := v.Directory.FindVariable(scopes, variableIdentifier)
		if variable.VariableType != v.typesStack.Peek() {
			fmt.Printf("Type error: cannot assign value of type '%s' to variable '%s' of type '%s'\n",
				v.typesStack.Peek(), variableIdentifier, variable.VariableType)
			os.Exit(1)
		}
		result = variable.VirtualDirection
	}

	// Generate assignment quadruple
	virtualAddressOp := memory.IdentifyOperator("=")
	v.generateQuadruple(virtualAddressOp, expressionResult, 0, result)
	// Pop the result type of the expression from the types stack since it has been consumed by the assignment
	v.typesStack.Pop()

	return nil
}

// VisitConditional processes if-else conditional structures
// Generates appropriate quadruples for conditional branching
func (v *Visitor) VisitConditional(ctx *generated.ConditionalContext) any {
	// Process if block
	if ctx.IfBlock() != nil {
		// Evaluate the condition expression
		conditionResult := v.Visit(ctx.IfBlock().ParenthesizedExpression())

		// Save position for GOTOF instruction that will be filled later
		jumpPosition := len(v.Quadruples)
		v.JumpsStack.Push(jumpPosition)
		virtualAddressOp := memory.IdentifyOperator("GOTOF")
		v.generateQuadruple(virtualAddressOp, conditionResult, 0, 0)

		// Process the code block within the if statement
		v.Visit(ctx.IfBlock().CodeBlock())
	}

	// Process else block if it exists
	if ctx.ElseBlock() != nil {
		// Get the position of if's GOTOF to update it
		falseJumpPosition := v.JumpsStack.Pop()

		// Add GOTO at the end of if block to skip else block
		v.JumpsStack.Push(len(v.Quadruples))
		virtualAddressOp := memory.IdentifyOperator("GOTO")
		v.generateQuadruple(virtualAddressOp, 0, 0, 0)

		// Update the GOTOF to jump to the beginning of the else block
		v.updateQuadruple(falseJumpPosition.(int), len(v.Quadruples))

		// Process the code block within the else statement
		v.Visit(ctx.ElseBlock().CodeBlock())
	}

	// Complete the conditional structure
	if ctx.SEMICOLON() != nil {
		// Get the last jump position (either from if or else) and update it
		lastJumpPosition := v.JumpsStack.Pop()
		v.updateQuadruple(lastJumpPosition.(int), len(v.Quadruples))
	}
	return nil
}

// VisitLoop processes loop structures (currently supports while loops)
// Handles the control flow for iterative execution
func (v *Visitor) VisitLoop(ctx *generated.LoopContext) any {
	if ctx.WHILE() != nil {
		// Save the position to return to for loop condition evaluation
		loopStartPosition := len(v.Quadruples)
		v.JumpsStack.Push(loopStartPosition)

		// Evaluate the loop condition
		conditionResult := v.Visit(ctx.ParenthesizedExpression())

		// Save position for GOTOF that will exit the loop when condition is false
		conditionalJumpPosition := len(v.Quadruples)
		v.JumpsStack.Push(conditionalJumpPosition)
		virtualAddressOpGotoF := memory.IdentifyOperator("GOTOF")
		v.generateQuadruple(virtualAddressOpGotoF, conditionResult, 0, 0)

		// Process the loop body
		v.Visit(ctx.CodeBlock())

		// Retrieve jump positions from stack
		exitJumpPosition := v.JumpsStack.Pop()
		returnToConditionPosition := v.JumpsStack.Pop()

		// Update the conditional jump to point to after the loop
		v.updateQuadruple(exitJumpPosition.(int), len(v.Quadruples)+1)

		// Add a jump back to the condition evaluation
		virtualAddressOpGoto := memory.IdentifyOperator("GOTO")
		v.generateQuadruple(virtualAddressOpGoto, 0, 0, returnToConditionPosition.(int))
	}
	return nil
}

// VisitFunctionCall processes function calls in the code
// Generates PARAM quadruples for arguments and a CALL quadruple
// generateQuadruple creates a new quadruple and adds it to the quadruples list
// A quadruple represents an intermediate code instruction with operator and operands
// Parameters:
func (v *Visitor) VisitFunctionCall(ctx *generated.FunctionCallContext) any {
	functionName := ctx.Identifier().GetText()
	virtualAddressOpEra := memory.IdentifyOperator("ERA")
	v.generateQuadruple(virtualAddressOpEra, functionName, 0, 0)

	expectedParams := len(v.Directory.FunctionsDirectory[functionName].Parameters)

	if expectedParams > 0 {
		if ctx.ArgumentList() == nil {
			fmt.Printf("Error calling function '%s': expected %d parameters but got %d\n",
				functionName, expectedParams, 0)
			os.Exit(1)
		}

		// Process all arguments
		var argumentValues []any
		for _, expression := range ctx.ArgumentList().AllExpression() {
			argumentValues = append(argumentValues, v.Visit(expression))
		}

		actualParams := len(argumentValues)
		if actualParams != expectedParams {
			fmt.Printf("Error calling function '%s': expected %d parameters but got %d\n",
				functionName, expectedParams, actualParams)
			os.Exit(1)
		}

		for i, value := range argumentValues {
			variableType, _ := v.Directory.LookupVariableByVirtualAddress(v.CurrentScope.ToStringSlice(), value.(int))
			parameterType := v.Directory.FunctionsDirectory[functionName].Parameters[i]
			if variableType != parameterType {
				fmt.Printf("Type mismatch in function '%s' call: parameter %d expected type '%s' but got '%s'\n",
					functionName, i+1, parameterType, variableType)
				os.Exit(1)
			}
		}

		// Generate parameter quadruples for each argument
		for _, argValue := range argumentValues {
			virtualAddressOpParam := memory.IdentifyOperator("PARAM")
			v.generateQuadruple(virtualAddressOpParam, argValue, 0, 0)
		}
	}

	// Generate the function call quadruple
	virtualAddressOpGosub := memory.IdentifyOperator("GOSUB")
	v.generateQuadruple(virtualAddressOpGosub, functionName, 0, 0)
	return nil
}

// VisitPrintStatement processes print statements
// Generates PRINT quadruples for each printable item
func (v *Visitor) VisitPrintStatement(ctx *generated.PrintStatementContext) any {
	for i, printable := range ctx.AllPrintable() {
		virtualAddressOpPrint := memory.IdentifyOperator("PRINT")
		if printable.Expression() != nil {
			// If it's an expression, evaluate it and print the result
			expressionResult := v.Visit(printable.Expression())
			v.generateQuadruple(virtualAddressOpPrint, 0, 0, expressionResult)
		} else {
			// If it's a literal, store it
			literalValue := ctx.Printable(i).GetText()
			virtualAddress, _ := v.Directory.LookupConstant(literalValue)
			v.generateQuadruple(virtualAddressOpPrint, 0, 0, virtualAddress)
		}
	}
	v.generateQuadruple(memory.IdentifyOperator("ALLPRINT"), 0, 0, 0)
	return nil
}

// VisitExpression processes expressions, including relational operations
// Returns a reference to the result (variable or temporary)
func (v *Visitor) VisitExpression(ctx *generated.ExpressionContext) any {
	// Process the left-hand arithmetic expression
	leftOperand := v.Visit(ctx.ArithmeticExpression())

	// Check if there's a relational component
	if ctx.RelationalExpression() != nil {
		// Get the relational operator
		relationalOperator := v.Visit(ctx.RelationalExpression().RelationalOperator())

		// Process the right-hand arithmetic expression
		rightOperand := v.Visit(ctx.RelationalExpression().ArithmeticExpression())

		rightOperandType := v.typesStack.Pop().(memory.DataType)
		leftOperandType := v.typesStack.Pop().(memory.DataType)
		resultType, ok := v.semanticCube.GetResultType(leftOperandType, rightOperandType, relationalOperator.(string))

		if !ok {
			fmt.Println("ERROR", resultType, rightOperandType, leftOperandType)
		}

		v.typesStack.Push(resultType)

		// Create a temporary variable to store the result
		resultTemp := v.newTemporaryVariable()
		virtualAddressOp := memory.IdentifyOperator(relationalOperator.(string))
		v.generateQuadruple(virtualAddressOp, leftOperand, rightOperand, resultTemp)
		return resultTemp
	}

	// If no relational component, return the arithmetic result
	return leftOperand
}

// VisitRelationalOperator identifies and returns the string representation
// of the relational operator
func (v *Visitor) VisitRelationalOperator(ctx *generated.RelationalOperatorContext) any {
	if ctx.GREATER() != nil {
		return ">"
	}
	if ctx.LESS() != nil {
		return "<"
	}
	if ctx.NOT_EQUAL() != nil {
		return "!="
	}
	return ""
}

// VisitArithmeticExpression processes arithmetic expressions (terms connected by + or -)
// Returns a reference to the result (variable or temporary)
func (v *Visitor) VisitArithmeticExpression(ctx *generated.ArithmeticExpressionContext) any {
	// Process the first term
	result := v.Visit(ctx.Term(0))

	// Process additional terms if they exist
	for i := range ctx.AllAdditiveOperator() {
		operator := v.Visit(ctx.AdditiveOperator(i))
		nextTerm := v.Visit(ctx.Term(i + 1))

		rightOperandType := v.typesStack.Pop().(memory.DataType)
		leftOperandType := v.typesStack.Pop().(memory.DataType)

		resultType, ok := v.semanticCube.GetResultType(leftOperandType, rightOperandType, operator.(string))
		if !ok {
			// TODO: Handle error
			fmt.Printf("Type mismatch error: cannot perform operation '%s' between types '%s' and '%s'\n",
				operator, leftOperandType, rightOperandType)
			os.Exit(1)
		}
		v.typesStack.Push(resultType)

		// Create a temporary variable for the result
		resultTemp := v.newTemporaryVariable()
		virtualAddressOpRelational := memory.IdentifyOperator(operator.(string))
		v.generateQuadruple(virtualAddressOpRelational, result, nextTerm, resultTemp)
		result = resultTemp
	}

	return result
}

// VisitAdditiveOperator identifies and returns the string representation
// of the additive operator (+ or -)
func (v *Visitor) VisitAdditiveOperator(ctx *generated.AdditiveOperatorContext) any {
	if ctx.OP_ADD() != nil {
		return "+"
	}
	if ctx.OP_SUBTRACT() != nil {
		return "-"
	}
	return ""
}

// VisitTerm processes terms (factors connected by * or /)
// Returns a reference to the result (variable or temporary)
func (v *Visitor) VisitTerm(ctx *generated.TermContext) any {
	// Process the first factor
	result := v.Visit(ctx.Factor(0))

	// Process additional factors if they exist
	for i := range ctx.AllMultiplicativeOperator() {
		operator := v.Visit(ctx.MultiplicativeOperator(i))
		nextFactor := v.Visit(ctx.Factor(i + 1))

		rightOperandType := v.typesStack.Pop().(memory.DataType)
		leftOperandType := v.typesStack.Pop().(memory.DataType)

		resultType, ok := v.semanticCube.GetResultType(leftOperandType, rightOperandType, operator.(string))
		if !ok {
			// TODO: Handle error
			fmt.Printf("Type mismatch error: cannot perform operation '%s' between types '%s' and '%s'\n",
				operator, leftOperandType, rightOperandType)
			os.Exit(1)
		}
		v.typesStack.Push(resultType)

		// Create a temporary variable for the result
		resultTemp := v.newTemporaryVariable()
		virtualAddressOpMultiplicative := memory.IdentifyOperator(operator.(string))
		v.generateQuadruple(virtualAddressOpMultiplicative, result, nextFactor, resultTemp)
		result = resultTemp
	}

	return result
}

// VisitMultiplicativeOperator identifies and returns the string representation
// of the multiplicative operator (* or /)
func (v *Visitor) VisitMultiplicativeOperator(ctx *generated.MultiplicativeOperatorContext) any {
	if ctx.OP_MULTIPLY() != nil {
		return "*"
	}
	if ctx.OP_DIVIDE() != nil {
		return "/"
	}
	return ""
}

// VisitFactor processes a factor (parenthesized expression or value)
// Returns a reference to the result
func (v *Visitor) VisitFactor(ctx *generated.FactorContext) any {
	if ctx.ParenthesizedExpression() != nil {
		return v.Visit(ctx.ParenthesizedExpression())
	}
	return v.Visit(ctx.ValueWithOptionalSign())
}

// VisitParenthesizedExpression processes expressions within parentheses
// Returns a reference to the result
func (v *Visitor) VisitParenthesizedExpression(ctx *generated.ParenthesizedExpressionContext) any {
	return v.Visit(ctx.Expression())
}

// VisitValueWithOptionalSign processes values that may have a leading sign
// Returns a reference to the result (variable or temporary)
func (v *Visitor) VisitValueWithOptionalSign(ctx *generated.ValueWithOptionalSignContext) any {
	// Process the base value
	valueResult := v.Visit(ctx.Value())

	// Check if there's a negative sign
	if ctx.AdditiveOperator() != nil && v.Visit(ctx.AdditiveOperator()) == ("-") {
		// Create a temporary for the negated value
		resultTemp := v.newTemporaryVariable()
		virtualAddressOp := memory.IdentifyOperator("NEG")
		v.generateQuadruple(virtualAddressOp, valueResult, 0, resultTemp)
		return resultTemp
	}

	return valueResult
}

// VisitValue processes basic values (identifiers or constants)
// Returns the identifier or constant value as a string
func (v *Visitor) VisitValue(ctx *generated.ValueContext) any {
	if ctx.Identifier() != nil {
		variableName := ctx.Identifier().GetText()
		scopes := v.CurrentScope.ToStringSlice()
		variableInfo, _ := v.Directory.FindVariable(scopes, variableName)
		v.typesStack.Push(variableInfo.VariableType)
		if len(v.debug) > 0 && v.debug[0] {
			return variableName
		}
		return variableInfo.VirtualDirection
	} else {
		constant := ctx.Constant().GetText()
		virtualAddress, constantType := v.Directory.LookupConstant(constant)
		v.typesStack.Push(constantType)
		if len(v.debug) > 0 && v.debug[0] {
			return constant
		}
		return virtualAddress
	}
}
