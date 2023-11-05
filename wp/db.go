package wp

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func init() {

	GetDB().Exec(
		`CREATE TABLE IF NOT EXISTS config (
			key VARCHAR(32) NOT NULL PRIMARY KEY,
			value  VARCHAR(255)
		);`,
	)
}

func SetConfig(name, value string) {

	db := GetDB()
	defer db.Close()

	var found = 0

	db.QueryRow("SELECT COUNT(*) as count FROM config WHERE key = $1", name).Scan(&found)

	if found == 0 {
		db.Exec(
			"INSERT INTO config (key,value) VALUES ($1, $2)",
			name,
			value,
		)
		return
	}

	db.Exec(
		"UPDATE config SET value = $1 WHERE key = $2",
		value,
		name,
	)
}

func GetConfig(name string) (v string) {

	db := GetDB()
	defer db.Close()

	row := db.QueryRow("SELECT value FROM config WHERE key = $1", name)

	if err := row.Scan(&v); err != nil {
		return ""
	}

	return
}

func GetDB() *sql.DB {

	os.MkdirAll("data/", os.ModePerm)

	db, err := sql.Open("sqlite3", "data/db")

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(1)

	return db
}
