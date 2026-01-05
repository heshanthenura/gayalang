package parser

import (
	"fmt"
	"strconv"

	"github.com/heshanthenura/gayalang/internal/ast"
	"github.com/heshanthenura/gayalang/internal/lexer"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  lexer.Token
	peekToken lexer.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	for p.curToken.Type != lexer.EOF {
		if p.curToken.Type == lexer.REQUEST {
			req := p.parseRequest()
			program.Requests = append(program.Requests, req)
		} else {
			p.nextToken()
		}
	}
	return program
}

func (p *Parser) parseRequest() ast.RequestNode {
	req := ast.RequestNode{}

	p.nextToken()
	req.Name = p.curToken.Literal

	p.nextToken()
	if p.curToken.Type != lexer.LBRACE {
		fmt.Printf("expected { after request name, got %s\n", p.curToken.Type)
		return req
	}

	p.nextToken()
	for p.curToken.Type != lexer.RBRACE && p.curToken.Type != lexer.EOF {
		switch p.curToken.Type {
		case lexer.GET, lexer.POST, lexer.PUT, lexer.PATCH, lexer.DELETE:
			req.Method, req.URL = p.parseMethodAndURL()
		case lexer.EXPECT:
			req.Expect = p.parseExpect()
		case lexer.SAVE:
			req.SaveVar = p.parseSaveVar()
		default:
			p.nextToken()
		}
	}

	p.nextToken()
	return req
}

func (p *Parser) parseMethodAndURL() (string, string) {
	method := p.curToken.Literal
	p.nextToken()
	url := ""
	if p.curToken.Type == lexer.STRING {
		url = p.curToken.Literal
	}
	p.nextToken()
	return method, url
}

func (p *Parser) parseExpect() ast.ExpectNode {
	p.nextToken()
	if p.curToken.Type != lexer.STATUS {
		return ast.ExpectNode{}
	}
	p.nextToken()
	if p.curToken.Type != lexer.EQUALS {
		return ast.ExpectNode{}
	}
	p.nextToken()
	status, _ := strconv.Atoi(p.curToken.Literal)
	p.nextToken()
	return ast.ExpectNode{Status: status}
}

func (p *Parser) parseSaveVar() string {
	p.nextToken()
	if p.curToken.Type != lexer.VAR {
		return ""
	}
	p.nextToken()
	save := p.curToken.Literal
	p.nextToken()
	return save
}
