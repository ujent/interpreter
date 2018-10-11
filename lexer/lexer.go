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


func (l *Lexer)NextToken() token.Token {
	var t token.Token
	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			lit := string(ch) + string(l.ch)
			t = token.Token{Type:token.EQ, Literal: lit}
		} else {
			t = newToken(token.ASSIGN, l.ch)
		}
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
	case '-':
		t = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			lit := string(ch) + string(l.ch)
			t = token.Token{Type:token.NOT_EQ, Literal: lit}
		} else {
			t = newToken(token.BANG, l.ch)
		}
	case '*':
		t = newToken(token.ASTERIKS, l.ch)
	case '/':
		t = newToken(token.SLASH, l.ch)
	case '<':
		t = newToken(token.LT, l.ch)
	case '>':
		t = newToken(token.GT, l.ch)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)

			return t
		}  else if isDigit(l.ch){
			t.Type = token.INT
			t.Literal = l.readNumber()

			return t
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()


	return t
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch){
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace(){
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	pos := l.position

	for isDigit(l.ch){
		l.readChar()
	}

	return l.input[pos:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input){
		return 0
	} else{
		return l.input[l.readPosition]
	}
}

func newToken(tokType token.TokenType, ch byte) token.Token{
	return  token.Token{Type: tokType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}


