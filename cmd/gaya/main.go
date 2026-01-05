package main

import (
	"fmt"
	"os"

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

	for _, req := range program.Requests {
		fmt.Println("Request:", req.Name)
		fmt.Println(" Method:", req.Method)
		fmt.Println(" URL:", req.URL)
		fmt.Println(" Expect Status:", req.Expect.Status)
		fmt.Println(" Save Var:", req.SaveVar)
		fmt.Println("---")
	}
}
