package main

type Lexer struct {
	input string
	pos   int // current position
	next  int // next to current
	ch    byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.next > len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.next]
	}
	l.pos = l.next
	l.next += 1
}

func (l *Lexer) NextToken() Token {
	var t Token

	switch l.ch {
	case '(':
		t = newToken(l.ch, LPAREN)
	case ')':
		t = newToken(l.ch, RPAREN)
	case '+':
		t = newToken(l.ch, PLUS)
	case '.':
		t = newToken(l.ch, PERIOD)
	}

	l.readChar()
	return t
}

func newToken(val byte, t TokenType) Token {
	return Token{Type: t, Literal: string(val)}
}
