package scanner

type State int

const (
	S_ID State = iota
	S_LABEL
	S_COMMA
	S_LPAREN
	S_RPAREN
	S_INT
	S_HEXINT
	S_REG
	S_WHITESPACE
	S_COMMENT

	// States that are not also kinds
	S_FAIL
	S_START
	S_DOT
	S_DOTID
	S_ZERO
	S_ZEROX
	S_MINUS
	S_DOLLARS
)

type Scanner struct {
}
