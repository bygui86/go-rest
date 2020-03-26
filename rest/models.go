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
	routes     []*serverRoute
}

type post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type serverRoute struct {
	PathTemplate    string   `json:"pathTemplate"`
	PathRegexp      string   `json:"pathRegexp"`
	QueriesTemplate []string `json:"queriesTemplate"`
	QueriesRegexp   []string `json:"queriesRegexp"`
	Methods         []string `json:"methods"`
}

type errorResponse struct {
	Request string `json:"request"`
	Message string `json:"message"`
}
