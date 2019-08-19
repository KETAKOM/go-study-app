package lib

import (
	"database/sql"
	"fmt"
	"os"
)

func NewDBConnection() (*sql.DB, error) {

	user := GetEnvWithDefault("DB_USER_NAME", "root")
	pass := GetEnvWithDefault("DB_PASSWORD", "")
	host := GetEnvWithDefault("DB_HOST", "localhost")
	port := GetEnvWithDefault("DB_PORT", "3306")
	db_name := GetEnvWithDefault("DB_NAME", "app")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, db_name)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	return db, nil

}

func GetEnvWithDefault(name, d string) string {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env
	}

	if len(d) != 0 {
		return d
	}
	return ""
}
