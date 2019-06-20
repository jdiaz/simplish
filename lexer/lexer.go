package lexer

import (
	"fmt"
	"simplish/token"
	"strings"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipBlankspace()

	// var peeked string
	// var peekedUpperCase string

	switch l.ch {
	case 'i':
		peeked := l.peekMultiWordOperator(1)
		if strings.ToUpper(peeked) == token.IS {
			n := len(token.IS)
			for i := 0; i < n; i++ {
				l.readChar()
			}
			tok = token.Token{Type: token.IS, Literal: peeked}
		}
	case 'g':
		peeked := l.peekMultiWordOperator(2)
		peekedUpperCase := strings.ToUpper(peeked)
		if peekedUpperCase == token.GT {
			n := len(token.GT)
			for i := 0; i < n; i++ {
				l.readChar()
			}
			tok = token.Token{Type: token.GT, Literal: peeked}
			break
		}
		peeked = l.peekMultiWordOperator(1)
		peekedUpperCase = strings.ToUpper(peeked)
		if peekedUpperCase == token.GE {
			n := len(token.GE)
			for i := 0; i < n; i++ {
				l.readChar()
			}
			tok = token.Token{Type: token.GE, Literal: peeked}
			break
		}
	case 'l':
		peeked := l.peekMultiWordOperator(2)
		peekedUpperCase := strings.ToUpper(peeked)
		if peekedUpperCase == token.LT {
			n := len(token.LT)
			for i := 0; i < n; i++ {
				l.readChar()
			}
			tok = token.Token{Type: token.LT, Literal: peeked}
			break
		}
		peeked = l.peekMultiWordOperator(1)
		peekedUpperCase = strings.ToUpper(peeked)
		if peekedUpperCase == token.LE {
			n := len(token.LE)
			for i := 0; i < n; i++ {
				l.readChar()
			}
			tok = token.Token{Type: token.LE, Literal: peeked}
			break
		}
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) skipBlankspace() {
	for isWordDelimiter(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// Peeks ahead in the input text in attempt to identify a multiword operator
// such as: "less than". The number of whitespaceAllow will delimit the last
// whitespace character to act as a delimiter i.e.
// peekMultiWordOperator(1) will peek forward until it finds a single whitespace
// character.
// peekMultiWordOperator(2) will peek forward until it finds a second ocurrance
// of the whitespace character
func (l *Lexer) peekMultiWordOperator(whitespaceAllowed int) string {
	if l.readPosition >= len(l.input) {
		return ""
	}
	nextCharPos := l.readPosition
	whitespaceCount := 0
	for whitespaceCount < whitespaceAllowed && nextCharPos != len(l.input) {
		fmt.Printf("\t%s\n", string(l.input[nextCharPos]))
		if isWordDelimiter(l.input[nextCharPos]) {
			whitespaceCount++
		}
		nextCharPos++
	}
	// Operator could be delimited by end of the string or a whitespace
	// In cases where whitespace bounds the operator nextCharPos points to the
	// location of the whitespace. However, if its the end the string nextCharPos
	// points to the last character of the operator.
	if whitespaceCount == whitespaceAllowed {
		return l.input[l.position : nextCharPos-1]
	}
	return l.input[l.position:nextCharPos]

}

func isWordDelimiter(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
