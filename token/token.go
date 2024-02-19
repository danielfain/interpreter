package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"
	STR   = "STRING"

	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	SEMICOLON = ";"
	COMMA     = ","
	PERIOD    = "."

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
