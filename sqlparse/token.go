package sqlparse

/*

token.go provides the token structure and functions.

*/

import (
	"fmt"
)

const (
	// NullToken indicates an undefined, or non-existent, token
	NullToken = iota
	// WhiteSpaceToken is a string of one or more white-space characters
	WhiteSpaceToken
	// IdentToken is a string that appears to be an identifier (or SQL keyword)
	IdentToken
	// NumericToken is a string that appears to represent a numeric value
	NumericToken
	// LineCommentToken is an SQL end of line comment
	LineCommentToken
	// BlockCommentToken is an SQL block comment
	BlockCommentToken
	// SingleQuotedToken is a single quoted string
	SingleQuotedToken
	// DoubleQuotedToken is a double quoted string
	DoubleQuotedToken
	// BacktickQuotedToken is a string enclosed in back-ticks '`blah blah blah`'
	BacktickQuotedToken
	// BracketQuotedToken is a string enclosed in square brackets '[blah blah blah]`
	BracketQuotedToken
	// LabelToken is a string that indicates a PL label (for Oracle
	//  and PostgreSQL this means "enclosed in double greater that/less
	//  than symbols '<< blah_blah_blah >>'"). For MySQL, MS-SQL, and
	//  MariaDB this is an identifier followed by a colon 'blah_blah:'
	LabelToken
	// KeywordToken is a string that matches an SQL (or PL) keyword
	KeywordToken
	// OtherToken is any string not identified as any other type of token
	OtherToken
    // TODO: Add OperatorToken? Others?
)

// Token provides a single token with type information
type Token struct {
	tokenString       string // the portion of the SQL that the token contains
	tokenType         int    // the indicator as to the kind of string that the token contains
	leadingWhiteSpace string // the white space preceeding the token
}

// Value returns the string contained in the token
func (t *Token) Value() (s string) {
	return t.tokenString
}

// Type returns the type of the token
func (t *Token) Type() (i int) {
	return t.tokenType
}

// WhiteSpace returns the white space that preceeded the token
func (t *Token) WhiteSpace() (s string) {
	return t.leadingWhiteSpace
}

// TypeName returns the string representation of the token type
func (t *Token) TypeName() (s string) {
	return typeName(t.Type())
}

func typeName(t int) (s string) {

	var typeNames = map[int]string{
		BacktickQuotedToken: "BacktickQuotedToken",
		BlockCommentToken:   "BlockCommentToken",
		BracketQuotedToken:  "BracketQuotedToken",
		DoubleQuotedToken:   "DoubleQuotedToken",
		IdentToken:          "IdentToken",
		KeywordToken:        "KeywordToken",
		LabelToken:          "LabelToken",
		LineCommentToken:    "LineCommentToken",
		NullToken:           "NullToken",
		NumericToken:        "NumericToken",
		OtherToken:          "OtherToken",
		SingleQuotedToken:   "SingleQuotedToken",
		WhiteSpaceToken:     "WhiteSpaceToken",
	}

	if s, ok := typeNames[t]; ok {
		return s
	}
	return ""
}

// String implements the Stringer interface for the token
func (t Token) String() string {
	return fmt.Sprintf("%s:  [%s]", t.TypeName(), t.Value())
}
