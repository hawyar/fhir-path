package main

import (
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
	//more := "(4+5).count()"
	//tests := []struct {
	//	tokenType TokenType
	//	literal   string
	//}{
	//	{LPAREN, "("},
	//	{INT, "4"},
	//	{PLUS, "+"},
	//	{INT, "5"},
	//	{RPAREN, ")"},
	//	{PERIOD, "."},
	//	{COUNT, "count"},
	//	{LPAREN, "("},
	//	{RPAREN, ")"},
	//}

	basic := `().`
	tests := []struct {
		tokenType TokenType
		literal   string
	}{
		{LPAREN, "("},
		{RPAREN, ")"},
		{PERIOD, "."},
	}

	l := NewLexer(basic)

	for i, val := range tests {
		current := l.NextToken()

		if current.Type != val.tokenType {
			t.Fatalf("tests[%d] - wrong tokenType. expected = %q, got = %q", i, val.tokenType, current.Type)
		}

		if current.Literal != val.literal {
			t.Fatalf("tests[%d] - wrong literal val. expected = %q, got = %q", i, val.literal, current.Literal)
		}
	}

	fmt.Println(l)
}
