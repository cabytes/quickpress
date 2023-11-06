package runtime

import "time"

type Post struct {
	ID          uint64    `json:"id"`
	Type        string    `json:"type"`
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Body        string    `json:"body"`
	CreatedAt   time.Time `json:"created_at"`
}

func init() {

	db := GetDB()
	defer db.Close()

	db.Exec(
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

	db.Exec(`CREATE UNIQUE INDEX unique_post ON posts (slug)`)
}

func GetPosts() (posts []*Post, err error) {

	db := GetDB()
	defer db.Close()

	posts = make([]*Post, 0)

	rows, err := db.Query(`SELECT * FROM posts`)

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

func GetPostBySlug(slug string) (post *Post, err error) {

	db := GetDB()
	defer db.Close()

	post = &Post{}

	err = db.QueryRow(`SELECT * FROM posts WHERE slug = $1`, slug).Scan(
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
