package main

import (
	"fmt"
	"os"

	"github.com/heshanthenura/gayalang/internal/executor"
	"github.com/heshanthenura/gayalang/internal/lexer"
	"github.com/heshanthenura/gayalang/internal/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gaya <filename>")
		return
	}

	filename := os.Args[1]

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	input := string(data)

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()

	ctx := executor.ExecuteProgram(program)
	fmt.Println("Execution context:", ctx)
}
