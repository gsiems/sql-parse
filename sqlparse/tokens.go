package sqlparse

/*

tokens.go provides the token list and the functionality for walking and
manipulating the token list.

*/

import (
	"strings"
)

// Tokens provides a list of tokens
type Tokens struct {
	tokens      []Token // the list of tokens
	idx         int     // the index of the current token
	length      int     // the length of the token list
	tokenIsOpen bool    // indicates if the end of the token has been reached or not
}

// Extend adds a new token to the token list and sets the type of the
// new token. If the current token is empty then the list is not
// extended and the current type is simply updated to the new type.
// If the previous token was a WhiteSpaceToken then the wihite space
// token gets replaced by the new token and the white space is associated
// with the new token
func (d *Tokens) Extend(newType int) {

	if d.Type() == WhiteSpaceToken && newType != WhiteSpaceToken {
		d.tokens[d.idx].leadingWhiteSpace = d.tokens[d.idx].tokenString
		d.tokens[d.idx].tokenString = ""
		d.tokens[d.idx].tokenType = newType

	} else {
		var nt Token
		nt.tokenType = newType
		d.tokens = append(d.tokens, nt)
		d.length = len(d.tokens)
		d.idx = d.length - 1
	}

	d.tokenIsOpen = true
}

// Push pushes a token onto the end of the token list
func (d *Tokens) Push(t Token) {
	d.tokens = append(d.tokens, t)
	d.length = len(d.tokens)
	d.idx = d.length - 1
	d.tokenIsOpen = true
}

// CloseToken flags the current token as closed.
func (d *Tokens) CloseToken() {
	d.tokenIsOpen = false
}

// SetType ensures that the current token is the new type. If, and only
// if, the current type is not the same as the new type then the token
// list is extended with the new token being set to the new type
func (d *Tokens) SetType(newType int) {
	if d.Type() != newType {
		d.Extend(newType)
	}
}

// UpdateType ensures that the current token is the new type. If the
// current type is not the same as the new type then the type is updated
// to the new type
func (d *Tokens) UpdateType(newType int) {
	if d.Type() != newType {
		if d.length > 0 {
			d.tokens[d.idx].tokenType = newType
		}
	}
}

// Type returns the type of the current token. If no such token exists
// then the NullToken value is returned.
func (d *Tokens) Type() int {
	return d.TypeN(0)
}

// TypeN returns the type of the token in the list that is distance N
// from the current token. If no such token exists then the NullToken
// value is returned.
func (d *Tokens) TypeN(n int) (t int) {
	if d.length > d.idx {
		if d.idx+n >= 0 && d.length > d.idx+n {
			return d.tokens[d.idx+n].tokenType
		}
	}
	return NullToken
}

// Concat adds the supplied string to the end of current token
func (d *Tokens) Concat(s string) {
	if d.length > d.idx {
		d.tokens[d.idx].tokenString = d.tokens[d.idx].tokenString + s
	}
}

// Next returns the current token in the list and advances the list to
// the next token.
func (d *Tokens) Next() (t Token) {
	if d.length > d.idx {
		t = d.tokens[d.idx]
		d.idx++
	}
	return t
}

// Peek returns the value of the current token in the list without
// advancing the list to the next token. If no such token exists then
// an empty string is returned.
func (d *Tokens) Peek() (s string) {
	return d.PeekN(0)
}

// PeekN returns the value of the token in the list that is distance N
// from the current token (without moving the list index). PeekN(0) is
// the same as Peek(). If no such token exists then an empty string is
// returned.
func (d *Tokens) PeekN(n int) (s string) {
	if d.length > d.idx {
		if d.idx+n >= 0 && d.length > d.idx+n {
			return d.tokens[d.idx+n].tokenString
		}
	}
	return ""
}

// Rewind resets the index of the token list to the beginning
func (d *Tokens) Rewind() {
	d.idx = 0
}

// Init initializes the token list by splitting the supplied data into
// individual characters and using that to populate the token list.
func (d *Tokens) Init(data string) {
	d.tokens = nil

	t := strings.Split(data, "")
	for i := 0; i < len(t); i++ {
		var nt Token
		nt.tokenString = t[i]
		d.tokens = append(d.tokens, nt)
	}

	d.length = len(d.tokens)
	d.idx = 0
}
