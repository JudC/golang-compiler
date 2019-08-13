package scanner

import "fmt"

// Kind are the types of tokens
type Kind int

const (
	ID Kind = iota
	WHITESPACE
	NUM
	NOT
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	RETURN
	IF
	ELSE
	WHILE
	PRINTLN
	WAIN
	BECOMES
	INT
	EQ
	NE
	LT
	GT
	LE
	GE
	PLUS
	MINUS
	STAR
	SLASH
	PCT
	COMMA
	SEMI
	NEW
	DELETE
	LBRACK
	RBRACK
	AMP
	NULL
)

type Token struct {
	Kind Kind
	Val  string
}

func (t Token) GetKind() Kind {
	return t.Kind
}

func (t Token) GetVal() string {
	return t.Val
}

func (t Token) ToString() string {
	kind := [...]string{
		"ID",
		"WHITESPACE",
		"NUM",
		"NOT",
		"LPAREN",
		"RPAREN",
		"LBRACE",
		"RBRACE",
		"RETURN",
		"IF",
		"ELSE",
		"WHILE",
		"PRINTLN",
		"WAIN",
		"BECOMES",
		"INT",
		"EQ",
		"NE",
		"LT",
		"GT",
		"LE",
		"GE",
		"PLUS",
		"MINUS",
		"STAR",
		"SLASH",
		"PCT",
		"COMMA",
		"SEMI",
		"NEW",
		"DELETE",
		"LBRACK",
		"RBRACK",
		"AMP",
		"NULL",
	}[t.Kind]

	return fmt.Sprintf("%v %v", kind, t.Val)
}
