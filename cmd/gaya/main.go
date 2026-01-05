package main

import (
	"fmt"

	"github.com/heshanthenura/gayalang/internal/lexer"
	"github.com/heshanthenura/gayalang/internal/parser"
)

func main() {
	input := `
request login {
    POST "https://api.example.com/login"
    expect status = 200
    save var token
}

request getData {
    GET "https://api.example.com/data"
    expect status = 200
}
`

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
