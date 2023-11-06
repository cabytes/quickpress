package wp

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/pflag"
)

func SetupAPI(mux *chi.Mux) {
	mux.Post("/api/auth", auth)
	mux.Post("/api/posts", createPost)
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func auth(w http.ResponseWriter, r *http.Request) {

	payload := AuthPayload{}

	data, err := io.ReadAll(r.Body)

	if err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(data, &payload); err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	password, err := pflag.CommandLine.GetString("password")

	if err != nil {
		panic(err)
	}

	if password == payload.Password {
		w.Write([]byte("Authenticated!"))
		return
	}

	w.Write([]byte("Hello world!"))
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}
