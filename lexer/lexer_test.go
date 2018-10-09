package lexer

import (
	"testing"
	"myinterpreter/token"
	"fmt"
)

func TestNextToken1(t *testing.T){
	code := "=+(){},;"
	tests := []struct{
		expectedType token.TokenType
		expectedLiteral string
	}{
		{expectedLiteral: "=", expectedType: token.ASSIGN},
		{expectedLiteral: "+", expectedType: token.PLUS},
		{expectedLiteral: "(", expectedType: token.LPAREN},
		{expectedLiteral: ")", expectedType: token.RPAREN},
		{expectedLiteral: "{", expectedType: token.LBRACE},
		{expectedLiteral: "}", expectedType: token.RBRACE},
		{expectedLiteral: ",", expectedType: token.COMMA},
		{expectedLiteral: ";", expectedType: token.SEMICOLON},
		{expectedLiteral: "", expectedType: token.EOF},
	}

	l := New(code)

	for i, tt := range tests  {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf(("tests[%d] - wrong token type. Expected=%q, got=%q"), i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(("tests[%d] - wrong literal. Expected=%q, got=%q"), i, tt.expectedLiteral, tok.Literal)
		}
	}

	fmt.Println(tests)
}

func TestNextToken2(t *testing.T){
	code := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
`
	tests := []struct{
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{ token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{ token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(code)

	for i, tt := range tests  {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf(("tests[%d] - wrong token type. Expected=%q, got=%q"), i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(("tests[%d] - wrong literal. Expected=%q, got=%q"), i, tt.expectedLiteral, tok.Literal)
		}
	}

	fmt.Println(tests)
}
