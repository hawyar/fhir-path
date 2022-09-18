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

	PLUS       = "+"
	MINUS      = "-"
	ASTERISK   = "*"
	DIVISION   = "/"
	GT         = ">"
	LT         = "<"
	EQ         = "="
	LT_EQ      = "<="
	GT_EQ      = ">="
	NOT_EQ     = "!="
	EQUIVALENT = "~"
	BANG       = "!"

	LPAREN      = "("
	RPAREN      = ")"
	LBRACKET    = "["
	RBRACKET    = "]"
	LBRACE      = "{"
	RBRACE      = "}"
	PERIOD      = "."
	COMMA       = ","
	AMPERSAND   = "&"
	DOLLAR_SIGN = "$"
	BACKTICK    = "`"

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
	SELECT = "select"
	COUNT  = "count"
	ALL    = "all"
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
	"count":    COUNT,
	"select":   SELECT,
	"exists":   EXISTS,
	"all":      ALL,
	"is":       IS,
}

func LookupIdent(ident string) TokenType {
	for k, v := range Keywords {
		if ident == k {
			return v
		}
	}
	return IDENT
}
