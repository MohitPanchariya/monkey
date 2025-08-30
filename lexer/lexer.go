package lexer

import "github.com/MohitPanchariya/monkey/token"

type Lexer struct {
	input        string // source code
	position     int    // position of the current character
	readPosition int    // position of the next character to be read
	ch           byte   // current character under examination
}

// `New` returns a new *Lexer.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// `readChar` advances the lexer to the next character.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// 0 represents NULL in ASCII
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// `skipWhiteSpace` skips over spaces, tabs, and new lines since
// they hold significance in monkey.
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// `isLetter` returns true if the character is a letter.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// `isDigit` returns true if the character is a digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// `readIdentifier` returns a parsed identifier.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// `readNumber` returns a parsed number.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// `NextToken` returns the next token in the source code.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

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
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// The below early return is required since readIdentifier reads
		// until it encounters a non letter character. Therefore, the lexer
		// has been advanced to the start of the next token by readIdentifier
		// and the readChar prior to exiting the function is not required.
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			return newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}
