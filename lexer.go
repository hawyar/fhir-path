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
	if l.next >= len(l.input) {
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
	case '.':
		t = newToken(l.ch, PERIOD)
	case 0:
		t.Literal = ""
		t.Type = EOF
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = LookupIdent(t.Literal)
		}
	}
	l.readChar()
	return t
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	//pos := l.pos
	for isLetter(l.ch) {
		l.readChar()
	}
	return ""
}

func newToken(val byte, t TokenType) Token {
	return Token{Type: t, Literal: string(val)}
}
