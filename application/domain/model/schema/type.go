package schema

const (
	// int
	TINYINT   = "tinyint"
	SMALLINT  = "smallint"
	MEDIUMINT = "mediumint"
	INT       = "int"
	BIGINT    = "bigint"

	// text
	CHAR       = "char"
	VARCHAR    = "varchar"
	TINYTEXT   = "tinytext"
	TEXT       = "text"
	MEDIUMTEXT = "mediumtext"
	LONGTEXT   = "longtext"
)

func ValidateType(s string) bool {
	switch s {
	case
		TINYINT,
		SMALLINT,
		MEDIUMINT,
		INT,
		BIGINT:
		return true
	case
		CHAR,
		VARCHAR,
		TINYTEXT,
		TEXT,
		MEDIUMTEXT,
		LONGTEXT:
		return true
	default:
		return false
	}
}
