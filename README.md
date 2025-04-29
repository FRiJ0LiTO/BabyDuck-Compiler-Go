# BabyDuck Language Parser

Este proyecto implementa un analizador léxico y sintáctico para el lenguaje BabyDuck utilizando [ANTLR](https://www.antlr.org/) y el lenguaje de programación Go. El objetivo es procesar y analizar programas escritos en BabyDuck, verificando su estructura y generando un árbol de análisis sintáctico.

## Estructura del Proyecto

### Lexer y Parser
El archivo `parser/BabyDuck.g4` define la gramática del lenguaje BabyDuck. Esta gramática incluye las reglas para palabras clave, operadores, delimitadores, constantes, identificadores y estructuras del lenguaje como declaraciones de variables, funciones, condicionales, bucles, y más.

ANTLR genera automáticamente el lexer y el parser a partir de esta gramática. Los archivos generados se encuentran en la carpeta `parser/` e incluyen:
- `BabyDuckLexer` (Lexer)
- `BabyDuckParser` (Parser)
- `BabyDuckBaseListener` (Listener base para eventos del árbol de análisis)

### Archivo Principal
El archivo `main.go` en la raíz del proyecto es el punto de entrada del programa. Este archivo:
1. Lee un archivo de entrada que contiene código en BabyDuck.
2. Utiliza el lexer y el parser generados para analizar el código.
3. Construye un árbol de análisis sintáctico y lo imprime en la consola.

### Archivo de Prueba
El archivo `test/main.text` contiene un ejemplo de código en BabyDuck que se utiliza para probar el lexer y el parser. Este archivo incluye declaraciones de variables, funciones, condicionales, bucles y llamadas a funciones.

## Ejecución del Programa

Para ejecutar el programa, asegúrate de tener Go instalado en tu sistema. Luego, sigue estos pasos:

1. Genera el lexer y el parser utilizando ANTLR:
   ```bash
   antlr4 -Dlanguage=Go -o parser BabyDuck.g4
   ```

2. Ejecuta el programa con el siguiente comando:
   ```bash
   go run main.go test/main.text
   ```

3. El programa imprimirá el árbol de análisis sintáctico del código en `test/main.text`.

## Dependencias

Este proyecto utiliza las siguientes dependencias:
- [ANTLR4 Go Runtime](https://github.com/antlr/antlr4/tree/master/runtime/Go/antlr): Biblioteca necesaria para ejecutar el lexer y el parser generados por ANTLR.

Para instalar las dependencias, ejecuta:
```bash
go get github.com/antlr4-go/antlr/v4
```

## Estructura de Archivos

```
.
├── main.go                # Archivo principal del programa
├── parser/
│   ├── BabyDuck.g4        # Gramática del lenguaje BabyDuck
│   ├── ...                # Archivos generados por ANTLR
├── test/
│   ├── main.text          # Archivo de prueba con código BabyDuck
└── go.mod                 # Archivo de configuración de módulos de Go
```
