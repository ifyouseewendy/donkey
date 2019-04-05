package repl

import (
	"bufio"
	"donkey/lexer"
	"donkey/parser"
	"fmt"
	"io"
)

// PROMPT is the prompt character for REPL
const PROMPT = ">> "

// Start is the entry for REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

// DONKEY_FACE is used in REPL for friendly compilation error
const DONKEY_FACE = `
                          /\          /\
                         ( \\        // )
                          \ \\      // /
                           \_\\||||//_/
                            \/ _  _ \
                           \/|(O)(O)|
                          \/ |      |
      ___________________\/  \      /
     //                //     |____|
    //                ||     /      \
   //|                \|     \ 0  0 /
  // \       )         V    / \____/
 //   \     /        (     /
""     \   /_________|  |_/
       /  /\   /     |  ||
      /  / /  /      \  ||
      | |  | |        | ||
      | |  | |        | ||
      |_|  |_|        |_||
       \_\  \_\        \_\\ Hard'96
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, DONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some donkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
