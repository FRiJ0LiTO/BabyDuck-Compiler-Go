package parser

import (
	"BabyDuck/grammar/generated"
	"BabyDuck/internal/memory"
	"BabyDuck/internal/symbol"
	"fmt"
	"os"
)

// DirectoryBuilder traverses the AST to build a symbol table and validate semantic rules.
// It implements the BaseBabyDuckListener interface to process parse tree events.
type DirectoryBuilder struct {
	generated.BaseBabyDuckListener
	Directory     *symbol.FunctionDirectory // Symbol table storing functions and variables
	MemoryManager *memory.Manager           // // Manages memory allocation for variables, constants
	Errors        []string                  // Collection of semantic errors found during parsing
	Debug         bool                      // Debugging Flag
}

// NewDirectoryBuilder creates and initializes a new DirectoryBuilder instance.
// Returns a pointer to the newly created DirectoryBuilder.
func NewDirectoryBuilder(debug bool) *DirectoryBuilder {
	return &DirectoryBuilder{
		Directory:     symbol.NewFunctionDirectory(),
		MemoryManager: memory.NewMemoryManager(memory.DefaultMemoryConfig),
		Errors:        []string{},
		Debug:         debug,
	}
}

// EnterFunctionDeclaration is called when entering a function declaration node in the parse tree.
// It registers the function in the directory and updates the current scope.
func (d *DirectoryBuilder) EnterFunctionDeclaration(ctx *generated.FunctionDeclarationContext) {
	functionName := ctx.Identifier().GetText()

	err := d.Directory.AddFunction(functionName)
	if err != nil {
		// Function already defined in current scope
		d.Errors = append(d.Errors, err.Error())
	}

	d.Directory.CurrentScope.Push(functionName)
}

// ExitFunctionDeclaration is called when exiting a function declaration node in the parse tree.
// It restores the scope that existed before entering the function.
func (d *DirectoryBuilder) ExitFunctionDeclaration(_ *generated.FunctionDeclarationContext) {
	d.Directory.CurrentScope.Pop()
	d.MemoryManager.ResetLocal()
}

// ExitVarDecl is called when exiting a variable declaration node in the parse tree.
// It registers all variables declared in this statement with their associated type.
func (d *DirectoryBuilder) ExitVarDecl(ctx *generated.VarDeclContext) {
	dataType := memory.DataType(ctx.Type_().GetText())
	identifiers := ctx.IdList().AllID()

	// Get the current scope from the stack
	currentScope := d.Directory.CurrentScope.Peek().(string)
	for _, identifier := range identifiers {
		varName := identifier.GetText()

		virtualAddress, err := d.allocateVirtualMemory(dataType)
		if err != nil {
			d.Errors = append(d.Errors, err.Error())
		}

		// Register Variable
		err = d.Directory.AddVariable(varName, dataType, virtualAddress, currentScope)
		if err != nil {
			// Variable already defined in current scope
			d.Errors = append(d.Errors, err.Error())
		}
	}

	// Determine scope prefix (Global/Local) for resource tracking
	scopePrefix := "Local"
	if currentScope == "program" {
		scopePrefix = "Global"
	}
	resourceType := scopePrefix + string(dataType)
	d.Directory.AddResource(currentScope, resourceType, len(identifiers))
}

// EnterParameter is called when entering a parameter declaration node in the parse tree.
// It registers each function parameter as a variable in the current function's scope.
func (d *DirectoryBuilder) EnterParameter(ctx *generated.ParameterContext) {
	paramName := ctx.Identifier().GetText()
	paramType := ctx.Type_().GetText()

	virtualAddress, err := d.allocateVirtualMemory(memory.DataType(paramType))
	if err != nil {
		d.Errors = append(d.Errors, err.Error())
	}

	scope := d.Directory.CurrentScope.Peek().(string)
	err = d.Directory.AddVariable(paramName, memory.DataType(paramType), virtualAddress, scope)
	if err != nil {
		// Parameter already defined in current function
		d.Errors = append(d.Errors, err.Error())
	}
}

// ExitAssignment is called when exiting an assignment statement node in the parse tree.
// It validates that the variable being assigned to exists in the current scope and validates the expression being assigned.
//
// Behavior:
//   - Extracts the variable name from the assignment context.
//   - Collects all active scopes starting from the current scope.
//   - Validates if the variable exists in any of the collected scopes using the symbol table.
//   - If the variable does not exist, appends an error message to the Errors list.
//   - If the assignment includes an expression, validates the expression.
func (d *DirectoryBuilder) ExitAssignment(ctx *generated.AssignmentContext) {
	variableName := ctx.Identifier().GetText()

	scopes := d.Directory.CurrentScope.ToStringSlice()
	_, exists := d.Directory.FindVariable(scopes, variableName)
	if !exists {
		d.Errors = append(d.Errors, "error: undefined variable ", variableName)
	}

	if ctx.Expression() != nil {
		d.validateExpression(ctx.Expression())
	}
}

// ExitParenthesizedExpression is called when exiting a parenthesized expression node.
// It validates the expression within the parentheses.
func (d *DirectoryBuilder) ExitParenthesizedExpression(ctx *generated.ParenthesizedExpressionContext) {
	if ctx.Expression() != nil {
		d.validateExpression(ctx.Expression())
	}
}

// ExitFunctionCall is called when exiting a function call node in the parse tree.
// It validates each argument expression in the function call.
func (d *DirectoryBuilder) ExitFunctionCall(ctx *generated.FunctionCallContext) {
	functionName := ctx.Identifier().GetText()

	exists := d.Directory.HasFunction(functionName)
	if !exists {
		fmt.Printf("error: undefined function '%s'\n", functionName)
		os.Exit(1)
	}

	if ctx.ArgumentList() != nil {
		for _, context := range ctx.ArgumentList().AllExpression() {
			d.validateExpression(context)
		}
	}
}

// ExitPrintStatement is called when exiting a print statement node in the parse tree.
// It validates each printable expression in the statement.
func (d *DirectoryBuilder) ExitPrintStatement(ctx *generated.PrintStatementContext) {
	for _, printable := range ctx.AllPrintable() {
		if printable.CONST_STRING() != nil {
			if d.Debug {
				fmt.Println("Found string literal:", printable.CONST_STRING().GetText())
			}
			_ = d.registerConstantInMemory(printable.CONST_STRING().GetText(), "string")
		}

		if printable.Expression() != nil {
			d.validateExpression(printable.Expression())
		}
	}
}

// validateExpression validates an expression node by delegating to the appropriate
// specialized validation method based on the expression type.
func (d *DirectoryBuilder) validateExpression(exp generated.IExpressionContext) {
	if exp.ArithmeticExpression() != nil {
		d.validateArithmeticExpression(exp.ArithmeticExpression())
	}

	if exp.RelationalExpression() != nil {
		d.validateRelationalExpression(exp.RelationalExpression())
	}
}

// validateRelationalExpression validates a relational expression node by validating
// its arithmetic expression components.
func (d *DirectoryBuilder) validateRelationalExpression(exp generated.IRelationalExpressionContext) {
	if exp.ArithmeticExpression() != nil {
		d.validateArithmeticExpression(exp.ArithmeticExpression())
	}
}

// validateArithmeticExpression validates an arithmetic expression by validating each
// of its term components.
func (d *DirectoryBuilder) validateArithmeticExpression(exp generated.IArithmeticExpressionContext) {
	for _, term := range exp.AllTerm() {
		d.validateTerm(term)
	}
}

// validateTerm validates a term by validating each of its factor components.
func (d *DirectoryBuilder) validateTerm(term generated.ITermContext) {
	for _, factor := range term.AllFactor() {
		d.validateFactor(factor)
	}
}

// validateFactor validates a factor based on whether it's a parenthesized expression
// or a value with an optional sign.
func (d *DirectoryBuilder) validateFactor(factor generated.IFactorContext) {
	if factor.ParenthesizedExpression() != nil {
		if factor.ParenthesizedExpression().Expression() != nil {
			d.validateExpression(factor.ParenthesizedExpression().Expression())
		}
	}

	if factor.ValueWithOptionalSign() != nil {
		d.validateValueWithOptionalSign(factor.ValueWithOptionalSign())
	}
}

// validateValueWithOptionalSign validates a value (which may have a sign) by determining
// whether it is an identifier (variable) or a constant and processing it accordingly.
//
// Parameters:
//   - value (generated.IValueWithOptionalSignContext): The context containing the value to validate.
//
// Behavior:
//   - If the value is an identifier, it validates the variable by checking its existence in the current scope.
//   - If the value is a constant, it registers the constant in the memory manager and symbol table.
func (d *DirectoryBuilder) validateValueWithOptionalSign(value generated.IValueWithOptionalSignContext) {
	if value.Value().Identifier() != nil {
		d.validateVariable(value)
	} else {
		d.registerConstant(value)
	}
}

// validateVariable checks if a variable exists in the current scope or any parent scope.
//
// Parameters:
//   - value (generated.IValueWithOptionalSignContext): The context containing the variable to validate.
//
// Behavior:
//   - Extracts the variable name from the provided context.
//   - Collects all active scopes starting from the current scope.
//   - Validates if the variable exists in any of the collected scopes using the symbol table.
//   - If the variable does not exist, appends an error message to the Errors list.
func (d *DirectoryBuilder) validateVariable(value generated.IValueWithOptionalSignContext) {
	variableName := value.Value().Identifier().GetText()

	scopes := d.Directory.CurrentScope.ToStringSlice()
	_, exists := d.Directory.FindVariable(scopes, variableName)
	if !exists {
		d.Errors = append(d.Errors, "error: undefined variable ", variableName)
	}
}

// registerConstant extracts and registers a constant (integer or float), considering its sign.
//
// Parameters:
//   - value (generated.IValueWithOptionalSignContext): The context containing the constant, possibly with a sign.
//
// Behavior:
//   - Determines whether the constant is negative by checking for a subtraction operator.
//   - Extracts the raw constant value (e.g., integer or float).
//   - Passes the constant value and its type to registerConstantInMemory for registration.
//   - If an error occurs during registration, it is appended to the Errors list.
func (d *DirectoryBuilder) registerConstant(value generated.IValueWithOptionalSignContext) {

	constant := value.Value().Constant()

	if d.Debug {
		fmt.Println("Processing constant:", constant.GetText())
	}

	var err error
	if constant.CONST_INT() != nil {
		constantVal := constant.CONST_INT().GetText()
		err = d.registerConstantInMemory(constantVal, memory.Integer)
	} else {
		constantVal := constant.CONST_FLOAT().GetText()
		err = d.registerConstantInMemory(constantVal, memory.Float)
	}

	if err != nil {
		d.Errors = append(d.Errors, err.Error())
	}
}

// allocateVirtualMemory allocates a virtual memory address for a variable based on its type and scope
//
// This function determines the variable's scope (global or local) and requests a memory address
// from the memory manager accordingly.
//
// Parameters:
//   - dataType: The data type of the variable (e.g., int, float, string)
//
// Returns:
//   - int: The allocated virtual memory address
//   - error: An error if address allocation fails
func (d *DirectoryBuilder) allocateVirtualMemory(dataType memory.DataType) (int, error) {
	// Determine the variable scope (Global or Local) based on current context
	var scopeType memory.VarType
	currentScope := d.Directory.CurrentScope.Peek()

	if currentScope == "program" {
		scopeType = memory.Global
	} else {
		scopeType = memory.Local
	}

	// Request a virtual memory address from the memory manager
	virtualAddress, err := d.MemoryManager.GetAddress(scopeType, dataType)
	if err != nil {
		return -1, fmt.Errorf("failed to allocate memory: %w", err)
	}

	return virtualAddress, nil
}

// registerConstantInMemory registers a constant in the constant memory segment and the symbol table.
//
// Parameters:
//   - constantValue (string): The textual representation of the constant value.
//   - dataType (memory.DataType): The data type of the constant (e.g., Integer, Float).
//   - isNegative (bool): Indicates whether the constant should be treated as negative.
//
// Returns:
//   - error: An error if memory allocation or registration fails.
//
// Behavior:
//   - Prepends a "-" to the constant value if isNegative is true.
//   - Requests a virtual address from the memory manager for the constant segment.
//   - Registers the constant in the symbol table using the assigned address.

func (d *DirectoryBuilder) registerConstantInMemory(constantValue string, dataType memory.DataType) error {

	// Allocate memory address in the constant segment
	virtualAddress, err := d.MemoryManager.GetAddress(memory.Constant, dataType)
	if err != nil {
		return fmt.Errorf("failed to allocate constant memory: %w", err)
	}

	// Register the constant in the directory/symbol table
	err = d.Directory.AddConstant(constantValue, dataType, virtualAddress)
	if err != nil {
		return fmt.Errorf("failed to register constant: %w", err)
	}

	return nil
}
