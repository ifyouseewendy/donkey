package token

import "strconv"

const (
	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"
	TRUE  = "TRUE"
	FALSE = "FALSE"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	SLASH    = "/"
	ASTERISK = "*"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"
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
	"let":    LET,
	"fn":     FUNCTION,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}
