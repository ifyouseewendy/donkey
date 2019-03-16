## Donkey

A language with a simple feature set following [Writing An Interpreter In Go](https://interpreterbook.com/)

## Progress

#### Lexer

Supported tokens

+ Operators: `=`, `+`, `-`, `*`, `/`, `!`, `<`, `>`, `==`, `!=`
+ Delimiters: `,`, `;`, `(`, `)`, `{`, `}`
+ Identifiers:
  - keywords: `let`, `fn`, `if`, `else`, `return`
  - literal integer, `0`, `123`
  - literal boolean, `true`, `false`
  - literal identifiers: `foo`, `bar`
+ EOF
+ Skip whitespace, ` `, `\t`, `\n`, `\r`
+ Mark errored token as Illegal
