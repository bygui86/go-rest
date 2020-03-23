package rest

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/bygui86/go-rest/logging"
)

const (
	contentTypeHeaderKey   = "Content-Type"
	contentTypeHeaderValue = "application/json"
)

func (s *Server) getPosts(writer http.ResponseWriter, request *http.Request) {
	logging.Log.Info("GET posts request")

	setJsonContentType(writer)
	err := json.NewEncoder(writer).Encode(s.posts)
	if err != nil {
		logging.SugaredLog.Errorf("Error on GET posts request (encode response): %s", err.Error())
	}
}

func (s *Server) createPost(writer http.ResponseWriter, request *http.Request) {
	logging.Log.Info("CREATE post request")

	setJsonContentType(writer)
	setStatusAccepted(writer)
	var post post
	decErr := json.NewDecoder(request.Body).Decode(&post)
	if decErr != nil {
		logging.SugaredLog.Errorf("Error on CREATE post request (decode request): %s", decErr.Error())
	}

	post.ID = strconv.Itoa(rand.Intn(1000000))
	s.posts = append(s.posts, &post)
	encErr := json.NewEncoder(writer).Encode(&post)
	if encErr != nil {
		logging.SugaredLog.Errorf("Error on CREATE post request (encode response): %s", encErr.Error())
	}
}

func (s *Server) getPost(writer http.ResponseWriter, request *http.Request) {
	logging.Log.Info("GET post request")

	setJsonContentType(writer)
	params := mux.Vars(request)
	for _, item := range s.posts {
		if item.ID == params[idKey] {
			err := json.NewEncoder(writer).Encode(item)
			if err != nil {
				logging.SugaredLog.Errorf("Error on GET post request (encode response): %s", err.Error())
			}
			return
		}
	}

	err := json.NewEncoder(writer).Encode(&post{})
	if err != nil {
		logging.SugaredLog.Errorf("Error on GET post request (encode response): %s", err.Error())
	}
}

func (s *Server) updatePost(writer http.ResponseWriter, request *http.Request) {
	logging.Log.Info("UPDATE posts request")

	setJsonContentType(writer)
	setStatusAccepted(writer)
	params := mux.Vars(request)
	for index, item := range s.posts {
		if item.ID == params[idKey] {
			/*
				posts[:index] >> from the beginning to index position
				posts[index+1:]... >> from index+1 position to the end
			*/
			s.posts = append(s.posts[:index], s.posts[index+1:]...)
			var post post
			decErr := json.NewDecoder(request.Body).Decode(&post)
			if decErr != nil {
				logging.SugaredLog.Errorf("Error on UPDATE post request (decode request): %s", decErr.Error())
			}

			post.ID = params[idKey]
			s.posts = append(s.posts, &post)
			encErr := json.NewEncoder(writer).Encode(&post)
			if encErr != nil {
				logging.SugaredLog.Errorf("Error on UPDATE post request (encode response): %s", encErr.Error())
			}
			return
		}
	}

	err := json.NewEncoder(writer).Encode(s.posts)
	if err != nil {
		logging.SugaredLog.Errorf("Error on UPDATE post request (encode response): %s", err.Error())
	}
}

func (s *Server) deletePost(writer http.ResponseWriter, request *http.Request) {
	logging.Log.Info("DELETE posts request")

	setJsonContentType(writer)
	setStatusAccepted(writer)
	params := mux.Vars(request)
	for index, item := range s.posts {

		if item.ID == params[idKey] {
			s.posts = append(s.posts[:index], s.posts[index+1:]...)
			break
		}
	}

	err := json.NewEncoder(writer).Encode(s.posts)
	if err != nil {
		logging.SugaredLog.Errorf("Error on DELETE post request (encode response): %s", err.Error())
	}
}
