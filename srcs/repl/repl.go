package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/RealConrad/monkey-interpreter/srcs/lexer"
	"github.com/RealConrad/monkey-interpreter/srcs/token"
)

const PROMPT = "Input> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		lexer := lexer.New(line)
		for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

