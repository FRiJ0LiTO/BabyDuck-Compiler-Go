package parser

import (
 "BabyDuck/grammar/generated"
 "BabyDuck/internal/symbol"
 "github.com/antlr4-go/antlr/v4"
)

// Parser handles the lexical analysis and parsing of BabyDuck source code.
// It coordinates the ANTLR-generated components and symbol table construction.
type Parser struct {
 inputStream antlr.CharStream             // Source code input stream
 lexer       *generated.BabyDuckLexer     // Lexical analyzer
 tokenStream *antlr.CommonTokenStream     // Stream of tokens from lexer
 parser      *generated.BabyDuckParser    // Syntactic parser
 builder     *DirectoryBuilder            // Symbol table constructor
}

// NewParser creates a new Parser instance for the given source code.
// It initializes all necessary ANTLR components for lexical and syntactic analysis.
//
// Parameters:
//   sourceCode (string): The source code to be parsed.
//
// Returns:
//   *Parser: A pointer to the newly created Parser instance.
func NewParser(sourceCode string, debug bool) *Parser {
 inputStream := antlr.NewInputStream(sourceCode)
 lexer := generated.NewBabyDuckLexer(inputStream)
 tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
 parser := generated.NewBabyDuckParser(tokenStream)

 return &Parser{
  inputStream: inputStream,
  lexer:       lexer,
  tokenStream: tokenStream,
  parser:      parser,
  builder:     NewDirectoryBuilder(debug),
 }
}

// Parse performs the complete analysis of the source code.
// It walks the parse tree to build the symbol table and collect semantic errors.
//
// Returns:
//   *symbol.FunctionDirectory: The constructed function directory containing symbol tables.
//   []string: A list of errors encountered during parsing.
func (p *Parser) Parse() ( generated.IProgramContext, *symbol.FunctionDirectory, []string) {
 parseTree := p.parser.Program()
 antlr.ParseTreeWalkerDefault.Walk(p.builder, parseTree)

 return parseTree, p.builder.Directory, p.builder.Errors
}