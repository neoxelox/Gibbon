package repl

import (
	"bufio"
	"fmt"
	"io"

	"Gibbon/cfg"
	"Gibbon/lexer"
	"Gibbon/token"
)

// Repl describes a Repl instance.
// "read-eval-print loop"
type Repl struct {
	in  io.Reader
	out io.Writer
}

// New creates a Repl instance.
func New(in io.Reader, out io.Writer) *Repl {
	r := &Repl{in: in, out: out}
	return r
}

// Start inits the Repl instance.
func (r *Repl) Start() {
	scanner := bufio.NewScanner(r.in)

	for {
		fmt.Printf(cfg.PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
