package lexer

import (
	"log"
	"testing"
)

func TestMain(t *testing.T) {
	input := `
request login {
    POST "https://webhook.site/ec2b106b-a4c6-4620-9873-14cf4af6fe20"
    expect status = 200
}

// testing
request getData {
    GET "https://judge0-be.vercel.app/api/indexasdad" 
    expect status = 200
	sdfsdfsdfd
}
`

	l := New(input)

	for {
		tok := l.NextToken()
		log.Printf("%d %d %s %s", tok.Line, tok.Column, tok.Type, tok.Literal)
		if tok.Type == EOF {
			break
		}
	}
}
