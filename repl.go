package main

import (
	"bufio"
	"fmt"
	"io"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, ">> ")

		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		l := NewLexer(line)
	}
}
