package lexer

import (
	"goInterpreter/token"
)
type Lexer struct {
	input        string
	position     int  //入力に卒kる現在の位置 (現在の文字を指し示す)
	readPosition int  //これから読み込む位置(現在の次の位置)
	ch           byte //現在検査中の文字
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 //ASCIIコードのNul文字に対応している
	} else {
		//まだ入力の最後に到達していない
		l.ch = l.input[l.readPosition] //文字列inputの次の文字を次の添字でl.chに代入
	}
	//int型次の位置を今の位置として更新
	l.position = l.readPosition
	l.readPosition++ //常に次に読もうとしている場所を指す
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '0':
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type: tokenType, //TokenType型(string)
		Literal: string(ch), //文字列型
	}
}