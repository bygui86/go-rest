package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Config     *Config
	Router     *mux.Router
	HTTPServer *http.Server
	Running    bool
	posts      []*Post
}

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
