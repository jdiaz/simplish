package lexer

import (
	"simplish/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := ` is 
	
	greater than 
		
age
	less than

	 ( true false 


		name
	greater/equals )

	21 
	
	less/equals`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IS, "is"},
		{token.GT, "greater than"},
		{token.IDENT, "age"},
		{token.LT, "less than"},
		{token.LPAREN, "("},
		{token.TRUE, "true"},
		{token.FALSE, "false"},
		{token.IDENT, "name"},
		{token.GE, "greater/equals"},
		{token.RPAREN, ")"},
		{token.INT, "21"},
		{token.LE, "less/equals"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, tt := range tests {
		token := lexer.NextToken()

		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected =%q, got=%q",
				i, tt.expectedType, token.Type)
		}

		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, token.Literal)
		}
	}
}
