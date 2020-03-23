package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	config     *config
	router     *mux.Router
	httpServer *http.Server
	running    bool
	posts      []*post
}

type post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
