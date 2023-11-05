package wp

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func init() {

	os.MkdirAll("data/", os.ModePerm)

	db := GetDB()
	defer db.Close()

	db.Exec(
		`CREATE TABLE IF NOT EXISTS config (
			key VARCHAR(32) NOT NULL PRIMARY KEY,
			value VARCHAR(255)
		);`,
	)

	db.Exec(
		`CREATE TABLE IF NOT EXISTS posts (
			id INTEGER NOT NULL PRIMARY KEY,
			slug VARCHAR(255) NOT NULL,
			title VARCHAR(120) NOT NULL,
			description VARCHAR(500),
			body TEXT
		)`,
	)

	db.Exec(`CREATE UNIQUE INDEX unique_post ON posts (slug)`)
}

func GetPostBySlug(slug string) (post *Post, err error) {

	db := GetDB()
	defer db.Close()

	post = &Post{}

	err = db.QueryRow(`SELECT * FROM posts WHERE slug = $1`, slug).Scan(
		&post.ID,
		&post.Slug,
		&post.Title,
		&post.Description,
		&post.Body,
	)

	return
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

	db, err := sql.Open("sqlite3", "data/db")

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(1)

	return db
}
