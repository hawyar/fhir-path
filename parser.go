package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("query is required")
	}

	raw := os.Args[1]

	fmt.Printf("input: %s \n", raw)
}
