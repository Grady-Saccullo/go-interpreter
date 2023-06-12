package repl

import (
	"bufio"
	"fmt"
	"github.com/Grady-Saccullo/go-interpreter/lexer"
	"github.com/Grady-Saccullo/go-interpreter/token"
	"io"
)

// PROMPT default characters to use to detonate new prompt line
const PROMPT = ">> "

// Start the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
