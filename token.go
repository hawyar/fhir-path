package main

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	DECIMAL  = "DECIMAL"
	DATE     = "DATE"
	DATETIME = "DATETIME"
	TIME     = "TIME"
	QUANTITY = "QUANTITY"

	// operators http://hl7.org/fhirpath/#operators
	PLUS       = "+"
	MINUS      = "-"
	ASTERISK   = "*"
	DIVISION   = "/"
	GT         = ">"
	LT         = "<"
	EQ         = "="
	NOT_EQ     = "!="
	EQUIVALENT = "~"
	BANG       = "!"

	// symbols http://hl7.org/fhirpath/#symbols
	LPAREN    = "("
	RPAREN    = ")"
	LBRACKET  = "["
	RBRACKET  = "]"
	LBRACE    = "{"
	RBRACE    = "}"
	PERIOD    = "."
	COMMA     = ","
	AMPERSAND = "&"
	BACKTICK  = "`"

	// keywords http://hl7.org/fhirpath/#keywords
	INDEX        = "$index"
	THIS         = "$this"
	TOTAL        = "$total"
	AND          = "and"
	AS           = "as"
	CONTAINS     = "contains"
	DAY          = "day"
	DAYS         = "days"
	DIV          = "div"
	FALSE        = "false"
	HOUR         = "hour"
	HOURS        = "hours"
	IMPLIES      = "implies"
	IN           = "in"
	IS           = "is"
	MILLISECOND  = "millisecond"
	MILLISECONDS = "milliseconds"
	MINUTE       = "minute"
	MINUTES      = "minutes"
	MOD          = "mod"
	MONTH        = "month"
	MONTHS       = "months"
	WHERE        = "where"
	OR           = "or"
	SECOND       = "second"
	SECONDS      = "seconds"
	TRUE         = "true"
	WEEK         = "week"
	WEEKS        = "weeks"
	XOR          = "xor"
	YEAR         = "year"
	YEARS        = "years"
	UNION        = "union"
	NOT          = "not"

	EXISTS = "exists"

	// todo: add functions http://hl7.org/fhirpath/#functions
	FNSELECT = "select"
	FNCOUNT  = "count"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func newToken(val byte, t TokenType) Token {
	return Token{Type: t, Literal: string(val)}
}

var Keywords = map[string]TokenType{
	"$index":   INDEX,
	"$this":    THIS,
	"$total":   TOTAL,
	"and":      AND,
	"as":       AS,
	"contains": CONTAINS,
	"count":    FNCOUNT,
	"select":   FNSELECT,
	"exists":   EXISTS,
}

func LookupIdent(ident string) TokenType {
	for k, v := range Keywords {
		if ident == k {
			return v
		}
	}
	return IDENT
}
