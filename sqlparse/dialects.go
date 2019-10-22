package sqlparse

import (
	d "github.com/gsiems/sql-parse/dialects"
)

/*

dialects.go provides the list of SQL dialects that sqlparse will attempt
to tokenize.

*/

// SQL Dialects
const (
	NullDialect = iota
	StandardSQL
	PostgreSQL
	SQLite
	MySQL
	Oracle
	MSSQL
	MariaDB
)

// SQLDialectName returns the string representation of the SQL dialect
func SQLDialectName(d int) (s string) {

	var names = map[int]string{
		StandardSQL: "StandardSQL",
		PostgreSQL:  "PostgreSQL",
		SQLite:      "SQLite",
		MySQL:       "MySQL",
		Oracle:      "Oracle",
		MSSQL:       "MSSQL",
		MariaDB:     "MariaDB",
	}

	if s, ok := names[d]; ok {
		return s
	}
	return ""
}

// SQLDialect returns the integer value of the SQL dialect
func SQLDialect(s string) (d int) {

	var vals = map[string]int{
		"StandardSQL": StandardSQL,
		"PostgreSQL":  PostgreSQL,
		"SQLite":      SQLite,
		"MySQL":       MySQL,
		"Oracle":      Oracle,
		"MSSQL":       MSSQL,
		"MariaDB":     MariaDB,
	}

	if d, ok := vals[s]; ok {
		return d
	}
	return 0
}

// IsKeyword returns true if the supplied string is defined as a
// keyword for the specified SQL dialect
func IsKeyword(s string, dialect int) bool {

	switch dialect {
	case PostgreSQL:
		return d.IsPostgreSQLKeyword(s)
	case SQLite:
		return d.IsSQLiteKeyword(s)
	case MySQL:
		return d.IsMySQLKeyword(s)
	case Oracle:
		return d.IsOracleKeyword(s)
	case MSSQL:
		return d.IsMSSQLKeyword(s)
	case MariaDB:
		return d.IsMariaDBKeyword(s)
	default:
		return d.IsStandardKeyword(s)
	}
}

// IsReservedKeyword returns true if the supplied string is defined as a
// reserved keyword for the specified SQL dialect
func IsReservedKeyword(s string, dialect int) bool {

	switch dialect {
	case PostgreSQL:
		return d.IsPostgreSQLReservedKeyword(s)
	case SQLite:
		return d.IsSQLiteReservedKeyword(s)
	case MySQL:
		return d.IsMySQLReservedKeyword(s)
	case Oracle:
		return d.IsOracleReservedKeyword(s)
	case MSSQL:
		return d.IsMSSQLReservedKeyword(s)
	case MariaDB:
		return d.IsMariaDBReservedKeyword(s)
	default:
		return d.IsStandardReservedKeyword(s)
	}
}

// IsIdentifier returns true if the supplied string is considered to be
// a non-quoted identifier for the specified SQL dialect
func IsIdentifier(s string, dialect int) bool {

	switch dialect {
	case PostgreSQL:
		return d.IsPostgreSQLIdentifier(s)
	case SQLite:
		return d.IsSQLiteIdentifier(s)
	case MySQL:
		return d.IsMySQLIdentifier(s)
	case Oracle:
		return d.IsOracleIdentifier(s)
	case MSSQL:
		return d.IsMSSQLIdentifier(s)
	case MariaDB:
		return d.IsMariaDBIdentifier(s)
	default:
		return d.IsStandardIdentifier(s)
	}
}
