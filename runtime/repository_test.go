package runtime

import (
	"database/sql"
	"os"
	"testing"
)

func testingRepo() (repo *Repository, err error) {

	tmpf, err := os.CreateTemp(".", "testdb")

	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", tmpf.Name())

	if err != nil {
		return nil, err
	}

	return NewRepository(db)
}

func TestGetPosts(t *testing.T) {

	repo, err := testingRepo()

	if err != nil {
		t.Fatal(err.Error())
		return
	}

	repo.Create(&Post{
		Slug: "test1",
	})

	repo.Create(&Post{
		Slug: "test2",
	})

	posts, err := repo.GetPosts()

	if err != nil {
		t.Fatal(err.Error())
		return
	}

	if len(posts) != 2 {
		t.Fatal("Expected two posts")
	}
}

func TestGetBySlug(t *testing.T) {

	repo, err := testingRepo()

	if err != nil {
		t.Fatal(err.Error())
		return
	}

	repo.Create(&Post{
		Slug: "test",
	})

	post, err := repo.GetPostBySlug("test")

	if err != nil {
		t.Fatal(err.Error())
		return
	}

	if post == nil {
		t.Fatal("Expected a post for test slug")
	}
}
