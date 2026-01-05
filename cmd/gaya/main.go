package main

import (
	"fmt"

	"github.com/heshanthenura/gayalang/internal/lexer"
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

	lex := lexer.New(input)

	fmt.Println("Tokens:")
	for tok := lex.NextToken(); tok.Type != lexer.EOF; tok = lex.NextToken() {
		fmt.Printf("[%d:%d] %-10s : %q\n", tok.Line, tok.Column, tok.Type, tok.Literal)
	}
	fmt.Println("EOF reached")
}
