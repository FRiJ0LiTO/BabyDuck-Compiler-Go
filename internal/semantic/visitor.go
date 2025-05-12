package semantic

import (
	"BabyDuck/grammar/generated"
	"BabyDuck/internal/symbol"
	"BabyDuck/structures/stack"
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

