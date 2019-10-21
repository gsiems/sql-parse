// Package sqlparse attempts to parse SQL, or SQL like, strings into a list of tokens.
package sqlparse

import (
	"strings"
)

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

	/*
	   BacktickQuotedToken
	   BlockCommentToken
	   BracketQuotedToken
	   DoubleQuotedToken
	   IdentToken
	   KeywordToken
	   LabelToken
	   LineCommentToken
	   NullToken
	   NumericToken
	   OperatorToken
	   OtherToken
	   PoundLineCommentToken
	   SingleQuotedToken
	   WhiteSpaceToken
	*/

	var tokenStart = map[string]int{
		"\"": DoubleQuotedToken,
		"'":  SingleQuotedToken,
		//"`":  BacktickQuotedToken,
		//"/*": BlockCommentToken     ,
		//"[":  BracketQuotedToken,
		//"--": LineCommentToken      ,
		//"#":  PoundLineCommentToken ,
	}

	var tokenEnd = map[int]string{
		BacktickQuotedToken:   "`",
		BlockCommentToken:     "*/",
		BracketQuotedToken:    "]",
		DoubleQuotedToken:     "\"",
		LineCommentToken:      "\n",
		PoundLineCommentToken: "\n",
		SingleQuotedToken:     "'",
	}

	chrs.Init(stmts)
	for {
		ch := chrs.Next()
		s := ch.Value()
		if s == "" {
			// nothing left to parse
			break
		}

		tokenType := tl.Type()
		switch tokenType {
		case BacktickQuotedToken, BracketQuotedToken, DoubleQuotedToken, SingleQuotedToken:
			tl.Concat(s)
			if te, ok := tokenEnd[tokenType]; ok {
				if s == te {
					tl.CloseToken()
				}
			}
		case LineCommentToken, PoundLineCommentToken:
			if s == "\n" {
				tl.SetType(WhiteSpaceToken)
			}
			tl.Concat(s)
		case BlockCommentToken:
			if s == "*" && chrs.Peek() == "/" {
				cn := chrs.Next()
				tl.Concat(s + cn.Value())
				tl.CloseToken()
			} else {
				// still in block comment
				tl.Concat(s)
			}

		default:
			if tt, ok := tokenStart[s]; ok {
				// Standard single char start of token
				// DoubleQuotedToken, SingleQuotedToken
				tl.Extend(tt)
				tl.Concat(s)
			} else if s == "#" && (dialect == MySQL || dialect == MariaDB) {
				tl.SetType(PoundLineCommentToken)
				tl.Concat(s)
			} else if s == "`" && (dialect == MySQL || dialect == MariaDB || dialect == SQLite) {
				// SQLite in compatibility mode
				tl.SetType(BacktickQuotedToken)
				tl.Concat(s)
			} else if s == "[" && (dialect == MSSQL || dialect == SQLite) {
				// SQLite in compatibility mode
				tl.SetType(BracketQuotedToken)
				tl.Concat(s)
			} else if s == "/" && chrs.Peek() == "*" {
				tl.SetType(BlockCommentToken)
				cn := chrs.Next()
				tl.Concat(s + cn.Value())
			} else if s == "-" && chrs.Peek() == "-" {
				tl.SetType(LineCommentToken)
				cn := chrs.Next()
				tl.Concat(s + cn.Value())
			} else if isWhiteSpaceChar(s) {
				tl.SetType(WhiteSpaceToken)
				tl.Concat(s)
			} else if s == "\\" {
				cn := chrs.Next()
				tl.Concat(s + cn.Value())
			} else if strings.Contains("(),;", s) {
				// start a new token regardless of the current state
				tl.Extend(OtherToken)
				tl.Concat(s)
				tl.CloseToken()
			} else {
				// Don't know (yet) what to do with it
				tl.SetType(OtherToken)
				tl.Concat(s)
			}
		}
	}
	return parsePassTwo(tl, dialect)
}

func parsePassTwo(tlIn Tokens, dialect int) (tlOut Tokens) {

	tlIn.Rewind()

	for {
		t := tlIn.Next()
		s := t.Value()
		if s == "" {
			// nothing left to parse
			break
		}

		tokenType := t.Type()
		switch tokenType {
		case NullToken, WhiteSpaceToken:
			// do nothing
		case BacktickQuotedToken, BlockCommentToken, BracketQuotedToken, DoubleQuotedToken, LineCommentToken, PoundLineCommentToken, SingleQuotedToken:
			tlOut.Push(t)
		default:
			// IdentToken
			// KeywordToken
			// LabelToken
			// NumericToken
			// OtherToken
			// OperatorToken
			if IsKeyword(s, dialect) {
				// KeywordToken
				tlOut.Push(t)
				tlOut.UpdateType(KeywordToken)
			} else if isNumericString(s) {
				// NumericToken
				tlOut.Push(t)
				tlOut.UpdateType(NumericToken)
			} else if isIdentString(s, dialect) {
				// IdentToken
				tlOut.Push(t)
				tlOut.UpdateType(IdentToken)
			} else {
				tlOut.Push(t)
			}
		}
	}

	return tlOut
}

// isWhiteSpaceChar determines whether or not the supplied character is
// considered to be a white space character
func isWhiteSpaceChar(s string) bool {
	const wsChars = " \n\r\t"
	return strings.Contains(wsChars, s)
}

// isNumericString determines whether or not the supplied string is
// considered to be a valid number
func isNumericString(s string) bool {
	const numChars = "0123456789."
	// "."
	// "E"
	// split upper on "E"
	// foreach check
	//   first can be +-. or digit
	//   remainder can be . or digit
	//   no more than one .

	if len(s) == 1 {
		if s == "+" || s == "-" {
			return false
		}
	}

	for _, element := range strings.Split(strings.ToUpper(s), "E") {

		if strings.Count(element, ".") > 1 {
			return false
		}

		chr := strings.Split(element, "")
		for i := 0; i < len(chr); i++ {
			matches := strings.Contains(numChars, chr[i])

			if !matches {
				if i > 0 {
					return false
				}
				if !(chr[i] == "+" || chr[i] == "-") {
					return false
				}
			}
		}
	}

	return true
}

// isIdentString determines whether or not the supplied string is a valid
// identifier for an SQL identifier of the specified dialect
func isIdentString(s string, dialect int) bool {
	const identChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789"
	const oraIdentChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789#$"
	const msIdentChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789#$@"

	// TODO: move this under dialects
	// it really isn't wuite this simple...

	chr := strings.Split(s, "")
	for i := 0; i < len(chr); i++ {

		matches := false

		if dialect == Oracle {
			// check for "starts with number, etc.?"
			matches = strings.Contains(oraIdentChars, chr[i])
		} else if dialect == MSSQL {
			matches = strings.Contains(msIdentChars, chr[i])
		} else {
			matches = strings.Contains(identChars, chr[i])
		}

		if !matches && chr[i] != "." {
			return false
		}
	}

	return true
}
