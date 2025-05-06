package main

import (
	"BabyDuck/internal/parser"
	"BabyDuck/internal/semantic"
	"BabyDuck/internal/symbol"
	"fmt"
	"os"
)

// printSymbolTable displays the contents of the function directory,
// showing all variables and their details for each scope.
func printSymbolTable(directory *symbol.FunctionDirectory) {
	for scope, variables := range directory.Directory {
		fmt.Printf("Scope (Function Name): %s\n", scope)
		for varName, varDetails := range variables {
			fmt.Printf("Variable: %s, Details: %v\n", varName, varDetails)
		}
		fmt.Println()
	}
}

func main() {
	// Check if a source file was provided
	if len(os.Args) < 2 {
		fmt.Println("Error: source file path required")
		os.Exit(1)
	}

	sourceFile := os.Args[1]

	// Read the source file
	sourceCode, err := os.ReadFile(sourceFile)
	if err != nil {
		fmt.Printf("Error reading source file: %v\n", err)
		os.Exit(1)
	}

	// Parse the source code
	newParser := parser.NewParser(string(sourceCode))
	symbolTable, errors := newParser.Parse()

	// Test Semantic Cube
	semanticCube := semantic.NewSemanticCube()
	fmt.Println(semanticCube.GetResultType("int", "int", "+"))

	fmt.Println("Function Directory built successfully")
	printSymbolTable(symbolTable)

	if len(errors) > 0 {
		fmt.Println("Compilation errors:")
		for _, err := range errors {
			fmt.Printf("- %s\n", err)
		}
	}
}
