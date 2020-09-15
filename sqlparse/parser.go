// Package sqlparse attempts to parse SQL, or SQL like, strings into a list of tokens.
package sqlparse

import (
	"strconv"
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

	chrs.Init(stmts)
	for {
		ch := chrs.Next()
		s := ch.Value()
		if s == "" {
			// nothing left to parse
			break
		}

		// if we are in a *delimited* token, check for the ending
		tokenType := tl.Type()
		switch {
		case isQuotedToken(tokenType):
			tl.Concat(s)
			if isTokenEnd(s, tokenType) {
				tl.CloseToken()
			}
			continue

		case isLineCommentToken(tokenType):
			if isTokenEnd(s, tokenType) {
				tl.SetType(WhiteSpaceToken)
			}
			tl.Concat(s)
			continue

		case isBlockCommentToken(tokenType):
			if isTokenEnd(s+chrs.Peek(), tokenType) {
				cn := chrs.Next()
				tl.Concat(s + cn.Value())
				tl.CloseToken()
			} else {
				// still in block comment
				tl.Concat(s)
			}
			continue
		}

		// check for the beginning of a *delimited* token
		tt := chkTokenStart(s, chrs.Peek(), dialect)
		switch {
		case isQuotedToken(tt):
			tl.Extend(tt)
			tl.Concat(s)
			continue
		case isCommentToken(tt):
			tl.SetType(tt)
			cn := chrs.Next()
			tl.Concat(s + cn.Value())
			continue
		}

		// other
		if isWhiteSpaceChar(s) {
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
			continue
		case BacktickQuotedToken, BlockCommentToken, BracketQuotedToken, DoubleQuotedToken, LineCommentToken, SingleQuotedToken:
			tlOut.Push(t)
			continue
		case PoundLineCommentToken:
			tlOut.Push(t)
			tlOut.UpdateType(LineCommentToken)
			continue
		}

		tt := chkTokenString(s, dialect)
		switch tt {
		case KeywordToken, OperatorToken, NumericToken, IdentToken:
			tlOut.Push(t)
			tlOut.UpdateType(tt)
			continue
		}

		tlOut.Push(t)
	}

	return parsePassThree(tlOut, dialect)
}

func parsePassThree(tlIn Tokens, dialect int) (tlOut Tokens) {

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
		case OtherToken:

			switch s {
			case "(", ")", ",", ";":
				tlOut.Push(t)
				continue
			}

			// by this point all that *should* be left are:
			//  - tagging labels,
			//  - tagging bind variable placeholders,
			//  - unquoted identifiers not flagged earlier (since that check may be too simplistic), and
			//  - parsing those strings where there was no space before
			//      and/or after an operator

			if IsLabel(s, dialect) {
				tlOut.Push(t)
				tlOut.UpdateType(LabelToken)
				continue
			}

			if isBindVar(s) {
				tlOut.Push(t)
				tlOut.UpdateType(BindParameterToken)
				continue
			}

			remainder := s
			var s2 string
			ws := t.WhiteSpace()
			for {
				s2, remainder = splitOnOperator(remainder, dialect)

				if s2 != "" {
					tt := chkTokenString(s2, dialect)
					switch tt {
					case KeywordToken, OperatorToken, NumericToken, IdentToken, BindParameterToken:
						tlOut.Extend(tt)
					default:
						tlOut.Extend(OtherToken)
					}

					// leading white space
					if ws != "" {
						tlOut.SetWhiteSpace(ws)
						ws = ""
					} else {
						tlOut.SetWhiteSpace(" ")
					}

					tlOut.Concat(s2)
					tlOut.CloseToken()
				}

				if remainder == "" {
					break
				}
			}

		default:
			tlOut.Push(t)
		}
	}

	return tlOut
}

func isQuotedToken(tokenType int) bool {
	switch tokenType {
	case DoubleQuotedToken, SingleQuotedToken, BacktickQuotedToken, BracketQuotedToken:
		return true
	}
	return false
}

func isLineCommentToken(tokenType int) bool {
	switch tokenType {
	case LineCommentToken, PoundLineCommentToken:
		return true
	}
	return false
}

func isBlockCommentToken(tokenType int) bool {
	switch tokenType {
	case BlockCommentToken:
		return true
	}
	return false
}

func isCommentToken(tokenType int) bool {

	if isLineCommentToken(tokenType) {
		return true
	}
	if isBlockCommentToken(tokenType) {
		return true
	}
	return false
}

func isTokenEnd(s string, tokenType int) bool {

	var tokenEnd = map[int]string{
		BacktickQuotedToken:   "`",
		BlockCommentToken:     "*/",
		BracketQuotedToken:    "]",
		DoubleQuotedToken:     "\"",
		LineCommentToken:      "\n",
		PoundLineCommentToken: "\n",
		SingleQuotedToken:     "'",
	}

	if te, ok := tokenEnd[tokenType]; ok {
		if s == te {
			return true
		}
	}

	return false
}

func splitOnOperator(s string, dialect int) (pre, remainder string) {

	maxOperatorLen := 3
	maxLen := maxOperatorLen
	pre = s

	// search for operators starting with the longest possible operator
	if maxLen > len(s) {
		maxLen = len(s)
	}
	for i := maxLen; i > 0; i-- {
		if len(s)-i < 0 {
			continue
		}

		for j := 0; j <= len(s)-i; j++ {

			var tstOp string
			if i == 1 {
				tstOp = string(s[j])
			} else {
				tstOp = s[j : j+i]
			}

			if IsOperator(tstOp, dialect) {
				if j == 0 {
					pre = tstOp
					remainder = s[len(pre):]
				} else {
					pre = s[0:j]
					remainder = s[j:]
				}
				return
			}
		}
	}
	return
}

func chkTokenStart(s, s2 string, dialect int) (d int) {

	if s == "\"" {
		return DoubleQuotedToken
	}

	if s == "'" {
		return SingleQuotedToken
	}

	if s == "#" && (dialect == MySQL || dialect == MariaDB) {
		return PoundLineCommentToken
	}

	if s == "`" && (dialect == MySQL || dialect == MariaDB || dialect == SQLite) {
		// SQLite in compatibility mode
		return BacktickQuotedToken
	}

	if s == "[" && (dialect == MSSQL || dialect == SQLite) {
		// SQLite in compatibility mode
		return BracketQuotedToken
	}

	if s == "/" && s2 == "*" {
		return BlockCommentToken
	}

	if s == "-" && s2 == "-" {
		return LineCommentToken
	}

	return NullToken
}

func chkTokenString(s string, dialect int) (d int) {

	if IsKeyword(s, dialect) {
		return KeywordToken
	}

	if IsOperator(s, dialect) {
		return OperatorToken
	}

	if isNumericString(s) {
		return NumericToken
	}

	if IsIdentifier(s, dialect) {
		return IdentToken
	}

	if isBindVar(s) {
		return BindParameterToken
	}

	return NullToken
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

	if strings.Count(strings.ToUpper(s), "E") > 1 {
		return false
	}

	for _, element := range strings.Split(strings.ToUpper(s), "E") {
		if !isNumber(element) {
			return false
		}
	}

	return true
}

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
	/*
		const numChars = "0123456789"

		decimalPoint := "."

		if len(s) == 1 {
			if s == "+" || s == "-" {
				return false
			}
		}

		if strings.Count(s, decimalPoint) > 1 {
			return false
		}

		chr := strings.Split(s, "")
		for i := 0; i < len(chr); i++ {
			matches := strings.Contains(numChars, chr[i]) || chr[i] == decimalPoint

			if !matches {
				if i > 0 {
					return false
				}
				if !(chr[i] == "+" || chr[i] == "-") {
					return false
				}
			}
		}

		return true
	*/
}

func isBindVar(s string) bool {
	// bind variables?
	// :x
	// ?
	// $x
	// other?

	if s == "?" {
		return true
	}
	if len(s) > 1 {
		if string(s[0]) == ":" && strings.Count(s, ":") == 1 {
			return true
		}
		if string(s[0]) == "$" && strings.Count(s, "$") == 1 {
			return true
		}
	}
	return false
}
