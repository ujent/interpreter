package lexer

import (
	"myinterpreter/token"
	)

type Lexer struct {
	input string
	position int
	readPosition int
	ch byte
}


func New(input string) *Lexer{
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) readChar(){

	if l.readPosition >= len(l.input){
		l.ch = 0
	} else{
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokType token.TokenType, ch byte) token.Token{
	return  token.Token{Type: tokType, Literal: string(ch)}
}


func (l *Lexer)NextToken() token.Token {
	var t token.Token

	switch l.ch {
	case '=':
		t = newToken(token.ASSIGN, l.ch)
	case ';':
		t = newToken(token.SEMICOLON, l.ch)
	case '(':
		t = newToken(token.LPAREN, l.ch)
	case ')':
		t = newToken(token.RPAREN, l.ch)
	case ',':
		t = newToken(token.COMMA, l.ch)
	case '+':
		t = newToken(token.PLUS, l.ch)
	case '{':
		t = newToken(token.LBRACE, l.ch)
	case '}':
		t = newToken(token.RBRACE, l.ch)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)

			return t
		}  else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()


	return t
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch){
		l.readChar()
	}

	return l.input[position:l.position]
}