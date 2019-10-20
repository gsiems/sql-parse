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
			} else if isWhiteSpace(s) {
				tl.SetType(WhiteSpaceToken)
				tl.Concat(s)

				// TODO ??
				// IdentToken
				// KeywordToken
				// LabelToken
				// NumericToken
			} else {
				// Don't know (yet) what to do with it
				tl.SetType(OtherToken)
				tl.Concat(s)
			}
		}

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
