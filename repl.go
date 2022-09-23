package main

import (
	"bufio"
	"fmt"
	"io"
)

func StartRepl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, ">> ")

		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		l := NewLexer(line)

		for tok := l.NextToken(); tok.Type != EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
