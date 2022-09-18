package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var ast = false

	flag.BoolVar(&ast, "ast", false, "Print the FHIR PATH abstract syntax tree")
	flag.Parse()

	if ast {
		fmt.Println("ast")
	}

	if len(os.Args[1:]) == 0 {
		fmt.Println("Usage: fhirpath <query> -A")
		fmt.Println("Example: fhirpath 'Patient.name.given.exists()' -A")
		os.Exit(1)
	}
	fmt.Println(os.Args[1:])
}
