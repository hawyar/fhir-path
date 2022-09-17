package main

import (
	"testing"
)

func TestOnCount(t *testing.T) {
	query := "(4+5).count()"

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
	l := NewLexer(query)

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

func TestOnExists(t *testing.T) {
	query := `Patient.name.given.exists()`

	expected := []struct {
		tokenType TokenType
		literal   string
	}{
		{IDENT, "Patient"},
		{PERIOD, "."},
		{IDENT, "name"},
		{PERIOD, "."},
		{IDENT, "given"},
		{PERIOD, "."},
		{EXISTS, "exists"},
		{LPAREN, "("},
		{RPAREN, ")"},
		{EOF, ""},
	}

	l := NewLexer(query)

	for i, val := range expected {
		tok := l.NextToken()

		if tok.Type != val.tokenType {
			t.Fatalf("test[%d] - wrong tokenType expected = %q, got = %q", i, val.tokenType, tok.Type)
		}

		if tok.Literal != val.literal {
			t.Fatalf("test[%d] - wrong literal val expected = %q, got = %q", i, val.literal, tok.Literal)
		}
	}
}

func TestOnSelect(t *testing.T) {
	query := `Bundle.entry.select(resource as Patient)`

	expected := []struct {
		tokenType TokenType
		literal   string
	}{
		{IDENT, "Bundle"},
		{PERIOD, "."},
		{IDENT, "entry"},
		{PERIOD, "."},
		{SELECT, "select"},
		{LPAREN, "("},
		{IDENT, "resource"},
		{AS, "as"},
		{IDENT, "Patient"},
		{RPAREN, ")"},
		{EOF, ""},
	}

	l := NewLexer(query)

	for i, val := range expected {
		token := l.NextToken()

		if token.Type != val.tokenType {
			t.Fatalf("test[%d] - wrong tokenType expected = %q, got = %q", i, val.tokenType, token.Type)
		}

		if token.Literal != val.literal {
			t.Fatalf("test[%d] - wrong literal val expected = %q, got = %q", i, val.literal, token.Literal)
		}
	}
}
