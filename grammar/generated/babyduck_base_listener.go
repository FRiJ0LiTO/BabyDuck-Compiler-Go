// Code generated from ./BabyDuck.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // BabyDuck

import "github.com/antlr4-go/antlr/v4"

// BaseBabyDuckListener is a complete listener for a parse tree produced by BabyDuckParser.
type BaseBabyDuckListener struct{}

var _ BabyDuckListener = &BaseBabyDuckListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBabyDuckListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBabyDuckListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBabyDuckListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBabyDuckListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseBabyDuckListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseBabyDuckListener) ExitProgram(ctx *ProgramContext) {}

// EnterProgramHeader is called when production programHeader is entered.
func (s *BaseBabyDuckListener) EnterProgramHeader(ctx *ProgramHeaderContext) {}

// ExitProgramHeader is called when production programHeader is exited.
func (s *BaseBabyDuckListener) ExitProgramHeader(ctx *ProgramHeaderContext) {}

// EnterProgramBody is called when production programBody is entered.
func (s *BaseBabyDuckListener) EnterProgramBody(ctx *ProgramBodyContext) {}

// ExitProgramBody is called when production programBody is exited.
func (s *BaseBabyDuckListener) ExitProgramBody(ctx *ProgramBodyContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseBabyDuckListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseBabyDuckListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterParenthesizedExpression is called when production parenthesizedExpression is entered.
func (s *BaseBabyDuckListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// ExitParenthesizedExpression is called when production parenthesizedExpression is exited.
func (s *BaseBabyDuckListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// EnterType is called when production type is entered.
func (s *BaseBabyDuckListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseBabyDuckListener) ExitType(ctx *TypeContext) {}

// EnterIdList is called when production idList is entered.
func (s *BaseBabyDuckListener) EnterIdList(ctx *IdListContext) {}

// ExitIdList is called when production idList is exited.
func (s *BaseBabyDuckListener) ExitIdList(ctx *IdListContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseBabyDuckListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseBabyDuckListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterVarsSection is called when production varsSection is entered.
func (s *BaseBabyDuckListener) EnterVarsSection(ctx *VarsSectionContext) {}

// ExitVarsSection is called when production varsSection is exited.
func (s *BaseBabyDuckListener) ExitVarsSection(ctx *VarsSectionContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseBabyDuckListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseBabyDuckListener) ExitConstant(ctx *ConstantContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseBabyDuckListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseBabyDuckListener) ExitParameter(ctx *ParameterContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BaseBabyDuckListener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BaseBabyDuckListener) ExitFunctionBody(ctx *FunctionBodyContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseBabyDuckListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseBabyDuckListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFunctionsSection is called when production functionsSection is entered.
func (s *BaseBabyDuckListener) EnterFunctionsSection(ctx *FunctionsSectionContext) {}

// ExitFunctionsSection is called when production functionsSection is exited.
func (s *BaseBabyDuckListener) ExitFunctionsSection(ctx *FunctionsSectionContext) {}

// EnterCodeBlock is called when production codeBlock is entered.
func (s *BaseBabyDuckListener) EnterCodeBlock(ctx *CodeBlockContext) {}

// ExitCodeBlock is called when production codeBlock is exited.
func (s *BaseBabyDuckListener) ExitCodeBlock(ctx *CodeBlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseBabyDuckListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseBabyDuckListener) ExitStatement(ctx *StatementContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseBabyDuckListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseBabyDuckListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterIfBlock is called when production ifBlock is entered.
func (s *BaseBabyDuckListener) EnterIfBlock(ctx *IfBlockContext) {}

// ExitIfBlock is called when production ifBlock is exited.
func (s *BaseBabyDuckListener) ExitIfBlock(ctx *IfBlockContext) {}

// EnterElseBlock is called when production elseBlock is entered.
func (s *BaseBabyDuckListener) EnterElseBlock(ctx *ElseBlockContext) {}

// ExitElseBlock is called when production elseBlock is exited.
func (s *BaseBabyDuckListener) ExitElseBlock(ctx *ElseBlockContext) {}

// EnterConditional is called when production conditional is entered.
func (s *BaseBabyDuckListener) EnterConditional(ctx *ConditionalContext) {}

// ExitConditional is called when production conditional is exited.
func (s *BaseBabyDuckListener) ExitConditional(ctx *ConditionalContext) {}

// EnterLoop is called when production loop is entered.
func (s *BaseBabyDuckListener) EnterLoop(ctx *LoopContext) {}

// ExitLoop is called when production loop is exited.
func (s *BaseBabyDuckListener) ExitLoop(ctx *LoopContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseBabyDuckListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseBabyDuckListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseBabyDuckListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseBabyDuckListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterPrintable is called when production printable is entered.
func (s *BaseBabyDuckListener) EnterPrintable(ctx *PrintableContext) {}

// ExitPrintable is called when production printable is exited.
func (s *BaseBabyDuckListener) ExitPrintable(ctx *PrintableContext) {}

// EnterPrintStatement is called when production printStatement is entered.
func (s *BaseBabyDuckListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production printStatement is exited.
func (s *BaseBabyDuckListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBabyDuckListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBabyDuckListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRelationalOperator is called when production relationalOperator is entered.
func (s *BaseBabyDuckListener) EnterRelationalOperator(ctx *RelationalOperatorContext) {}

// ExitRelationalOperator is called when production relationalOperator is exited.
func (s *BaseBabyDuckListener) ExitRelationalOperator(ctx *RelationalOperatorContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseBabyDuckListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseBabyDuckListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterAdditiveOperator is called when production additiveOperator is entered.
func (s *BaseBabyDuckListener) EnterAdditiveOperator(ctx *AdditiveOperatorContext) {}

// ExitAdditiveOperator is called when production additiveOperator is exited.
func (s *BaseBabyDuckListener) ExitAdditiveOperator(ctx *AdditiveOperatorContext) {}

// EnterArithmeticExpression is called when production arithmeticExpression is entered.
func (s *BaseBabyDuckListener) EnterArithmeticExpression(ctx *ArithmeticExpressionContext) {}

// ExitArithmeticExpression is called when production arithmeticExpression is exited.
func (s *BaseBabyDuckListener) ExitArithmeticExpression(ctx *ArithmeticExpressionContext) {}

// EnterMultiplicativeOperator is called when production multiplicativeOperator is entered.
func (s *BaseBabyDuckListener) EnterMultiplicativeOperator(ctx *MultiplicativeOperatorContext) {}

// ExitMultiplicativeOperator is called when production multiplicativeOperator is exited.
func (s *BaseBabyDuckListener) ExitMultiplicativeOperator(ctx *MultiplicativeOperatorContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseBabyDuckListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseBabyDuckListener) ExitTerm(ctx *TermContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseBabyDuckListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseBabyDuckListener) ExitFactor(ctx *FactorContext) {}

// EnterValueWithOptionalSign is called when production valueWithOptionalSign is entered.
func (s *BaseBabyDuckListener) EnterValueWithOptionalSign(ctx *ValueWithOptionalSignContext) {}

// ExitValueWithOptionalSign is called when production valueWithOptionalSign is exited.
func (s *BaseBabyDuckListener) ExitValueWithOptionalSign(ctx *ValueWithOptionalSignContext) {}

// EnterValue is called when production value is entered.
func (s *BaseBabyDuckListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseBabyDuckListener) ExitValue(ctx *ValueContext) {}
