package lexer

import (
    "testing"
    "donkey/token"
)

type testsType struct {
    expectedType    token.TokenType
    expectedLiteral string
}

var testCases = []struct {
    input string
    tests []testsType
} {
    {
        input: "=+(){},;",
        tests: []testsType {
            {token.ASSIGN, "="},
            {token.PLUS, "+"},
            {token.LPAREN, "("},
            {token.RPAREN, ")"},
            {token.LBRACE, "{"},
            {token.RBRACE, "}"},
            {token.COMMA, ","},
            {token.SEMICOLON, ";"},
            {token.EOF, ""},
        },
    },
}

func TestNextToken(t *testing.T) {
    for _, testCase := range testCases {
        input := testCase.input
        tests := testCase.tests

        l := New(input)

        for i, tt := range tests {
            tok := l.NextToken()

            if tok.Type != tt.expectedType {
                t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
                    i, tt.expectedType, tok.Type)
            }

            if tok.Literal != tt.expectedLiteral {
                t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
                    i, tt.expectedLiteral, tok.Literal)
            }
        }
    }
}
