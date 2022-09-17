package main

type Lexer struct {
	input string
	pos   int
	next  int
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

	l.skipWhitespace()

	switch l.ch {
	case '(':
		t = newToken(l.ch, LPAREN)
	case ')':
		t = newToken(l.ch, RPAREN)
	case '[':
		t = newToken(l.ch, LBRACKET)
	case ']':
		t = newToken(l.ch, RBRACKET)
	case '{':
		t = newToken(l.ch, LBRACE)
	case '}':
		t = newToken(l.ch, RBRACE)
	case '.':
		t = newToken(l.ch, PERIOD)
	case ',':
		t = newToken(l.ch, COMMA)
	case '&':
		t = newToken(l.ch, AMPERSAND)
	case '`':
		t = newToken(l.ch, BACKTICK)
	case '+':
		t = newToken(l.ch, PLUS)
	case '-':
		t = newToken(l.ch, MINUS)
	case '*':
		t = newToken(l.ch, ASTERISK)
	case '/':
		t = newToken(l.ch, DIVISION)
	case '>':
		t = newToken(l.ch, GT)
	case '<':
		t = newToken(l.ch, LT)
	case '=':
		t = newToken(l.ch, EQ)
	case '~':
		t = newToken(l.ch, EQUIVALENT)
	case '!':
		if l.peekChar() == '=' {
			curr := l.ch
			l.readChar()
			literal := curr + l.ch
			t = newToken(literal, NOT_EQ)
		}
		t = newToken(l.ch, BANG)
	case 0:
		t.Literal = ""
		t.Type = EOF
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = LookupIdent(t.Literal)
			return t
		} else if isDigit(l.ch) {
			t = newToken(l.ch, INT)
			t.Literal = l.readNumber()
			return t
		} else {
			t = newToken(l.ch, ILLEGAL)
		}
	}
	l.readChar()
	return t
}

func (l *Lexer) skipWhitespace() {
	for l.ch == '\t' || l.ch == '\n' || l.ch == '\r' || l.ch == ' ' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.next >= len(l.input) {
		return 0
	} else {
		return l.input[l.next]
	}
}

func (l *Lexer) readIdentifier() string {
	first := l.pos
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[first:l.pos]
}

func (l *Lexer) readNumber() string {
	first := l.pos
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[first:l.pos]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
