// Code generated from ./BabyDuck.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // BabyDuck

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by BabyDuckParser.
type BabyDuckVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BabyDuckParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#programHeader.
	VisitProgramHeader(ctx *ProgramHeaderContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#programBody.
	VisitProgramBody(ctx *ProgramBodyContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#parenthesizedExpression.
	VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#idList.
	VisitIdList(ctx *IdListContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#varDecl.
	VisitVarDecl(ctx *VarDeclContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#varsSection.
	VisitVarsSection(ctx *VarsSectionContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#functionBody.
	VisitFunctionBody(ctx *FunctionBodyContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#functionsSection.
	VisitFunctionsSection(ctx *FunctionsSectionContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#codeBlock.
	VisitCodeBlock(ctx *CodeBlockContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#ifBlock.
	VisitIfBlock(ctx *IfBlockContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#elseBlock.
	VisitElseBlock(ctx *ElseBlockContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#conditional.
	VisitConditional(ctx *ConditionalContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#loop.
	VisitLoop(ctx *LoopContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#printable.
	VisitPrintable(ctx *PrintableContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#printStatement.
	VisitPrintStatement(ctx *PrintStatementContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#relationalOperator.
	VisitRelationalOperator(ctx *RelationalOperatorContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#additiveOperator.
	VisitAdditiveOperator(ctx *AdditiveOperatorContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#arithmeticExpression.
	VisitArithmeticExpression(ctx *ArithmeticExpressionContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#multiplicativeOperator.
	VisitMultiplicativeOperator(ctx *MultiplicativeOperatorContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#valueWithOptionalSign.
	VisitValueWithOptionalSign(ctx *ValueWithOptionalSignContext) interface{}

	// Visit a parse tree produced by BabyDuckParser#value.
	VisitValue(ctx *ValueContext) interface{}
}
