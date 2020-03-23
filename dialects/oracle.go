package dialects

import "strings"

/*
Oracle keywords

https://docs.oracle.com/cd/B10501_01/appdev.920/a42525/apb.htm

Includes both SQL and PL/SQL keywords

*/

// map[keyword]isReserved
var oracleKeywords = map[string]bool{
	"ABORT":          false,
	"ACCEPT":         false,
	"ACCESS":         true,
	"ADD":            true,
	"ADMIN":          false,
	"AFTER":          false,
	"ALLOCATE":       false,
	"ALL":            true,
	"ALTER":          true,
	"ANALYZE":        false,
	"AND":            true,
	"ANY":            true,
	"ARCHIVE":        false,
	"ARCHIVELOG":     false,
	"ARRAY":          false,
	"ARRAYLEN":       true,
	"ASC":            true,
	"ASSERT":         false,
	"ASSIGN":         false,
	"AS":             true,
	"AT":             false,
	"AUDIT":          true,
	"AUTHORIZATION":  false,
	"AVG":            false,
	"BACKUP":         false,
	"BASE_TABLE":     false,
	"BECOME":         false,
	"BEFORE":         false,
	"BEGIN":          false,
	"BETWEEN":        true,
	"BINARY_INTEGER": false,
	"BLOCK":          false,
	"BODY":           false,
	"BOOLEAN":        false,
	"BY":             true,
	"CACHE":          false,
	"CANCEL":         false,
	"CASCADE":        false,
	"CASE":           false,
	"CHANGE":         false,
	"CHARACTER":      false,
	"CHAR_BASE":      false,
	"CHAR":           true,
	"CHECKPOINT":     false,
	"CHECK":          true,
	"CLOSE":          false,
	"CLUSTERS":       false,
	"CLUSTER":        true,
	"COBOL":          false,
	"COLAUTH":        false,
	"COLUMNS":        false,
	"COLUMN":         true,
	"COMMENT":        true,
	"COMMIT":         false,
	"COMPILE":        false,
	"COMPRESS":       true,
	"CONNECT":        true,
	"CONSTANT":       false,
	"CONSTRAINT":     false,
	"CONSTRAINTS":    false,
	"CONTENTS":       false,
	"CONTINUE":       false,
	"CONTROLFILE":    false,
	"COUNT":          false,
	"CRASH":          false,
	"CREATE":         true,
	"CURRENT":        true,
	"CURRVAL":        false,
	"CURSOR":         false,
	"CYCLE":          false,
	"DATA_BASE":      false,
	"DATABASE":       false,
	"DATAFILE":       false,
	"DATE":           true,
	"DBA":            false,
	"DEBUGOFF":       false,
	"DEBUGON":        false,
	"DEC":            false,
	"DECIMAL":        true,
	"DECLARE":        false,
	"DEFAULT":        true,
	"DEFINITION":     false,
	"DELAY":          false,
	"DELETE":         true,
	"DELTA":          false,
	"DESC":           true,
	"DIGITS":         false,
	"DISABLE":        false,
	"DISMOUNT":       false,
	"DISPOSE":        false,
	"DISTINCT":       true,
	"DO":             false,
	"DOUBLE":         false,
	"DROP":           true,
	"DUMP":           false,
	"EACH":           false,
	"ELSE":           true,
	"ELSIF":          false,
	"ENABLE":         false,
	"END":            false,
	"ENTRY":          false,
	"ESCAPE":         false,
	"EVENTS":         false,
	"EXCEPT":         false,
	"EXCEPTION":      false,
	"EXCEPTION_INIT": false,
	"EXCEPTIONS":     false,
	"EXCLUSIVE":      true,
	"EXEC":           false,
	"EXECUTE":        false,
	"EXISTS":         true,
	"EXIT":           false,
	"EXPLAIN":        false,
	"EXTENT":         false,
	"EXTERNALLY":     false,
	"FALSE":          false,
	"FETCH":          false,
	"FILE":           true,
	"FLOAT":          true,
	"FLUSH":          false,
	"FORCE":          false,
	"FOREIGN":        false,
	"FORM":           false,
	"FORTRAN":        false,
	"FOR":            true,
	"FOUND":          false,
	"FREELIST":       false,
	"FREELISTS":      false,
	"FROM":           true,
	"FUNCTION":       false,
	"GENERIC":        false,
	"GO":             false,
	"GOTO":           false,
	"GRANT":          true,
	"GROUPS":         false,
	"GROUP":          true,
	"HAVING":         true,
	"IDENTIFIED":     true,
	"IF":             false,
	"IMMEDIATE":      true,
	"INCLUDING":      false,
	"INCREMENT":      true,
	"INDEXES":        false,
	"INDEX":          true,
	"INDICATOR":      false,
	"INITIAL":        true,
	"INITRANS":       false,
	"INSERT":         true,
	"INSTANCE":       false,
	"INTEGER":        true,
	"INTERSECT":      true,
	"INT":            false,
	"INTO":           true,
	"IN":             true,
	"IS":             true,
	"KEY":            false,
	"LANGUAGE":       false,
	"LAYER":          false,
	"LEVEL":          true,
	"LIKE":           true,
	"LIMITED":        false,
	"LINK":           false,
	"LISTS":          false,
	"LOCK":           true,
	"LOGFILE":        false,
	"LONG":           true,
	"LOOP":           false,
	"MANAGE":         false,
	"MANUAL":         false,
	"MAXDATAFILES":   false,
	"MAXEXTENTS":     true,
	"MAX":            false,
	"MAXINSTANCES":   false,
	"MAXLOGFILES":    false,
	"MAXLOGHISTORY":  false,
	"MAXLOGMEMBERS":  false,
	"MAXTRANS":       false,
	"MAXVALUE":       false,
	"MINEXTENTS":     false,
	"MIN":            false,
	"MINUS":          true,
	"MINVALUE":       false,
	"MLSLABEL":       false,
	"MODE":           true,
	"MOD":            false,
	"MODIFY":         true,
	"MODULE":         false,
	"MOUNT":          false,
	"NATURAL":        false,
	"NEW":            false,
	"NEXT":           false,
	"NEXTVAL":        false,
	"NOARCHIVELOG":   false,
	"NOAUDIT":        true,
	"NOCACHE":        false,
	"NOCOMPRESS":     true,
	"NOCYCLE":        false,
	"NOMAXVALUE":     false,
	"NOMINVALUE":     false,
	"NONE":           false,
	"NOORDER":        false,
	"NORESETLOGS":    false,
	"NORMAL":         false,
	"NOSORT":         false,
	"NOTFOUND":       true,
	"NOT":            true,
	"NOWAIT":         true,
	"NULL":           true,
	"NUMBER_BASE":    false,
	"NUMBER":         true,
	"NUMERIC":        false,
	"OFF":            false,
	"OFFLINE":        true,
	"OF":             true,
	"OLD":            false,
	"ONLINE":         true,
	"ONLY":           false,
	"ON":             true,
	"OPEN":           false,
	"OPTIMAL":        false,
	"OPTION":         true,
	"ORDER":          true,
	"OR":             true,
	"OTHERS":         false,
	"OUT":            false,
	"OWN":            false,
	"PACKAGE":        false,
	"PARALLEL":       false,
	"PARTITION":      false,
	"PCTFREE":        true,
	"PCTINCREASE":    false,
	"PCTUSED":        false,
	"PLAN":           false,
	"PLI":            false,
	"POSITIVE":       false,
	"PRAGMA":         false,
	"PRECISION":      false,
	"PRIMARY":        false,
	"PRIOR":          true,
	"PRIVATE":        false,
	"PRIVILEGES":     true,
	"PROCEDURE":      false,
	"PROFILE":        false,
	"PUBLIC":         true,
	"QUOTA":          false,
	"RAISE":          false,
	"RANGE":          false,
	"RAW":            true,
	"READ":           false,
	"REAL":           false,
	"RECORD":         false,
	"RECOVER":        false,
	"REFERENCES":     false,
	"REFERENCING":    false,
	"RELEASE":        false,
	"REMR":           false,
	"RENAME":         true,
	"RESETLOGS":      false,
	"RESOURCE":       true,
	"RESTRICTED":     false,
	"RETURN":         false,
	"REUSE":          false,
	"REVERSE":        false,
	"REVOKE":         true,
	"ROLE":           false,
	"ROLES":          false,
	"ROLLBACK":       false,
	"ROWID":          true,
	"ROWLABEL":       true,
	"ROWNUM":         true,
	"ROWS":           true,
	"ROW":            true,
	"ROWTYPE":        false,
	"RUN":            false,
	"SAVEPOINT":      false,
	"SCHEMA":         false,
	"SCN":            false,
	"SECTION":        false,
	"SEGMENT":        false,
	"SELECT":         true,
	"SEPARATE":       false,
	"SEQUENCE":       false,
	"SESSION":        true,
	"SET":            true,
	"SHARED":         false,
	"SHARE":          true,
	"SIZE":           true,
	"SMALLINT":       true,
	"SNAPSHOT":       false,
	"SOME":           false,
	"SORT":           false,
	"SPACE":          false,
	"SQLBUF":         true,
	"SQLCODE":        false,
	"SQLERRM":        false,
	"SQLERROR":       false,
	"SQL":            false,
	"SQLSTATE":       false,
	"START":          true,
	"STATEMENT":      false,
	"STATEMENT_ID":   false,
	"STATISTICS":     false,
	"STDDEV":         false,
	"STOP":           false,
	"STORAGE":        false,
	"SUBTYPE":        false,
	"SUCCESSFUL":     true,
	"SUM":            false,
	"SWITCH":         false,
	"SYNONYM":        true,
	"SYSDATE":        true,
	"SYSTEM":         false,
	"TABAUTH":        false,
	"TABLES":         false,
	"TABLESPACE":     false,
	"TABLE":          true,
	"TASK":           false,
	"TEMPORARY":      false,
	"TERMINATE":      false,
	"THEN":           true,
	"THREAD":         false,
	"TIME":           false,
	"TO":             true,
	"TRACING":        false,
	"TRANSACTION":    false,
	"TRIGGERS":       false,
	"TRIGGER":        true,
	"TRUE":           false,
	"TRUNCATE":       false,
	"TYPE":           false,
	"UID":            true,
	"UNDER":          false,
	"UNION":          true,
	"UNIQUE":         true,
	"UNLIMITED":      false,
	"UNTIL":          false,
	"UPDATE":         true,
	"USE":            false,
	"USER":           true,
	"USING":          false,
	"VALIDATE":       true,
	"VALUES":         true,
	"VARCHAR2":       true,
	"VARCHAR":        true,
	"VARIANCE":       false,
	"VIEWS":          false,
	"VIEW":           true,
	"WHENEVER":       true,
	"WHEN":           false,
	"WHERE":          true,
	"WHILE":          false,
	"WITH":           true,
	"WORK":           false,
	"WRITE":          false,
	"XOR":            false,
}

var oracleOperators = map[string]bool{
	"^=":  true,
	"<=":  true,
	"<":   true,
	"=":   true,
	">=":  true,
	">":   true,
	"¬=":  true,
	"||":  true,
	"-":   true,
	":=":  true,
	"!=":  true,
	"/":   true,
	"(+)": true,
	"*":   true,
	"+":   true,
}

// IsOracleKeyword returns a boolean indicating if the supplied string
// is considered to be a keyword in Oracle
func IsOracleKeyword(s string) bool {

	_, ok := oracleKeywords[strings.ToUpper(s)]
	return ok
}

// IsOracleReservedKeyword returns a boolean indicating if the supplied
// string is considered to be a reserved keyword in Oracle
func IsOracleReservedKeyword(s string) bool {

	if val, ok := oracleKeywords[strings.ToUpper(s)]; ok {
		return val
	}
	return false
}

// IsOracleOperator returns a boolean indicating if the supplied string
// is considered to be an operator in Oracle
func IsOracleOperator(s string) bool {

	_, ok := oracleOperators[s]
	return ok
}

// IsOracleLabel returns a boolean indicating if the supplied string
// is considered to be a label in Oracle
func IsOracleLabel(s string) bool {
	if len(s) < 5 {
		return false
	}
	if s[0:2] != "<<" {
		return false
	}
	if s[len(s)-2:len(s)] != ">>" {
		return false
	}
	if IsOracleIdentifier(s[2 : len(s)-2]) {
		return true
	}
	return false
}

// IsOracleIdentifier returns a boolean indicating if the supplied
// string is considered to be a non-quoted Oracle identifier.
func IsOracleIdentifier(s string) bool {

	// - Nonquoted identifiers must begin with an alphabetic character
	//    from the database character set.
	// - Additonal characters may include numbers and the underscore (_),
	//    dollar sign ($), and pound sign (#)

	const firstIdentChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const identChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_#$"

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
