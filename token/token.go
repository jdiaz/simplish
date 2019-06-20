package token

type TokenType = string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers
	IDENT = "IDENT" // add, foobar, x , y etc
	INT   = "INT"   // 123456

	// Operators
	IS  = "IS"
	NOT = "NOT"
	AND = "AND"
	OR  = "OR"
	LT  = "LESS THAN"
	LE  = "LESS/EQUALS"
	GT  = "GREATER THAN"
	GE  = "GREATER/EQUALS"

	EQ = "EQUALS"

	LPAREN = "("
	RPAREN = ")"

	// Keywords
	TRUE  = "TRUE"
	FALSE = "FALSE"
)

var keywords = map[string]TokenType{
	"true":  TRUE,
	"false": FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
