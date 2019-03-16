package token

import "strconv"

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func LookupIdentifier(literal string) TokenType {
	if v, ok := keywords[literal]; ok {
		return v
	}
	return IDENT
}

func ParseDigit(literal string) TokenType {
	if literal[0] == '0' {
		return ILLEGAL
	}

	if _, err := strconv.Atoi(literal); err != nil {
		return ILLEGAL
	}

	return INT
}

var keywords = map[string]TokenType{
	"let": LET,
	"fn":  FUNCTION,
}
