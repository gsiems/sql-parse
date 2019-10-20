// Package sqlparse attempts to parse SQL, or SQL like, strings into a list of tokens.
package sqlparse

import "strings"

/*

parser.go provides the actual parsing logic

*/

// ParseStatements takes a string of one or more SQL-ish looking
// statements and/or procedural SQL blocks and splits them into a list
// of word, symbol, comment, quoted string, etc. tokens. The dialect of
// the SQL being submitted is used to better tokenize the submitted string.
func ParseStatements(stmts string, dialect int) Tokens {
	var chrs Tokens
	var tl Tokens

	chrs.Init(stmts)
	for {
		ch := chrs.Next()
		s := ch.Value()
		if s == "" {
			// nothing left to parse
			break
		}
		// TODO: lots of stuff...
		tl.Push(ch)

	}
	return tl
}

// isWhiteSpace determines whether or not the supplied character is
// considered to be a white space character
func isWhiteSpace(s string) bool {
	const wsChars = " \n\r\t"
	return strings.Contains(wsChars, s)
}

// isNumeric determines whether or not the supplied character is
// considered to be a numeric character
func isNumeric(s string) bool {
	const numChars = "0123456789."
	return strings.Contains(numChars, s)
}

// isIdent determines whether or not the supplied character is a valid
// character for an SQL identifier of the specified dialect
func isIdent(s string, dialect int) bool {
	const identChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789"
	const oraIdentChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789#$"
	const msIdentChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789#$@"

	if dialect == Oracle {
		return strings.Contains(oraIdentChars, s)
	} else if dialect == MSSQL {
		return strings.Contains(msIdentChars, s)
	} else {
		return strings.Contains(identChars, s)
	}
}
