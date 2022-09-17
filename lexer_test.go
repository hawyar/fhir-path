package main

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	//basic := "(4+5).count()"
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
	basic := "Patient.name.where(given = 'Wouter').exists()"
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
	}

	l := NewLexer(basic)

	for i, val := range tests {
		current := l.NextToken()

		if current.Type != val.tokenType {
			t.Fatalf("test[%d] - wrong tokenType. expected = %q, got = %q", i, val.tokenType, current.Type)
		}

		if current.Literal != val.literal {
			t.Fatalf("test[%d] - wrong literal val. expected = %q, got = %q", i, val.literal, current.Literal)
		}
	}
}
