package semantic

import (
	"BabyDuck/grammar/generated"
	"BabyDuck/internal/symbol"
	"BabyDuck/structures/stack"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
)

type Quadruple struct {
	Operator     interface{}
	LeftOperand  interface{}
	RightOperand interface{}
	Result       string
	Scope        string
}

type Visitor struct {
	generated.BaseBabyDuckVisitor
	Directory           *symbol.FunctionDirectory
	JumpsStack          *stack.Stack
	CurrentScope        *stack.Stack // Stack of active scopes, with the innermost scope at the end.
	Quadruples          []Quadruple
	TempVariableCounter int
}

func NewVisitor(directory *symbol.FunctionDirectory) *Visitor {
	return &Visitor{
		Directory:           directory,
		JumpsStack:          stack.New(),
		CurrentScope:        stack.New(),
		Quadruples:          []Quadruple{},
		TempVariableCounter: 0,
	}
}

// newTemporaryVariable generates a new unique temporary variable name
// Each call increments the counter to ensure uniqueness
// Returns: A string representing the new temporary variable (e.g., "t0", "t1", etc.)
func (v *Visitor) newTemporaryVariable() string {
	// Create variable name with current counter value
	temporaryVariable := "t" + strconv.Itoa(v.TempVariableCounter)

	// Increment counter for next call
	v.TempVariableCounter++

	return temporaryVariable
}

// generateQuadruple creates a new quadruple and adds it to the quadruples list
// A quadruple represents an intermediate code instruction with operator and operands
// Parameters:
//   - operator: The operation to perform (e.g., "+", "-", "=", "GOTO")
//   - leftOperand: The first operand of the operation
//   - rightOperand: The second operand of the operation (maybe empty)
//   - result: The destination where the result will be stored
func (v *Visitor) generateQuadruple(operator interface{}, leftOperand interface{}, rightOperand interface{}, result string) {
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
func (v *Visitor) updateQuadruple(index int, result string) {
	quadruple := v.Quadruples[index]
	quadruple.Result = result
	v.Quadruples[index] = quadruple
}

func (v *Visitor) PrintQuadruplesTable() {
	fmt.Println("\n=== Quadruples Table ===")
	fmt.Println("Index | Operator | Operand 1 | Operand 2 |      Result     | Scope")
	fmt.Println("-------|----------|------------|------------|--------------------|----------")

	for i, quad := range v.Quadruples {
		fmt.Printf("%6d | %8s | %10s | %10s | %18s | %s\n",
			i, quad.Operator, quad.LeftOperand, quad.RightOperand, quad.Result, quad.Scope)
	}

	fmt.Println("=============================================================")
}

// Visit is the main entry point for traversing the parse tree
// It dispatches to the appropriate visitor method based on the node type
func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
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
func (v *Visitor) VisitProgram(ctx *generated.ProgramContext) interface{} {
	// Push program scope to the stack
	v.CurrentScope.Push("Program")

	// Generate program start quadruple
	v.generateQuadruple("PROGRAM", "", "", "")

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
	v.generateQuadruple("END", "", "", "")

	return nil
}

// VisitProgramBody processes the main code block of the program
func (v *Visitor) VisitProgramBody(ctx *generated.ProgramBodyContext) interface{} {
	return v.Visit(ctx.CodeBlock())
}

// VisitFunctionDeclaration processes function declarations
// Creates a new scope and generates function-related quadruples
func (v *Visitor) VisitFunctionDeclaration(ctx *generated.FunctionDeclarationContext) interface{} {
	// Get function name and push to scope stack
	functionName := ctx.Identifier().GetText()
	v.TempVariableCounter = 0
	v.CurrentScope.Push(functionName)

	// Generate function declaration quadruple
	v.generateQuadruple("FUNC", functionName, "", "")

	// Process function body
	return v.Visit(ctx.FunctionBody())
}

// VisitFunctionBody processes the code block within a function
// Handles scope management and generates end function quadruple
func (v *Visitor) VisitFunctionBody(ctx *generated.FunctionBodyContext) interface{} {
	// Process function code block
	v.Visit(ctx.CodeBlock())

	// Pop function scope from stack
	v.CurrentScope.Pop()

	// Generate end function quadruple
	v.generateQuadruple("ENDFUNC", "", "", "")
	v.TempVariableCounter = 0
	return nil
}

// VisitCodeBlock processes a block of statements
// Visits each statement in the block sequentially
func (v *Visitor) VisitCodeBlock(ctx *generated.CodeBlockContext) interface{} {
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
func (v *Visitor) VisitStatement(ctx *generated.StatementContext) interface{} {
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
func (v *Visitor) VisitAssignment(ctx *generated.AssignmentContext) interface{} {
	// Evaluate the expression on the right side
	expressionResult := v.Visit(ctx.Expression())

	// Get the identifier on the left side
	variableIdentifier := ctx.Identifier().GetText()

	// Generate assignment quadruple
	v.generateQuadruple("=", expressionResult.(string), "", variableIdentifier)

	return nil
}

// VisitConditional processes if-else conditional structures
// Generates appropriate quadruples for conditional branching
func (v *Visitor) VisitConditional(ctx *generated.ConditionalContext) interface{} {
	// Process if block
	if ctx.IfBlock() != nil {
		// Evaluate the condition expression
		v.Visit(ctx.IfBlock().ParenthesizedExpression())

		// Save position for GOTOF instruction that will be filled later
		jumpPosition := len(v.Quadruples)
		v.JumpsStack.Push(jumpPosition)
		v.generateQuadruple("GOTOF", "", "", "")

		// Process the code block within the if statement
		v.Visit(ctx.IfBlock().CodeBlock())
	}

	// Process else block if it exists
	if ctx.ElseBlock() != nil {
		// Get the position of if's GOTOF to update it
		falseJumpPosition := v.JumpsStack.Pop()

		// Add GOTO at the end of if block to skip else block
		v.JumpsStack.Push(len(v.Quadruples))
		v.generateQuadruple("GOTO", "", "", "")

		// Update the GOTOF to jump to the beginning of the else block
		v.updateQuadruple(falseJumpPosition.(int), strconv.Itoa(len(v.Quadruples)))

		// Process the code block within the else statement
		v.Visit(ctx.ElseBlock().CodeBlock())
	}

	// Complete the conditional structure
	if ctx.SEMICOLON() != nil {
		// Get the last jump position (either from if or else) and update it
		lastJumpPosition := v.JumpsStack.Pop()
		v.updateQuadruple(lastJumpPosition.(int), strconv.Itoa(len(v.Quadruples)))
	}
	return nil
}

// VisitLoop processes loop structures (currently supports while loops)
// Handles the control flow for iterative execution
func (v *Visitor) VisitLoop(ctx *generated.LoopContext) interface{} {
	if ctx.WHILE() != nil {
		// Save the position to return to for loop condition evaluation
		loopStartPosition := len(v.Quadruples)
		v.JumpsStack.Push(loopStartPosition)

		// Evaluate the loop condition
		v.Visit(ctx.ParenthesizedExpression())

		// Save position for GOTOF that will exit the loop when condition is false
		conditionalJumpPosition := len(v.Quadruples)
		v.JumpsStack.Push(conditionalJumpPosition)
		v.generateQuadruple("GOTOF", "", "", "")

		// Process the loop body
		v.Visit(ctx.CodeBlock())

		// Retrieve jump positions from stack
		exitJumpPosition := v.JumpsStack.Pop()
		returnToConditionPosition := v.JumpsStack.Pop()

		// Update the conditional jump to point to after the loop
		v.updateQuadruple(exitJumpPosition.(int), strconv.Itoa(len(v.Quadruples)+1))

		// Add a jump back to the condition evaluation
		v.generateQuadruple("GOTO", "", "", strconv.Itoa(returnToConditionPosition.(int)))
	}
	return nil
}

// VisitFunctionCall processes function calls in the code
// Generates PARAM quadruples for arguments and a CALL quadruple
func (v *Visitor) VisitFunctionCall(ctx *generated.FunctionCallContext) interface{} {
	// Process all arguments
	var argumentValues []string
	for _, expression := range ctx.ArgumentList().AllExpression() {
		argumentValues = append(argumentValues, v.Visit(expression).(string))
	}

	// Generate parameter quadruples for each argument
	for _, argValue := range argumentValues {
		v.generateQuadruple("PARAM", argValue, "", "")
	}

	// Generate the function call quadruple
	functionName := ctx.Identifier().GetText()
	v.generateQuadruple("CALL", functionName, argumentValues, "")
	return nil
}

// VisitPrintStatement processes print statements
// Generates PRINT quadruples for each printable item
func (v *Visitor) VisitPrintStatement(ctx *generated.PrintStatementContext) interface{} {
	for i, printable := range ctx.AllPrintable() {
		if printable.Expression() != nil {
			// If it's an expression, evaluate it and print the result
			expressionResult := v.Visit(printable.Expression())
			v.generateQuadruple("PRINT", "", "", expressionResult.(string))
		} else {
			// If it's a literal, print it directly
			literalValue := ctx.Printable(i).GetText()
			v.generateQuadruple("PRINT", "", "", literalValue)
		}
	}
	return nil
}

// VisitExpression processes expressions, including relational operations
// Returns a reference to the result (variable or temporary)
func (v *Visitor) VisitExpression(ctx *generated.ExpressionContext) interface{} {
	// Process the left-hand arithmetic expression
	leftOperand := v.Visit(ctx.ArithmeticExpression())

	// Check if there's a relational component
	if ctx.RelationalExpression() != nil {
		// Get the relational operator
		relationalOperator := v.Visit(ctx.RelationalExpression().RelationalOperator())

		// Process the right-hand arithmetic expression
		rightOperand := v.Visit(ctx.RelationalExpression().ArithmeticExpression())

		// Create a temporary variable to store the result
		resultTemp := v.newTemporaryVariable()
		v.generateQuadruple(relationalOperator.(string), leftOperand.(string), rightOperand.(string), resultTemp)
		return resultTemp
	}

	// If no relational component, return the arithmetic result
	return leftOperand
}

// VisitRelationalOperator identifies and returns the string representation
// of the relational operator
func (v *Visitor) VisitRelationalOperator(ctx *generated.RelationalOperatorContext) interface{} {
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
func (v *Visitor) VisitArithmeticExpression(ctx *generated.ArithmeticExpressionContext) interface{} {
	// Process the first term
	result := v.Visit(ctx.Term(0))

	// Process additional terms if they exist
	for i := range ctx.AllAdditiveOperator() {
		operator := v.Visit(ctx.AdditiveOperator(i))
		nextTerm := v.Visit(ctx.Term(i + 1))

		// Create a temporary variable for the result
		resultTemp := v.newTemporaryVariable()
		v.generateQuadruple(operator.(string), result.(string), nextTerm.(string), resultTemp)
		return resultTemp
	}

	return result
}

// VisitAdditiveOperator identifies and returns the string representation
// of the additive operator (+ or -)
func (v *Visitor) VisitAdditiveOperator(ctx *generated.AdditiveOperatorContext) interface{} {
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
func (v *Visitor) VisitTerm(ctx *generated.TermContext) interface{} {
	// Process the first factor
	result := v.Visit(ctx.Factor(0))

	// Process additional factors if they exist
	for i := range ctx.AllMultiplicativeOperator() {
		operator := v.Visit(ctx.MultiplicativeOperator(i))
		nextFactor := v.Visit(ctx.Factor(i + 1))

		// Create a temporary variable for the result
		resultTemp := v.newTemporaryVariable()
		v.generateQuadruple(operator.(string), result.(string), nextFactor.(string), resultTemp)
		return resultTemp
	}

	return result
}

// VisitMultiplicativeOperator identifies and returns the string representation
// of the multiplicative operator (* or /)
func (v *Visitor) VisitMultiplicativeOperator(ctx *generated.MultiplicativeOperatorContext) interface{} {
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
func (v *Visitor) VisitFactor(ctx *generated.FactorContext) interface{} {
	if ctx.ParenthesizedExpression() != nil {
		return v.Visit(ctx.ParenthesizedExpression())
	}
	return v.Visit(ctx.ValueWithOptionalSign())
}

// VisitParenthesizedExpression processes expressions within parentheses
// Returns a reference to the result
func (v *Visitor) VisitParenthesizedExpression(ctx *generated.ParenthesizedExpressionContext) interface{} {
	return v.Visit(ctx.Expression())
}

// VisitValueWithOptionalSign processes values that may have a leading sign
// Returns a reference to the result (variable or temporary)
func (v *Visitor) VisitValueWithOptionalSign(ctx *generated.ValueWithOptionalSignContext) interface{} {
	// Process the base value
	valueResult := v.Visit(ctx.Value())

	// Check if there's a negative sign
	if ctx.AdditiveOperator() != nil && v.Visit(ctx.AdditiveOperator()) == ("-") {
		// Create a temporary for the negated value
		resultTemp := v.newTemporaryVariable()
		v.generateQuadruple("NEG", valueResult.(string), "", resultTemp)
		return resultTemp
	}

	return valueResult
}

// VisitValue processes basic values (identifiers or constants)
// Returns the identifier or constant value as a string
func (v *Visitor) VisitValue(ctx *generated.ValueContext) interface{} {
	if ctx.Identifier() != nil {
		return ctx.Identifier().GetText()
	} else {
		return ctx.Constant().GetText()
	}
}
