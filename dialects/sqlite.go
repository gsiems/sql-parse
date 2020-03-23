package dialects

import "strings"

/*
SQLite keywords

Keywords from https://www.sqlite.org/lang_keywords.html

The bool is set to false as there is no indication if the keywords are
reserved or not.

From the documentation:

"If you want to use a keyword as a name, you need to quote it. There
    are four ways of quoting keywords in SQLite:

    * 'keyword'  A keyword in single quotes is a string literal.


    * "keyword"  A keyword in double-quotes is an identifier.

    * [keyword]  A keyword enclosed in square brackets is an
        identifier. This is not standard SQL. This quoting mechanism is
        used by MS Access and SQL Server and is included in SQLite for
        compatibility.

    * `keyword`  A keyword enclosed in grave accents (ASCII code 96) is
        an identifier. This is not standard SQL. This quoting mechanism is
        used by MySQL and is included in SQLite for compatibility.

For resilience when confronted with historical SQL statements, SQLite
will sometimes bend the quoting rules above:

If a keyword in single quotes (ex: 'key' or 'glob') is used in
a context where an identifier is allowed but where a string
literal is not allowed, then the token is understood to be an
identifier instead of a string literal.

If a keyword in double quotes (ex: "key" or "glob") is used in
a context where it cannot be resolved to an identifier but
where a string literal is allowed, then the token is understood
to be a string literal instead of an identifier."

*/

// map[keyword]isReserved
var sqliteKeywords = map[string]bool{
	"ABORT":             false,
	"ACTION":            false,
	"ADD":               false,
	"AFTER":             false,
	"ALL":               false,
	"ALTER":             false,
	"ANALYZE":           false,
	"AND":               false,
	"AS":                false,
	"ASC":               false,
	"ATTACH":            false,
	"AUTOINCREMENT":     false,
	"BEFORE":            false,
	"BEGIN":             false,
	"BETWEEN":           false,
	"BY":                false,
	"CASCADE":           false,
	"CASE":              false,
	"CAST":              false,
	"CHECK":             false,
	"COLLATE":           false,
	"COLUMN":            false,
	"COMMIT":            false,
	"CONFLICT":          false,
	"CONSTRAINT":        false,
	"CREATE":            false,
	"CROSS":             false,
	"CURRENT":           false,
	"CURRENT_DATE":      false,
	"CURRENT_TIME":      false,
	"CURRENT_TIMESTAMP": false,
	"DATABASE":          false,
	"DEFAULT":           false,
	"DEFERRABLE":        false,
	"DEFERRED":          false,
	"DELETE":            false,
	"DESC":              false,
	"DETACH":            false,
	"DISTINCT":          false,
	"DO":                false,
	"DROP":              false,
	"EACH":              false,
	"ELSE":              false,
	"END":               false,
	"ESCAPE":            false,
	"EXCEPT":            false,
	"EXCLUDE":           false,
	"EXCLUSIVE":         false,
	"EXISTS":            false,
	"EXPLAIN":           false,
	"FAIL":              false,
	"FILTER":            false,
	"FIRST":             false,
	"FOLLOWING":         false,
	"FOR":               false,
	"FOREIGN":           false,
	"FROM":              false,
	"FULL":              false,
	"GLOB":              false,
	"GROUP":             false,
	"GROUPS":            false,
	"HAVING":            false,
	"IF":                false,
	"IGNORE":            false,
	"IMMEDIATE":         false,
	"IN":                false,
	"INDEX":             false,
	"INDEXED":           false,
	"INITIALLY":         false,
	"INNER":             false,
	"INSERT":            false,
	"INSTEAD":           false,
	"INTERSECT":         false,
	"INTO":              false,
	"IS":                false,
	"ISNULL":            false,
	"JOIN":              false,
	"KEY":               false,
	"LAST":              false,
	"LEFT":              false,
	"LIKE":              false,
	"LIMIT":             false,
	"MATCH":             false,
	"NATURAL":           false,
	"NO":                false,
	"NOT":               false,
	"NOTHING":           false,
	"NOTNULL":           false,
	"NULL":              false,
	"NULLS":             false,
	"OF":                false,
	"OFFSET":            false,
	"ON":                false,
	"OR":                false,
	"ORDER":             false,
	"OTHERS":            false,
	"OUTER":             false,
	"OVER":              false,
	"PARTITION":         false,
	"PLAN":              false,
	"PRAGMA":            false,
	"PRECEDING":         false,
	"PRIMARY":           false,
	"QUERY":             false,
	"RAISE":             false,
	"RANGE":             false,
	"RECURSIVE":         false,
	"REFERENCES":        false,
	"REGEXP":            false,
	"REINDEX":           false,
	"RELEASE":           false,
	"RENAME":            false,
	"REPLACE":           false,
	"RESTRICT":          false,
	"RIGHT":             false,
	"ROLLBACK":          false,
	"ROW":               false,
	"ROWS":              false,
	"SAVEPOINT":         false,
	"SELECT":            false,
	"SET":               false,
	"TABLE":             false,
	"TEMP":              false,
	"TEMPORARY":         false,
	"THEN":              false,
	"TIES":              false,
	"TO":                false,
	"TRANSACTION":       false,
	"TRIGGER":           false,
	"UNBOUNDED":         false,
	"UNION":             false,
	"UNIQUE":            false,
	"UPDATE":            false,
	"USING":             false,
	"VACUUM":            false,
	"VALUES":            false,
	"VIEW":              false,
	"VIRTUAL":           false,
	"WHEN":              false,
	"WHERE":             false,
	"WINDOW":            false,
	"WITH":              false,
	"WITHOUT":           false,
}

var sqliteOperators = map[string]bool{
	"~":  true,
	"<":  true,
	"<<": true,
	"<=": true,
	"<>": true,
	"=":  true,
	"==": true,
	">":  true,
	">=": true,
	">>": true,
	"|":  true,
	"||": true,
	"-":  true,
	"!=": true,
	"/":  true,
	"*":  true,
	"&":  true,
	"%":  true,
	"+":  true,
}

// IsSQLiteKeyword returns a boolean indicating if the supplied string
// is considered to be a keyword in SQLite
func IsSQLiteKeyword(s string) bool {

	_, ok := sqliteKeywords[strings.ToUpper(s)]
	return ok
}

// IsSQLiteReservedKeyword returns a boolean indicating if the supplied
// string is considered to be a reserved keyword in SQLite
func IsSQLiteReservedKeyword(s string) bool {

	if val, ok := sqliteKeywords[strings.ToUpper(s)]; ok {
		return val
	}
	return false
}

// IsSQLiteOperator returns a boolean indicating if the supplied string
// is considered to be an operator in SQLite
func IsSQLiteOperator(s string) bool {

	_, ok := sqliteOperators[s]
	return ok
}

// IsSQLiteLabel returns a boolean indicating if the supplied string
// is considered to be a label in SQLite
func IsSQLiteLabel(s string) bool {
	return false
}

// IsSQLiteIdentifier returns a boolean indicating if the supplied
// string is considered to be a non-quoted PostgreSQL identifier.
func IsSQLiteIdentifier(s string) bool {

	// generally unknown...
	// - cannot start with a number
	// - alpha and underscore are okay
	// - cannot contain dashes

	// Just guessing:
	const firstIdentChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	const identChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

	chr := strings.Split(s, "")
	for i := 0; i < len(chr); i++ {

		if i == 0 {
			matches := strings.Contains(firstIdentChars, chr[i])
			if !matches {
				return false
			}

		} else {
			matches := strings.Contains(identChars, chr[i])
			if !matches && chr[i] != "." {
				return false
			}

		}
	}

	return true
}
