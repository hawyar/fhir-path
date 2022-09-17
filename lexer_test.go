package main

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	basic := "(4+5).count()"

	tests := []struct {
		tokenType TokenType
		literal   string
	}{
		{LPAREN, "("},
		{INT, "4"},
		{PLUS, "+"},
		{INT, "5"},
		{RPAREN, ")"},
		{PERIOD, "."},
		{COUNT, "count"},
		{LPAREN, "("},
		{RPAREN, ")"},
		{EOF, ""},
	}
	l := NewLexer(basic)

	for i, val := range tests {
		tok := l.NextToken()

		if tok.Type != val.tokenType {
			t.Fatalf("test[%d] - wrong tokenType. expected = %q, got = %q", i, val.tokenType, tok.Type)
		}

		if tok.Literal != val.literal {
			t.Fatalf("test[%d] - wrong literal val. expected = %q, got = %q", i, val.literal, tok.Literal)
		}
	}
}
