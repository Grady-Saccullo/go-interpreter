package lexer

import "github.com/Grady-Saccullo/go-interpreter/token"

type Lexer struct {
	input string
	// current position in input (points to current ch)
	position int
	// current reading position in input (after reading current ch)
	readPosition int
	// current character under examination
	ch byte
}

// New create a new lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar read the next character and update underlying lexer positions
func (l *Lexer) readChar() {
	// if the next character is the last then we need to set
	// 	the readPosition to "NUL" (ASCII 0)
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	// set the current position to the previously "peaked" position
	l.position = l.readPosition
	// move the read position forward
	l.readPosition += 1
}

// NextToken gets the next token from the current lexer input and
// invokes a read of the next token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readCharacter(isLetter)
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readCharacter(isDigit)
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}

	l.readChar()
	return tok
}

type determiner func(ch byte) bool

// readCharacter uses the current input in the lexer to "yank" out the
// entire subset of characters until the determiner returns false
// and advances the current position of the Lexer internally
func (l *Lexer) readCharacter(determiner determiner) string {
	position := l.position
	for determiner(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// skipWhitespace continues to read through characters in the current input
// until a non-skippable character is found
//
// ---
//
// skippable characters include: ' ', '\t', '\n', '\r'
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// isLetter determines if a character is valid, specifically for identifiers
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit determines if a character is a valid number
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// newToken create a token for the character being currently read
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
