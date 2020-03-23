package dialects

import "strings"

/*
PostgreSQL keywords

https://www.postgresql.org/docs/current/sql-keywords-appendix.html

*/

// map[keyword]isReserved
var pgKeywords = map[string]bool{
	"ABORT":             false,
	"ABSOLUTE":          false,
	"ACCESS":            false,
	"ACTION":            false,
	"ADD":               false,
	"ADMIN":             false,
	"AFTER":             false,
	"AGGREGATE":         false,
	"ALL":               true,
	"ALSO":              false,
	"ALTER":             false,
	"ALWAYS":            false,
	"ANALYSE":           true,
	"ANALYZE":           true,
	"AND":               true,
	"ANY":               true,
	"ARRAY":             true,
	"ASC":               true,
	"ASSERTION":         false,
	"ASSIGNMENT":        false,
	"AS":                true,
	"ASYMMETRIC":        true,
	"AT":                false,
	"ATTACH":            false,
	"ATTRIBUTE":         false,
	"AUTHORIZATION":     true,
	"BACKWARD":          false,
	"BEFORE":            false,
	"BEGIN":             false,
	"BETWEEN":           false,
	"BIGINT":            false,
	"BINARY":            true,
	"BIT":               false,
	"BOOLEAN":           false,
	"BOTH":              true,
	"BY":                false,
	"CACHE":             false,
	"CALLED":            false,
	"CALL":              false,
	"CASCADED":          false,
	"CASCADE":           false,
	"CASE":              true,
	"CAST":              true,
	"CATALOG":           false,
	"CHAIN":             false,
	"CHARACTER":         false,
	"CHARACTERISTICS":   false,
	"CHAR":              false,
	"CHECKPOINT":        false,
	"CHECK":             true,
	"CLASS":             false,
	"CLOSE":             false,
	"CLUSTER":           false,
	"COALESCE":          false,
	"COLLATE":           true,
	"COLLATION":         true,
	"COLUMNS":           false,
	"COLUMN":            true,
	"COMMENT":           false,
	"COMMENTS":          false,
	"COMMIT":            false,
	"COMMITTED":         false,
	"CONCURRENTLY":      true,
	"CONFIGURATION":     false,
	"CONFLICT":          false,
	"CONNECTION":        false,
	"CONSTRAINTS":       false,
	"CONSTRAINT":        true,
	"CONTENT":           false,
	"CONTINUE":          false,
	"CONVERSION":        false,
	"COPY":              false,
	"COST":              false,
	"CREATE":            true,
	"CROSS":             true,
	"CSV":               false,
	"CUBE":              false,
	"CURRENT_CATALOG":   true,
	"CURRENT_DATE":      true,
	"CURRENT":           false,
	"CURRENT_ROLE":      true,
	"CURRENT_SCHEMA":    true,
	"CURRENT_TIMESTAMP": true,
	"CURRENT_TIME":      true,
	"CURRENT_USER":      true,
	"CURSOR":            false,
	"CYCLE":             false,
	"DATABASE":          false,
	"DATA":              false,
	"DAY":               false,
	"DEALLOCATE":        false,
	"DEC":               false,
	"DECIMAL":           false,
	"DECLARE":           false,
	"DEFAULTS":          false,
	"DEFAULT":           true,
	"DEFERRABLE":        true,
	"DEFERRED":          false,
	"DEFINER":           false,
	"DELETE":            false,
	"DELIMITER":         false,
	"DELIMITERS":        false,
	"DEPENDS":           false,
	"DESC":              true,
	"DETACH":            false,
	"DICTIONARY":        false,
	"DISABLE":           false,
	"DISCARD":           false,
	"DISTINCT":          true,
	"DOCUMENT":          false,
	"DOMAIN":            false,
	"DO":                true,
	"DOUBLE":            false,
	"DROP":              false,
	"EACH":              false,
	"ELSE":              true,
	"ENABLE":            false,
	"ENCODING":          false,
	"ENCRYPTED":         false,
	"END":               true,
	"ENUM":              false,
	"ESCAPE":            false,
	"EVENT":             false,
	"EXCEPT":            true,
	"EXCLUDE":           false,
	"EXCLUDING":         false,
	"EXCLUSIVE":         false,
	"EXECUTE":           false,
	"EXISTS":            false,
	"EXPLAIN":           false,
	"EXTENSION":         false,
	"EXTERNAL":          false,
	"EXTRACT":           false,
	"FALSE":             true,
	"FAMILY":            false,
	"FETCH":             true,
	"FILTER":            false,
	"FIRST":             false,
	"FLOAT":             false,
	"FOLLOWING":         false,
	"FORCE":             false,
	"FOREIGN":           true,
	"FOR":               true,
	"FORWARD":           false,
	"FREEZE":            true,
	"FROM":              true,
	"FULL":              true,
	"FUNCTION":          false,
	"FUNCTIONS":         false,
	"GENERATED":         false,
	"GLOBAL":            false,
	"GRANTED":           false,
	"GRANT":             true,
	"GREATEST":          false,
	"GROUPING":          false,
	"GROUPS":            false,
	"GROUP":             true,
	"HANDLER":           false,
	"HAVING":            true,
	"HEADER":            false,
	"HOLD":              false,
	"HOUR":              false,
	"IDENTITY":          false,
	"IF":                false,
	"ILIKE":             true,
	"IMMEDIATE":         false,
	"IMMUTABLE":         false,
	"IMPLICIT":          false,
	"IMPORT":            false,
	"INCLUDE":           false,
	"INCLUDING":         false,
	"INCREMENT":         false,
	"INDEXES":           false,
	"INDEX":             false,
	"INHERIT":           false,
	"INHERITS":          false,
	"INITIALLY":         true,
	"INLINE":            false,
	"INNER":             true,
	"INOUT":             false,
	"INPUT":             false,
	"INSENSITIVE":       false,
	"INSERT":            false,
	"INSTEAD":           false,
	"INTEGER":           false,
	"INTERSECT":         true,
	"INTERVAL":          false,
	"INT":               false,
	"INTO":              true,
	"IN":                true,
	"INVOKER":           false,
	"ISNULL":            true,
	"ISOLATION":         false,
	"IS":                true,
	"JOIN":              true,
	"KEY":               false,
	"LABEL":             false,
	"LANGUAGE":          false,
	"LARGE":             false,
	"LAST":              false,
	"LATERAL":           true,
	"LEADING":           true,
	"LEAKPROOF":         false,
	"LEAST":             false,
	"LEFT":              true,
	"LEVEL":             false,
	"LIKE":              true,
	"LIMIT":             true,
	"LISTEN":            false,
	"LOAD":              false,
	"LOCAL":             false,
	"LOCALTIMESTAMP":    true,
	"LOCALTIME":         true,
	"LOCATION":          false,
	"LOCKED":            false,
	"LOCK":              false,
	"LOGGED":            false,
	"MAPPING":           false,
	"MATCH":             false,
	"MATERIALIZED":      false,
	"MAXVALUE":          false,
	"METHOD":            false,
	"MINUTE":            false,
	"MINVALUE":          false,
	"MODE":              false,
	"MONTH":             false,
	"MOVE":              false,
	"NAME":              false,
	"NAMES":             false,
	"NATIONAL":          false,
	"NATURAL":           true,
	"NCHAR":             false,
	"NEW":               false,
	"NEXT":              false,
	"NO":                false,
	"NONE":              false,
	"NOTHING":           false,
	"NOTIFY":            false,
	"NOTNULL":           true,
	"NOT":               true,
	"NOWAIT":            false,
	"NULLIF":            false,
	"NULLS":             false,
	"NULL":              true,
	"NUMERIC":           false,
	"OBJECT":            false,
	"OF":                false,
	"OFF":               false,
	"OFFSET":            true,
	"OIDS":              false,
	"OLD":               false,
	"ONLY":              true,
	"ON":                true,
	"OPERATOR":          false,
	"OPTION":            false,
	"OPTIONS":           false,
	"ORDER":             true,
	"ORDINALITY":        false,
	"OR":                true,
	"OTHERS":            false,
	"OUTER":             true,
	"OUT":               false,
	"OVER":              false,
	"OVERLAPS":          true,
	"OVERLAY":           false,
	"OVERRIDING":        false,
	"OWNED":             false,
	"OWNER":             false,
	"PARALLEL":          false,
	"PARSER":            false,
	"PARTIAL":           false,
	"PARTITION":         false,
	"PASSING":           false,
	"PASSWORD":          false,
	"PLACING":           true,
	"PLANS":             false,
	"POLICY":            false,
	"POSITION":          false,
	"PRECEDING":         false,
	"PRECISION":         false,
	"PREPARED":          false,
	"PREPARE":           false,
	"PRESERVE":          false,
	"PRIMARY":           true,
	"PRIOR":             false,
	"PRIVILEGES":        false,
	"PROCEDURAL":        false,
	"PROCEDURE":         false,
	"PROCEDURES":        false,
	"PROGRAM":           false,
	"PUBLICATION":       false,
	"QUOTE":             false,
	"RANGE":             false,
	"READ":              false,
	"REAL":              false,
	"REASSIGN":          false,
	"RECHECK":           false,
	"RECURSIVE":         false,
	"REFERENCES":        true,
	"REFERENCING":       false,
	"REF":               false,
	"REFRESH":           false,
	"REINDEX":           false,
	"RELATIVE":          false,
	"RELEASE":           false,
	"RENAME":            false,
	"REPEATABLE":        false,
	"REPLACE":           false,
	"REPLICA":           false,
	"RESET":             false,
	"RESTART":           false,
	"RESTRICT":          false,
	"RETURNING":         true,
	"RETURNS":           false,
	"REVOKE":            false,
	"RIGHT":             true,
	"ROLE":              false,
	"ROLLBACK":          false,
	"ROLLUP":            false,
	"ROUTINE":           false,
	"ROUTINES":          false,
	"ROW":               false,
	"ROWS":              false,
	"RULE":              false,
	"SAVEPOINT":         false,
	"SCHEMA":            false,
	"SCHEMAS":           false,
	"SCROLL":            false,
	"SEARCH":            false,
	"SECOND":            false,
	"SECURITY":          false,
	"SELECT":            true,
	"SEQUENCE":          false,
	"SEQUENCES":         false,
	"SERIALIZABLE":      false,
	"SERVER":            false,
	"SESSION":           false,
	"SESSION_USER":      true,
	"SET":               false,
	"SETOF":             false,
	"SETS":              false,
	"SHARE":             false,
	"SHOW":              false,
	"SIMILAR":           true,
	"SIMPLE":            false,
	"SKIP":              false,
	"SMALLINT":          false,
	"SNAPSHOT":          false,
	"SOME":              true,
	"SQL":               false,
	"STABLE":            false,
	"STANDALONE":        false,
	"START":             false,
	"STATEMENT":         false,
	"STATISTICS":        false,
	"STDIN":             false,
	"STDOUT":            false,
	"STORAGE":           false,
	"STORED":            false,
	"STRICT":            false,
	"STRIP":             false,
	"SUBSCRIPTION":      false,
	"SUBSTRING":         false,
	"SUPPORT":           false,
	"SYMMETRIC":         true,
	"SYSID":             false,
	"SYSTEM":            false,
	"TABLESAMPLE":       true,
	"TABLES":            false,
	"TABLESPACE":        false,
	"TABLE":             true,
	"TEMP":              false,
	"TEMPLATE":          false,
	"TEMPORARY":         false,
	"TEXT":              false,
	"THEN":              true,
	"TIES":              false,
	"TIME":              false,
	"TIMESTAMP":         false,
	"TO":                true,
	"TRAILING":          true,
	"TRANSACTION":       false,
	"TRANSFORM":         false,
	"TREAT":             false,
	"TRIGGER":           false,
	"TRIM":              false,
	"TRUE":              true,
	"TRUNCATE":          false,
	"TRUSTED":           false,
	"TYPE":              false,
	"TYPES":             false,
	"UNBOUNDED":         false,
	"UNCOMMITTED":       false,
	"UNENCRYPTED":       false,
	"UNION":             true,
	"UNIQUE":            true,
	"UNKNOWN":           false,
	"UNLISTEN":          false,
	"UNLOGGED":          false,
	"UNTIL":             false,
	"UPDATE":            false,
	"USER":              true,
	"USING":             true,
	"VACUUM":            false,
	"VALIDATE":          false,
	"VALIDATOR":         false,
	"VALID":             false,
	"VALUE":             false,
	"VALUES":            false,
	"VARCHAR":           false,
	"VARIADIC":          true,
	"VARYING":           false,
	"VERBOSE":           true,
	"VERSION":           false,
	"VIEW":              false,
	"VIEWS":             false,
	"VOLATILE":          false,
	"WHEN":              true,
	"WHERE":             true,
	"WHITESPACE":        false,
	"WINDOW":            true,
	"WITHIN":            false,
	"WITHOUT":           false,
	"WITH":              true,
	"WORK":              false,
	"WRAPPER":           false,
	"WRITE":             false,
	"XMLATTRIBUTES":     false,
	"XMLCONCAT":         false,
	"XMLELEMENT":        false,
	"XMLEXISTS":         false,
	"XML":               false,
	"XMLFOREST":         false,
	"XMLNAMESPACES":     false,
	"XMLPARSE":          false,
	"XMLPI":             false,
	"XMLROOT":           false,
	"XMLSERIALIZE":      false,
	"XMLTABLE":          false,
	"YEAR":              false,
	"YES":               false,
	"ZONE":              false,
}

var pgOperators = map[string]bool{
	"^":   true,
	"~":   true,
	"~*":  true,
	"<<":  true,
	"<=":  true,
	"<>":  true,
	"<":   true,
	"=":   true,
	">=":  true,
	">>":  true,
	">":   true,
	"||/": true,
	"||":  true,
	"|/":  true,
	"|":   true,
	"-":   true,
	":=":  true,
	"::":  true,
	"!~":  true,
	"!~*": true,
	"!=":  true,
	"!!":  true,
	"!":   true,
	"/":   true,
	"@":   true,
	"*":   true,
	"&":   true,
	"#":   true,
	"%":   true,
	"+":   true,
}

// IsPostgreSQLKeyword returns a boolean indicating if the supplied string
// is considered to be a keyword in PostgreSQL
func IsPostgreSQLKeyword(s string) bool {

	_, ok := pgKeywords[strings.ToUpper(s)]
	return ok
}

// IsPostgreSQLReservedKeyword returns a boolean indicating if the supplied
// string is considered to be a reserved keyword in PostgreSQL
func IsPostgreSQLReservedKeyword(s string) bool {

	if val, ok := pgKeywords[strings.ToUpper(s)]; ok {
		return val
	}
	return false
}

// IsPostgreSQLOperator returns a boolean indicating if the supplied string
// is considered to be an operator in PostgreSQL
func IsPostgreSQLOperator(s string) bool {

	_, ok := pgOperators[s]
	return ok
}

// IsPostgreSQLLabel returns a boolean indicating if the supplied string
// is considered to be a label in PostgreSQL
func IsPostgreSQLLabel(s string) bool {
	if len(s) < 5 {
		return false
	}
	if s[0:2] != "<<" {
		return false
	}
	if s[len(s)-2:len(s)] != ">>" {
		return false
	}
	if IsPostgreSQLIdentifier(s[2 : len(s)-2]) {
		return true
	}
	return false
}

// IsPostgreSQLIdentifier returns a boolean indicating if the supplied
// string is considered to be a non-quoted PostgreSQL identifier.
func IsPostgreSQLIdentifier(s string) bool {

	// "SQL identifiers and key words must begin with a letter (a-z, but
	// also letters with diacritical marks and non-Latin letters) or an
	// underscore (_). Subsequent characters in an identifier or key word
	// can be letters, underscores, digits (0-9), or dollar signs ($).
	// Note that dollar signs are not allowed in identifiers according to
	// the letter of the SQL standard, so their use might render
	// applications less portable."

	const firstIdentChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	const identChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_$"

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
