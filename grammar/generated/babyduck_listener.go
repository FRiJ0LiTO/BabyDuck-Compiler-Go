// Code generated from ./BabyDuck.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // BabyDuck

import "github.com/antlr4-go/antlr/v4"

// BabyDuckListener is a complete listener for a parse tree produced by BabyDuckParser.
type BabyDuckListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterProgramHeader is called when entering the programHeader production.
	EnterProgramHeader(c *ProgramHeaderContext)

	// EnterProgramBody is called when entering the programBody production.
	EnterProgramBody(c *ProgramBodyContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterParenthesizedExpression is called when entering the parenthesizedExpression production.
	EnterParenthesizedExpression(c *ParenthesizedExpressionContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterIdList is called when entering the idList production.
	EnterIdList(c *IdListContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterVarsSection is called when entering the varsSection production.
	EnterVarsSection(c *VarsSectionContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFunctionsSection is called when entering the functionsSection production.
	EnterFunctionsSection(c *FunctionsSectionContext)

	// EnterCodeBlock is called when entering the codeBlock production.
	EnterCodeBlock(c *CodeBlockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterIfBlock is called when entering the ifBlock production.
	EnterIfBlock(c *IfBlockContext)

	// EnterElseBlock is called when entering the elseBlock production.
	EnterElseBlock(c *ElseBlockContext)

	// EnterConditional is called when entering the conditional production.
	EnterConditional(c *ConditionalContext)

	// EnterLoop is called when entering the loop production.
	EnterLoop(c *LoopContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterPrintable is called when entering the printable production.
	EnterPrintable(c *PrintableContext)

	// EnterPrintStatement is called when entering the printStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRelationalOperator is called when entering the relationalOperator production.
	EnterRelationalOperator(c *RelationalOperatorContext)

	// EnterRelationalExpression is called when entering the relationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterAdditiveOperator is called when entering the additiveOperator production.
	EnterAdditiveOperator(c *AdditiveOperatorContext)

	// EnterArithmeticExpression is called when entering the arithmeticExpression production.
	EnterArithmeticExpression(c *ArithmeticExpressionContext)

	// EnterMultiplicativeOperator is called when entering the multiplicativeOperator production.
	EnterMultiplicativeOperator(c *MultiplicativeOperatorContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterValueWithOptionalSign is called when entering the valueWithOptionalSign production.
	EnterValueWithOptionalSign(c *ValueWithOptionalSignContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitProgramHeader is called when exiting the programHeader production.
	ExitProgramHeader(c *ProgramHeaderContext)

	// ExitProgramBody is called when exiting the programBody production.
	ExitProgramBody(c *ProgramBodyContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitParenthesizedExpression is called when exiting the parenthesizedExpression production.
	ExitParenthesizedExpression(c *ParenthesizedExpressionContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitIdList is called when exiting the idList production.
	ExitIdList(c *IdListContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitVarsSection is called when exiting the varsSection production.
	ExitVarsSection(c *VarsSectionContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFunctionsSection is called when exiting the functionsSection production.
	ExitFunctionsSection(c *FunctionsSectionContext)

	// ExitCodeBlock is called when exiting the codeBlock production.
	ExitCodeBlock(c *CodeBlockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitIfBlock is called when exiting the ifBlock production.
	ExitIfBlock(c *IfBlockContext)

	// ExitElseBlock is called when exiting the elseBlock production.
	ExitElseBlock(c *ElseBlockContext)

	// ExitConditional is called when exiting the conditional production.
	ExitConditional(c *ConditionalContext)

	// ExitLoop is called when exiting the loop production.
	ExitLoop(c *LoopContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitPrintable is called when exiting the printable production.
	ExitPrintable(c *PrintableContext)

	// ExitPrintStatement is called when exiting the printStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRelationalOperator is called when exiting the relationalOperator production.
	ExitRelationalOperator(c *RelationalOperatorContext)

	// ExitRelationalExpression is called when exiting the relationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitAdditiveOperator is called when exiting the additiveOperator production.
	ExitAdditiveOperator(c *AdditiveOperatorContext)

	// ExitArithmeticExpression is called when exiting the arithmeticExpression production.
	ExitArithmeticExpression(c *ArithmeticExpressionContext)

	// ExitMultiplicativeOperator is called when exiting the multiplicativeOperator production.
	ExitMultiplicativeOperator(c *MultiplicativeOperatorContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitValueWithOptionalSign is called when exiting the valueWithOptionalSign production.
	ExitValueWithOptionalSign(c *ValueWithOptionalSignContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
