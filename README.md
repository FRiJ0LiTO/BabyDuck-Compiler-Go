# BabyDuck Language Compiler

A complete compiler implementation for the BabyDuck programming language, built using Go and ANTLR4. This project provides lexical analysis, syntax parsing, semantic analysis, intermediate code generation, and virtual machine execution.

## Features

- **Lexical Analysis**: Tokenizes BabyDuck source code using ANTLR4-generated lexer
- **Syntax Analysis**: Parses tokens into an Abstract Syntax Tree (AST)
- **Semantic Analysis**: Type checking, variable declaration validation, and scope management
- **Intermediate Code Generation**: Generates quadruples for virtual machine execution
- **Virtual Machine**: Executes intermediate code with memory management and function calls
- **Symbol Table Management**: Tracks variables, functions, and constants across scopes
- **Memory Management**: Virtual memory allocation with different segments for variables and constants

## Architecture

### Core Components

1. **Grammar Definition** (`grammar/BabyDuck.g4`)
   - ANTLR4 grammar specification for the BabyDuck language
   - Defines tokens, keywords, operators, and language syntax rules

2. **Parser** (`internal/parser/`)
   - Lexical and syntactic analysis
   - Symbol table construction during parsing
   - Semantic error detection

3. **Semantic Analyzer** (`internal/semantic/`)
   - Type checking using semantic cube
   - Intermediate code generation (quadruples)
   - Expression evaluation and validation

4. **Virtual Machine** (`internal/virtualMachine/`)
   - Executes intermediate code quadruples
   - Memory management with global/local scope separation
   - Function call stack management

5. **Memory Management** (`internal/memory/`)
   - Virtual memory allocation
   - Separate address spaces for different data types
   - Memory configuration and address management

6. **Symbol Table** (`internal/symbol/`)
   - Function directory management
   - Variable and constant tracking
   - Scope-based symbol resolution

## Getting Started

### Prerequisites

- Go 1.24 or higher
- ANTLR4 (for grammar modifications)

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd BabyDuck
```

2. Install dependencies:
```bash
go mod download
```

3. Build the project:
```bash
go build -o babyduck cmd/compiler/main.go
```

### Usage

Compile and execute a BabyDuck program:

```bash
go run cmd/compiler/main.go <source-file>
```

Or using the built executable:

```bash
./babyduck <source-file>
```

### Example Programs

The `tests/` directory contains several example programs:

- **Basic Program** (`tests/main`): Demonstrates variables, arithmetic, conditionals, loops, and function calls
- **Factorial** (`tests/factorial`): Recursive factorial calculation
- **Fibonacci** (`tests/fibonacci`): Fibonacci sequence generation
- **Function Execution** (`tests/executeFunctions`): Multiple function calls and parameter passing

## BabyDuck Language Syntax

### Program Structure

```babyduck
program ProgramName;

var
    variable1, variable2: int;
    variable3: float;

void functionName(parameter: int) [
    var localVar: int;
    {
        // function body
    }
];

main
{
    // main program body
}
end
```

### Data Types

- `int`: Integer numbers
- `float`: Floating-point numbers
- `string`: String literals (for print statements)

### Variables

```babyduck
var
    counter, sum: int;
    average, total: float;
```

### Functions

```babyduck
void functionName(param1: int, param2: float) [
    var localVariable: int;
    {
        // function implementation
    }
];
```

### Control Structures

#### Conditionals
```babyduck
if (condition) {
    // statements
} else {
    // statements
};
```

#### Loops
```babyduck
while (condition) do {
    // statements
};
```

### Operators

#### Arithmetic Operators
- `+` Addition
- `-` Subtraction  
- `*` Multiplication
- `/` Division

#### Relational Operators
- `>` Greater than
- `<` Less than
- `!=` Not equal

#### Assignment Operator
- `=` Assignment

### Print Statement

```babyduck
print("Hello World");
print(variable1, variable2);
print("Result is", result);
```

## Project Structure

```
.
├── cmd/
│   └── compiler/
│       └── main.go              # Main entry point
├── grammar/
│   └── BabyDuck.g4             # ANTLR4 grammar definition
├── internal/
│   ├── memory/
│   │   ├── allocator.go        # Memory configuration and operators
│   │   └── manager.go          # Memory management system
│   ├── parser/
│   │   ├── builder.go          # Symbol table builder
│   │   └── parser.go           # Parser implementation
│   ├── semantic/
│   │   ├── cube.go             # Semantic type checking cube
│   │   └── visitor.go          # AST visitor for code generation
│   ├── symbol/
│   │   ├── constants.go        # Constants management
│   │   ├── directory.go        # Function directory
│   │   └── table.go            # Symbol table implementation
│   └── virtualMachine/
│       ├── memory.go           # Virtual machine memory
│       └── vm.go               # Virtual machine execution
├── structures/
│   ├── queue/
│   │   └── queue.go            # Queue data structure
│   └── stack/
│       └── stack.go            # Stack data structure
├── tests/                      # Example BabyDuck programs
├── go.mod                      # Go module definition
└── README.md                   # This file
```

## Compilation Process

1. **Lexical Analysis**: Source code is tokenized using ANTLR4-generated lexer
2. **Syntax Analysis**: Tokens are parsed into an Abstract Syntax Tree
3. **Symbol Table Construction**: Variables and functions are registered during parsing
4. **Semantic Analysis**: Type checking and semantic validation
5. **Intermediate Code Generation**: AST is converted to quadruples
6. **Virtual Machine Execution**: Quadruples are executed by the virtual machine

## Memory Organization

The virtual machine uses a segmented memory model:

- **Global Variables**: 1000-3999 (int: 1000-1999, float: 2000-2999, bool: 3000-3999)
- **Local Variables**: 4000-6999 (int: 4000-4999, float: 5000-5999, bool: 6000-6999)
- **Temporary Variables**: 7000-9999 (int: 7000-7999, float: 8000-8999, bool: 9000-9999)
- **Constants**: 10000-12999 (int: 10000-10999, float: 11000-11999, string: 12000-12999)

## Error Handling

The compiler provides comprehensive error detection:

- **Lexical Errors**: Invalid tokens and characters
- **Syntax Errors**: Grammar violations and malformed statements
- **Semantic Errors**: 
  - Undefined variables and functions
  - Type mismatches in operations
  - Duplicate declarations
  - Parameter count mismatches in function calls

## Dependencies

- **github.com/antlr4-go/antlr/v4**: ANTLR4 Go runtime for parser generation

## Development

### Modifying the Grammar

1. Edit `grammar/BabyDuck.g4`
2. Regenerate parser code:
```bash
antlr4 -Dlanguage=Go -o grammar/generated grammar/BabyDuck.g4
```
3. Rebuild the project

### Testing

Run example programs to test functionality:

```bash
go run cmd/compiler/main.go tests/main
go run cmd/compiler/main.go tests/factorial
go run cmd/compiler/main.go tests/fibonacci
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Submit a pull request

## License

This project is licensed under the MIT License
