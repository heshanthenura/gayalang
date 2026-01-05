package lexer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

const (
	// Special
	EOF     TokenType = "EOF"
	ILLEGAL TokenType = "ILLEGAL"

	// Identifiers + literals
	IDENT  TokenType = "IDENT"
	STRING TokenType = "STRING"
	NUMBER TokenType = "NUMBER"

	// Keywords
	REQUEST TokenType = "REQUEST"
	EXPECT  TokenType = "EXPECT"
	STATUS  TokenType = "STATUS"
	SAVE    TokenType = "SAVE"
	VAR     TokenType = "VAR"

	// HTTP Methods
	GET    TokenType = "GET"
	POST   TokenType = "POST"
	PUT    TokenType = "PUT"
	PATCH  TokenType = "PATCH"
	DELETE TokenType = "DELETE"

	// Symbols
	LBRACE  TokenType = "{"
	RBRACE  TokenType = "}"
	EQUALS  TokenType = "="
	COMMENT TokenType = "//"
)

var keywords = map[string]TokenType{
	"request": REQUEST,
	"expect":  EXPECT,
	"status":  STATUS,
	"save":    SAVE,
	"var":     VAR,

	"GET":    GET,
	"POST":   POST,
	"PUT":    PUT,
	"PATCH":  PATCH,
	"DELETE": DELETE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
