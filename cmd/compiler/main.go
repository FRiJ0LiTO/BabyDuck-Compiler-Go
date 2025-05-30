package main

import (
	"BabyDuck/internal/parser"
	"BabyDuck/internal/semantic"
	"BabyDuck/internal/virtualMachine"
	"fmt"
	"os"
)

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
	newParser := parser.NewParser(string(sourceCode), false)
	parseTree, symbolTable, errors := newParser.Parse()

	if len(errors) > 0 {
		fmt.Println("Compilation errors:")
		for _, err := range errors {
			fmt.Printf("- %s\n", err)
		}
	}

	// Quadruples
	visitor := semantic.NewVisitor(symbolTable, false)
	visitor.Visit(parseTree)

	// Virtual Machine
	vm := virtualMachine.NewVirtualMachine(*symbolTable, visitor.Quadruples)
	vm.Execute()
}
