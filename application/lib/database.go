package lib

import (
	"database/sql"
	"fmt"
	"os"
)

func NewDBConnection() (*sql.DB, error) {

	user := getEnvWithDefault("DB_USER_NAME", "root")
	pass := getEnvWithDefault("DB_PASSWORD", "")
	host := getEnvWithDefault("DB_HOST", "localhost")
	port := getEnvWithDefault("DB_PORT", "3306")
	db_name := getEnvWithDefault("DB_NAME", "app")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, db_name)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	return db, nil

}

func getEnvWithDefault(name, d string) string {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env
	}

	if len(d) != 0 {
		return d
	}
	return ""
}
