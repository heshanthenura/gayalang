package lexer

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	line         int
	column       int
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
		line:  1,
	}
	l.readChar() // set first character
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	if l.ch != 0 {
		l.column++
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line++
			l.column = 0
		}
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifierOrNumber() (string, bool) {
	start := l.position
	hasLetter := false

	for isLetter(l.ch) || isDigit(l.ch) {
		if isLetter(l.ch) {
			hasLetter = true
		}
		l.readChar()
	}

	lit := l.input[start:l.position]
	return lit, hasLetter
}

func (l *Lexer) readString() string {
	l.readChar()
	start := l.position
	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}
	value := l.input[start:l.position]
	l.readChar()
	return value
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	startLine := l.line
	startCol := l.column
	tok := Token{
		Line:   startLine,
		Column: startCol,
	}

	switch l.ch {
	case '{':
		tok = l.newToken(LBRACE, "{")
	case '}':
		tok = l.newToken(RBRACE, "}")
	case '=':
		tok = l.newToken(EQUALS, "=")
	case '"':
		tok.Type = STRING
		tok.Literal = l.readString()
		return tok
	case 0:
		tok.Type = EOF
	default:
		if isLetter(l.ch) {
			lit, _ := l.readIdentifierOrNumber()
			tok.Type = LookupIdent(lit)
			tok.Literal = lit
			return tok
		} else if isDigit(l.ch) {
			lit, hasLetter := l.readIdentifierOrNumber()
			if hasLetter {
				tok.Type = ILLEGAL
				tok.Literal = "invalid number: " + lit
				return tok
			}
			tok.Type = NUMBER
			tok.Literal = lit
			return tok
		} else {
			tok = l.newToken(ILLEGAL, string(l.ch))
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) newToken(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
		Line:    l.line,
		Column:  l.column,
	}
}
