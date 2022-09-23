package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var input = ""

	flag.StringVar(&input, "i", "", "The FHIR path query")
	flag.Parse()

	if len(os.Args[1:]) == 0 {
		fmt.Println("starting repl")
		StartRepl(os.Stdin, os.Stdout)
	}

	l := NewLexer(input)

	fmt.Println("tokens: ")
	for token := l.NextToken(); token.Type != EOF; token = l.NextToken() {
		fmt.Printf("%+v \n", token)
	}
	os.Exit(1)
}
