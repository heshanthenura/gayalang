package ast

type Node interface {
	TokenLiteral() string
}

type Program struct {
	Requests []RequestNode
}

type RequestNode struct {
	Name    string
	Method  string
	URL     string
	Expect  ExpectNode
	SaveVar string
}

type ExpectNode struct {
	Status int
}

func (r RequestNode) TokenLiteral() string { return r.Name }
func (e ExpectNode) TokenLiteral() string  { return "" }
func (p Program) TokenLiteral() string     { return "" }
