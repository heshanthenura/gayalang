package lexer

import (
	"testing"
)

func TestInvalidIdentifier(t *testing.T) {
	input := `12asd12 123bds 45valid 78`

	expected := []struct {
		typ     TokenType
		literal string
	}{
		{ILLEGAL, "invalid number: 12asd12"},
		{ILLEGAL, "invalid number: 123bds"},
		{ILLEGAL, "invalid number: 45valid"},
		{NUMBER, "78"},
		{EOF, ""},
	}

	l := New(input)

	for i, exp := range expected {
		tok := l.NextToken()

		if tok.Type != exp.typ {
			t.Fatalf("test[%d] - wrong token type. got=%q, want=%q", i, tok.Type, exp.typ)
		}
		if tok.Literal != exp.literal {
			t.Fatalf("test[%d] - wrong literal. got=%q, want=%q", i, tok.Literal, exp.literal)
		}
	}
}
