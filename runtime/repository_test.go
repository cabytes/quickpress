package runtime

import (
	"database/sql"
	"os"
	"testing"
)

func TestGetBySlug(t *testing.T) {

	tmpf, err := os.CreateTemp(".", "testdb")

	defer os.Remove(tmpf.Name())

	if err != nil {
		t.Fatal(err)
		return
	}

	db, err := sql.Open("sqlite3", tmpf.Name())

	if err != nil {
		t.Fatal(err.Error())
		return
	}

	db.SetMaxOpenConns(1)

	repo, err := NewRepository(db)

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

	db.Close()
}
