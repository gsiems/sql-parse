package dialects

import "strings"

/*
Microsoft SQL-Server keywords

https://docs.microsoft.com/en-us/sql/t-sql/language-elements/reserved-keywords-transact-sql?view=sql-server-ver15

The isReserved value is set to false as there is no indication (from
the above link) if the keywords are reserved or not.

*/

// map[keyword]isReserved
var mssqlKeywords = map[string]bool{
	"ADD":                            false,
	"ALL":                            false,
	"ALTER":                          false,
	"AND":                            false,
	"ANY":                            false,
	"AS":                             false,
	"ASC":                            false,
	"AUTHORIZATION":                  false,
	"BACKUP":                         false,
	"BEGIN":                          false,
	"BETWEEN":                        false,
	"BREAK":                          false,
	"BROWSE":                         false,
	"BULK":                           false,
	"BY":                             false,
	"CASCADE":                        false,
	"CASE":                           false,
	"CHECK":                          false,
	"CHECKPOINT":                     false,
	"CLOSE":                          false,
	"CLUSTERED":                      false,
	"COALESCE":                       false,
	"COLLATE":                        false,
	"COLUMN":                         false,
	"COMMIT":                         false,
	"COMPUTE":                        false,
	"CONSTRAINT":                     false,
	"CONTAINS":                       false,
	"CONTAINSTABLE":                  false,
	"CONTINUE":                       false,
	"CONVERT":                        false,
	"CREATE":                         false,
	"CROSS":                          false,
	"CURRENT":                        false,
	"CURRENT_DATE":                   false,
	"CURRENT_TIME":                   false,
	"CURRENT_TIMESTAMP":              false,
	"CURRENT_USER":                   false,
	"CURSOR":                         false,
	"DATABASE":                       false,
	"DBCC":                           false,
	"DEALLOCATE":                     false,
	"DECLARE":                        false,
	"DEFAULT":                        false,
	"DELETE":                         false,
	"DENY":                           false,
	"DESC":                           false,
	"DISK":                           false,
	"DISTINCT":                       false,
	"DISTRIBUTED":                    false,
	"DOUBLE":                         false,
	"DROP":                           false,
	"DUMP":                           false,
	"ELSE":                           false,
	"END":                            false,
	"ERRLVL":                         false,
	"ESCAPE":                         false,
	"EXCEPT":                         false,
	"EXEC":                           false,
	"EXECUTE":                        false,
	"EXISTS":                         false,
	"EXIT":                           false,
	"EXTERNAL":                       false,
	"FETCH":                          false,
	"FILE":                           false,
	"FILLFACTOR":                     false,
	"FOR":                            false,
	"FOREIGN":                        false,
	"FREETEXT":                       false,
	"FREETEXTTABLE":                  false,
	"FROM":                           false,
	"FULL":                           false,
	"FUNCTION":                       false,
	"GOTO":                           false,
	"GRANT":                          false,
	"GROUP":                          false,
	"HAVING":                         false,
	"HOLDLOCK":                       false,
	"IDENTITY":                       false,
	"IDENTITYCOL":                    false,
	"IDENTITY_INSERT":                false,
	"IF":                             false,
	"IN":                             false,
	"INDEX":                          false,
	"INNER":                          false,
	"INSERT":                         false,
	"INTERSECT":                      false,
	"INTO":                           false,
	"IS":                             false,
	"JOIN":                           false,
	"KEY":                            false,
	"KILL":                           false,
	"LABEL":                          false,
	"LEFT":                           false,
	"LIKE":                           false,
	"LINENO":                         false,
	"LOAD":                           false,
	"MERGE":                          false,
	"NATIONAL":                       false,
	"NOCHECK":                        false,
	"NONCLUSTERED":                   false,
	"NOT":                            false,
	"NULL":                           false,
	"NULLIF":                         false,
	"OF":                             false,
	"OFF":                            false,
	"OFFSETS":                        false,
	"ON":                             false,
	"OPEN":                           false,
	"OPENDATASOURCE":                 false,
	"OPENQUERY":                      false,
	"OPENROWSET":                     false,
	"OPENXML":                        false,
	"OPTION":                         false,
	"OR":                             false,
	"ORDER":                          false,
	"OUTER":                          false,
	"OVER":                           false,
	"PERCENT":                        false,
	"PIVOT":                          false,
	"PLAN":                           false,
	"PRECISION":                      false,
	"PRIMARY":                        false,
	"PRINT":                          false,
	"PROC":                           false,
	"PROCEDURE":                      false,
	"PUBLIC":                         false,
	"RAISERROR":                      false,
	"READ":                           false,
	"READTEXT":                       false,
	"RECONFIGURE":                    false,
	"REFERENCES":                     false,
	"REPLICATION":                    false,
	"RESTORE":                        false,
	"RESTRICT":                       false,
	"RETURN":                         false,
	"REVERT":                         false,
	"REVOKE":                         false,
	"RIGHT":                          false,
	"ROLLBACK":                       false,
	"ROWCOUNT":                       false,
	"ROWGUIDCOL":                     false,
	"RULE":                           false,
	"SAVE":                           false,
	"SCHEMA":                         false,
	"SECURITYAUDIT":                  false,
	"SELECT":                         false,
	"SEMANTICKEYPHRASETABLE":         false,
	"SEMANTICSIMILARITYDETAILSTABLE": false,
	"SEMANTICSIMILARITYTABLE":        false,
	"SESSION_USER":                   false,
	"SET":                            false,
	"SETUSER":                        false,
	"SHUTDOWN":                       false,
	"SOME":                           false,
	"STATISTICS":                     false,
	"SYSTEM_USER":                    false,
	"TABLE":                          false,
	"TABLESAMPLE":                    false,
	"TEXTSIZE":                       false,
	"THEN":                           false,
	"TO":                             false,
	"TOP":                            false,
	"TRAN":                           false,
	"TRANSACTION":                    false,
	"TRIGGER":                        false,
	"TRUNCATE":                       false,
	"TRY_CONVERT":                    false,
	"TSEQUAL":                        false,
	"UNION":                          false,
	"UNIQUE":                         false,
	"UNPIVOT":                        false,
	"UPDATE":                         false,
	"UPDATETEXT":                     false,
	"USE":                            false,
	"USER":                           false,
	"VALUES":                         false,
	"VARYING":                        false,
	"VIEW":                           false,
	"WAITFOR":                        false,
	"WHEN":                           false,
	"WHERE":                          false,
	"WHILE":                          false,
	"WITH":                           false,
	"WITHIN GROUP":                   false,
	"WRITETEXT":                      false,
}

var mssqlOperators = map[string]bool{
	"^":  true,
	"^=": true,
	"~":  true,
	"<":  true,
	"<=": true,
	"<>": true,
	"=":  true,
	">":  true,
	">=": true,
	"|":  true,
	"|=": true,
	"-":  true,
	"-=": true,
	"::": true,
	"!<": true,
	"!=": true,
	"!>": true,
	"/":  true,
	"/=": true,
	"*":  true,
	"*=": true,
	"&":  true,
	"&=": true,
	"%":  true,
	"%=": true,
	"+":  true,
	"+=": true,
}

// IsMSSQLKeyword returns a boolean indicating if the supplied string
// is considered to be a keyword in MS-SQL
func IsMSSQLKeyword(s string) bool {

	_, ok := mssqlKeywords[strings.ToUpper(s)]
	return ok
}

// IsMSSQLReservedKeyword returns a boolean indicating if the supplied
// string is considered to be a reserved keyword in MS-SQL
func IsMSSQLReservedKeyword(s string) bool {

	if val, ok := mssqlKeywords[strings.ToUpper(s)]; ok {
		return val
	}
	return false
}

// IsMSSQLOperator returns a boolean indicating if the supplied string
// is considered to be an operator in MSSQL
func IsMSSQLOperator(s string) bool {

	_, ok := mssqlOperators[s]
	return ok
}

// IsMSSQLIdentifier returns a boolean indicating if the supplied
// string is considered to be a non-quoted MSSQL identifier.
func IsMSSQLIdentifier(s string) bool {

	/*

		From the documentstion found:

		   The first character must be one of the following:

		       A letter as defined by the Unicode Standard 3.2. The Unicode
		       definition of letters includes Latin characters from a through
		       z, from A through Z, and also letter characters from other
		       languages.

		       The underscore (_), at sign (@), or number sign (#).

		   ...

		   Subsequent characters can include the following:
		       Letters as defined in the Unicode Standard 3.2.
		       Decimal numbers from either Basic Latin or other national scripts.
		       The at sign, dollar sign ($), number sign, or underscore.

		       Embedded spaces or special characters are not allowed.

		       Supplementary characters are not allowed.

	*/

	const firstIdentChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_#@"
	const identChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_#@$"

	chr := strings.Split(s, "")
	for i := 0; i < len(chr); i++ {

		//matches := false

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
