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
	l.readChar()
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

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
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
	return ('a' <= ch && ch <= 'z') ||
		('A' <= ch && ch <= 'Z') ||
		ch == '_'
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

	return l.input[start:l.position], hasLetter
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

func (l *Lexer) readComment() string {
	start := l.position

	l.readChar()
	l.readChar()

	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}

	return l.input[start:l.position]
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	tok := Token{
		Line:   l.line,
		Column: l.column,
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
	case '/':
		if l.peekChar() == '/' {
			tok.Type = COMMENT
			tok.Literal = l.readComment()
			return tok
		}
		tok = l.newToken(ILLEGAL, "/")
	case 0:
		tok.Type = EOF
	default:
		if isLetter(l.ch) {
			lit, _ := l.readIdentifierOrNumber()
			tok.Type = LookupIdent(lit)
			tok.Literal = lit
			return tok
		}

		if isDigit(l.ch) {
			lit, hasLetter := l.readIdentifierOrNumber()
			if hasLetter {
				tok.Type = ILLEGAL
				tok.Literal = "invalid number: " + lit
				return tok
			}
			tok.Type = NUMBER
			tok.Literal = lit
			return tok
		}

		tok = l.newToken(ILLEGAL, string(l.ch))
	}

	l.readChar()
	return tok
}

func (l *Lexer) newToken(t TokenType, lit string) Token {
	return Token{
		Type:    t,
		Literal: lit,
		Line:    l.line,
		Column:  l.column,
	}
}
