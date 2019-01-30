package lexer

import "monkey/token"

var characters = map[byte]token.TokenType{
	'=': token.ASSIGN,
	';': token.SEMICOLON,
	'(': token.LPAREN,
	')': token.RPAREN,
	',': token.COMMA,
	'+': token.PLUS,
	'{': token.LBRACE,
	'}': token.RBRACE,
	0:   token.EOF,
}

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	tokenType := characters[l.ch]
	tok = newToken(tokenType, l.ch)

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	literal := string(ch)
	if ch == 0 {
		literal = ""
	}
	return token.Token{Type: tokenType, Literal: literal}
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
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
	l.readPosition += 1
}
