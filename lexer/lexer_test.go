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
    {
        input: "=%?;",
        tests: []testsType {
            {token.ASSIGN, "="},
            {token.ILLEGAL, "%"},
            {token.ILLEGAL, "?"},
            {token.SEMICOLON, ";"},
        },
    },
    {
        input: `let five = 5;
    let ten = 10;

    let add = fn(x, y) {
      x + y;
    };

    let result = add(five, ten);
    `,
        tests: []testsType {
            {token.LET, "let"},
            {token.IDENT, "five"},
            {token.ASSIGN, "="},
            {token.INT, "5"},
            {token.SEMICOLON, ";"},
            {token.LET, "let"},
            {token.IDENT, "ten"},
            {token.ASSIGN, "="},
            {token.INT, "10"},
            {token.SEMICOLON, ";"},
            {token.LET, "let"},
            {token.IDENT, "add"},
            {token.ASSIGN, "="},
            {token.FUNCTION, "fn"},
            {token.LPAREN, "("},
            {token.IDENT, "x"},
            {token.COMMA, ","},
            {token.IDENT, "y"},
            {token.RPAREN, ")"},
            {token.LBRACE, "{"},
            {token.IDENT, "x"},
            {token.PLUS, "+"},
            {token.IDENT, "y"},
            {token.SEMICOLON, ";"},
            {token.RBRACE, "}"},
            {token.SEMICOLON, ";"},
            {token.LET, "let"},
            {token.IDENT, "result"},
            {token.ASSIGN, "="},
            {token.IDENT, "add"},
            {token.LPAREN, "("},
            {token.IDENT, "five"},
            {token.COMMA, ","},
            {token.IDENT, "ten"},
            {token.RPAREN, ")"},
            {token.SEMICOLON, ";"},
            {token.EOF, ""},
        },
    },
    {
        input: "let five = 05",
        tests: []testsType {
            {token.LET, "let"},
            {token.IDENT, "five"},
            {token.ASSIGN, "="},
            {token.ILLEGAL, "05"},
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
