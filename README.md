# Donkey

A language with a simple feature set following [Writing An Interpreter In Go](https://interpreterbook.com/)

## Lexer

<details>
<summary>details</summary>
<p>

|            | Token      | Example    | Precedence |
|------------|------------|------------|------------|
| Identifier | INT        | 1          |            |
|            | TRUE/FALSE | true/false |            |
|            | IDENT      | foo        |            |
| Keyword    | LET        | let        |            |
|            | FUNCTION   | fn         | 6          |
|            | RETURN     | return     |            |
|            | IF         | if         |            |
|            | ELSE       | else       |            |
| Operator   | ASSIGN     | =          |            |
|            | EQ         | ==         | 1          |
|            | NOT_EQ     | !=         | 1          |
|            | LT         | <          | 2          |
|            | GT         | >          | 2          |
|            | PLUS       | +          | 3          |
|            | MINUS      | -          | i3, p5     |
|            | ASTERISK   | *          | 4          |
|            | SLASH      | /          | 4          |
|            | BANG       | !          | p5         |
| Delimiter  | COMMA      | ,          |            |
|            | SEMICOLON  | ;          |            |
|            | LPAREN     | (          |            |
|            | RPAREN     | )          |            |
|            | LBRACE     | {          |            |
|            | RBRACE     | }          |            |
| SPECIAL    | EOF        | \EOF       |            |
|            | ILLEGAL    |            |            |

* Skip whitespace, ` `, `\t`, `\n`, `\r`

</p>
</details>

## Parser

> There are two main strategies when parsing a programming language: top-down parsing or bottom-up parsing. For example, "recursive descent parsing", "Early parsing" or "predictive parsing" are all variations of top down pang. The difference between top down and bottom up parsers is that the former starts with constructing root node of the AST and then descends while the latter does it the other way around.
The parser we are going to write is a recursive descent parser. And in particular, it's a "top down operator precedence" parser, sometimes called "Pratt parser", after its inventor Vaughan Pratt. A recursive descent parser, which works from the top down, is often recommended for newcomers to parsing, since it closely mirrors the way we think about ASTs and their construction.

After lexical check, parser's job is to do syntax check, which takes the token streams from lexer and assembles an AST.

Parser is a list of statements. In the Monkey programming language, every statement besides let and return statements is an expression.

### AST

In Monkey, a program is a series of statements. Every statements besides let and return statements are expression statements.
An expression statement is not really a distinct statement; it's only a wrapper which consists solely of one expression.

```
Programm = [Statement]
Statement = LetStatement | ReturnStatement | ExpressionStatement
```

For example,

```
// let statement
let x = 5;

// return statement
return 5;

// expression statement
x + 10;
```

### Expression

+ identifiers
  - integer literal
  - boolean literal
  - identifier
+ operators
  - prefix (unary expression)
  - infix (binary expression)
  - parentheses (grouped expression)
+ if expressions
+ function expressions
  - function literal
  - function call

### Identifiers

```
5
true
foo
```

### Operators prefix

Token set: `- !`

```
-5
!true
```

### Operator infix

Token set: `+ - * / == != < >`

```
5 + 5
5 - 5
5 * 5
5 / 5
foo == bar
foo != bar
foo <  bar
foo >  bar
```

### Operator parentheses

Token set: `( )`

```
5 * (5 + 5)
((5 + 5) * 5) * 5
```

### If expression

Patten: `if (<condition>) <consequence> else <alternative>`


```
let result = if (10 > 5) { true } else { false };
```


### Function literal

Pattern: `fn <parameters> <block statement>`

```
fn(x, y) { return x + y }

// bind in a let statement
let add = fn(x, y) { return x + y }
```

### Function call

Pattern: `<expression>(<comma separated expressions>)`

```
multiply(5, add(5, 5))
```

### Parsing

Pratt Parsing, first decribed by Vaughan Pratt in the 1973 paper "Top down
operator precedence", as an alternative parsers based on CFG
(context-free grammars) and BNF (Backus-Naur-Form).

A Pratt parser's main idea is the association of parsing functions (which Pratt
calls "semantic code") with token types, instead of associating parsing functions
with grammar rules (defined in BNF)
. Whenever this token type is encountered, the parsing functions are called to
parse the appropriate expression and return an AST node that represents it.
Each token type can have up to two parsing functions associated with it,
depending on whether the token is found in a prefix or an infix position.

I think the core idea is in `parseExpression`, which always parse `prefix` operator first, then parse `infix` recursively based on precedence.

```go
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}
```

The prefix and infix parsing functions are registered during initialization. To be noted, one operator can be both prefix and infix depending on the context.

```go
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.LPAREN, p.parseCallExpression)

	p.nextToken()
	p.nextToken()

	return p
}
```
