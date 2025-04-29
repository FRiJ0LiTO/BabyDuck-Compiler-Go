package main

import (
	"BabyDuck/parser"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"os"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewBabyDuckLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parserBabyDuck := parser.NewBabyDuckParser(stream)
	fmt.Println("Parser", parserBabyDuck.GetCurrentToken().GetText())
	parserBabyDuck.BuildParseTrees = true

	tree := parserBabyDuck.Program()

	fmt.Println(tree.ToStringTree(nil, parserBabyDuck))

}
