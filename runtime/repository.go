package runtime

import (
	"database/sql"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gosimple/slug"
)

type Post struct {
	ID          uint64    `json:"id"`
	Type        string    `json:"type"`
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Body        string    `json:"body"`
	CreatedAt   time.Time `json:"created_at"`
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) (repo *Repository, err error) {

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	repo = &Repository{db}

	repo.init()

	return
}

func (repo *Repository) init() {

	repo.db.Exec(
		`CREATE TABLE IF NOT EXISTS config (
			key VARCHAR(32) NOT NULL PRIMARY KEY,
			value VARCHAR(255)
		);`,
	)

	repo.db.Exec(
		`CREATE TABLE IF NOT EXISTS posts (
			id INTEGER NOT NULL PRIMARY KEY,
			type VARCHAR(16) CHECK( type IN ('post', 'page') ) NOT NULL DEFAULT 'post',
			slug VARCHAR(255) NOT NULL,
			title VARCHAR(120) NOT NULL,
			description VARCHAR(500),
			body TEXT,
			created_at DATETIME DEFAULT (datetime('now','localtime'))
		)`,
	)

	repo.db.Exec(`CREATE UNIQUE INDEX unique_post ON posts (slug)`)

	repo.DefineConfig(map[string]string{
		"blog_title_separator": " - ",
		"blog_title":           "Awesome Blog",
	})
}

// Build blog title for title tag
func (repo *Repository) BlogTitle(title string) string {
	return strings.ToUpper(title[:1]) + title[1:] + repo.GetConfig("blog_title_separator") + repo.GetConfig("blog_title")
}

func (repo *Repository) DefineConfig(config map[string]string) error {

	for name, value := range config {

		var found = 0

		repo.db.QueryRow("SELECT COUNT(*) as count FROM config WHERE key = $1", name).Scan(&found)

		if found == 0 {
			repo.db.Exec(
				"INSERT INTO config (key,value) VALUES ($1, $2)",
				name,
				value,
			)
		}
	}

	return nil
}

func (repo *Repository) SetConfig(name, value string) {

	var found = 0

	repo.db.QueryRow("SELECT COUNT(*) as count FROM config WHERE key = $1", name).Scan(&found)

	if found == 0 {
		repo.db.Exec(
			"INSERT INTO config (key,value) VALUES ($1, $2)",
			name,
			value,
		)
		return
	}

	repo.db.Exec(
		"UPDATE config SET value = $1 WHERE key = $2",
		value,
		name,
	)
}

func (repo *Repository) GetConfig(name string) (v string) {

	row := repo.db.QueryRow("SELECT value FROM config WHERE key = $1", name)

	if err := row.Scan(&v); err != nil {
		return ""
	}

	return
}

func (repo *Repository) GetPosts() (posts []*Post, err error) {

	posts = make([]*Post, 0)

	rows, err := repo.db.Query(`SELECT * FROM posts`)

	for rows.Next() {

		post := &Post{}

		err = rows.Scan(
			&post.ID,
			&post.Type,
			&post.Slug,
			&post.Title,
			&post.Description,
			&post.Body,
			&post.CreatedAt,
		)

		if err != nil {
			return posts, err
		}

		posts = append(posts, post)
	}

	return
}

func (repo *Repository) GetPostBySlug(slugInput string) (post *Post, err error) {

	slug := slug.Make(slugInput)

	post = &Post{}

	err = repo.db.QueryRow(`SELECT * FROM posts WHERE slug = $1`, slug).Scan(
		&post.ID,
		&post.Type,
		&post.Slug,
		&post.Title,
		&post.Description,
		&post.Body,
		&post.CreatedAt,
	)

	return
}
