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
		//"`":  BacktickQuotedToken,   // dialect dependent
		//"[":  BracketQuotedToken,    // dialect dependent
		//"#":  PoundLineCommentToken, // dialect dependent
		//"/*": BlockCommentToken,     // multi-char
		//"--": LineCommentToken,      // multi-char
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
			} else if isOperatorChar(s) {
				tl.SetType(OperatorToken)
				tl.Concat(s)
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
		case BacktickQuotedToken, BlockCommentToken, BracketQuotedToken, DoubleQuotedToken, LineCommentToken, SingleQuotedToken:
			tlOut.Push(t)
		case PoundLineCommentToken:
			tlOut.Push(t)
			tlOut.UpdateType(LineCommentToken)
		default:
			// IdentToken
			// KeywordToken
			// LabelToken
			// NumericToken
			// OtherToken
			if IsKeyword(s, dialect) {
				// KeywordToken
				tlOut.Push(t)
				tlOut.UpdateType(KeywordToken)
			} else if isNumericString(s) {
				// By operator tokens previously and having queries
				// with minimal whitespace the numbers, especially those
				// in scientific notation are potentially split over
				// several tokens.
				//
				// Unsigned numbers should be fine, including those in
				// scientific notation.
				//
				// Signed numbers, and those in scientific notation where
				// the exponent is signed, will need to be consolidated.
				//
				// If the current numeric ends in an "E" and the next
				// token is either "+" or "-" and the token after that is
				// a number then join them.
				var tmp string
				if strings.HasSuffix(strings.ToUpper(s), "E") {
					if (tlIn.Peek() == "+" || tlIn.Peek() == "-") && tlIn.WhiteSpace() == "" {
						if isNumericString(tlIn.PeekN(1)) && tlIn.WhiteSpaceN(1) == "" {
							tn := tlIn.Next()
							tmp = tn.Value()
							tn = tlIn.Next()
							tmp = tmp + tn.Value()
						}
					}
				}

				// If the previous token was a sign ("+" or "-") and the
				// token prior to that was a keyword, comma, operator, or
				// open paren then it is very unlikely that the signed
				// token isn't part of the number (as opposed to being an
				// arithmetic operation).
				prevVal := tlOut.PeekN(-1)
				prevType := tlOut.TypeN(-1)

				if (prevType == KeywordToken || prevType == OperatorToken || prevVal == ",") && (tlOut.Peek() == "+" || tlOut.Peek() == "-") && t.WhiteSpace() == "" {
					// Previous is part of NumericToken
					tlOut.UpdateType(NumericToken)
					tlOut.Concat(s + tmp)
				} else {

					// Current is NumericToken
					tlOut.Push(t)
					tlOut.UpdateType(NumericToken)
					tlOut.Concat(tmp)

				}
			} else if IsIdentifier(s, dialect) {
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

// isOperatorChar determines whether or not the supplied character is
// considered to be an operator character
func isOperatorChar(s string) bool {
	// TODO: this may also need to move to dialects

	const opChars = "^~<=>|-!/@*&#%+"
	return strings.Contains(opChars, s)
}

// isNumericString determines whether or not the supplied string is
// considered to be a valid number (or portion thereof)
func isNumericString(s string) bool {
	const numChars = "0123456789."
	// "."
	// "E"
	// split upper on "E"
	// foreach check
	//   first can be "+", "-", ".", or digit
	//   remainder can be "." or digit
	//   no more than one "." or "E"

	if len(s) == 1 {
		if s == "+" || s == "-" {
			return false
		}
	}

	if strings.Count(strings.ToUpper(s), "E") > 1 {
		return false
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
