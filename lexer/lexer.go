package lexer

import (
	"interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.char {
	case '=':
		tok = newToken(token.ASSIGN, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '-':
		tok = newToken(token.MINUS, l.char)
	case '/':
		tok = newToken(token.DIVIDE, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case 0:
		tok = newToken(token.EOF, l.char)
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			literal, tokenType := l.readNumber()
			tok.Literal = literal
			tok.Type = tokenType
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	firstPosition := l.position

	for isLetter(l.char) {
		l.readChar()
	}

	return l.input[firstPosition:l.position]
}

func (l *Lexer) readNumber() (string, token.TokenType) {
	firstPosition := l.position

	for isDigit(l.char) {
		l.readChar()
	}

	if l.char != '.' {
		return l.input[firstPosition:l.position], token.INT
	}

	l.readChar()

	for isDigit(l.char) {
		l.readChar()
	}

	return l.input[firstPosition:l.position], token.FLOAT
}

func numberType(numberLiteral string) token.TokenType {
	for i := range len(numberLiteral) {
		if numberLiteral[i] == '.' {
			return token.FLOAT
		}
	}
	return token.INT
}

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
