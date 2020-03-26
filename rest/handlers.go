package rest

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/bygui86/go-rest/logging"
)

const (
	errorMessageFormat         = "Error on %s (%s): %s"
	errorEncodeResponseMessage = "encode response"
	errorDecodeResponseMessage = "decode request"
	errorPostNotFoundMessage   = "post not found"
)

func (s *Server) getPosts(writer http.ResponseWriter, request *http.Request) {
	requestStr := "GET posts request"
	logging.Log.Info(requestStr)

	setJsonContentType(writer)
	err := json.NewEncoder(writer).Encode(s.posts)
	if err != nil {
		logging.SugaredLog.Errorf(errorMessageFormat, requestStr, errorEncodeResponseMessage, err.Error())
		setStatusInternalServerError(writer)
		s.returnErrorResponse(writer, requestStr, err.Error())
	}
}

func (s *Server) getPostByPath(writer http.ResponseWriter, request *http.Request) {
	requestStr := "GET post request by path-var"
	logging.Log.Info(requestStr)

	pathVars := mux.Vars(request)
	if pathVars[idKey] == "" {
		errMsg := fmt.Sprintf("%s not found in URL", idKey)
		logging.SugaredLog.Errorf(errorMessageFormat, requestStr, "retrieve path var", errMsg)
		setStatusBadRequest(writer)
		s.returnErrorResponse(writer, requestStr, errMsg)
		return
	}

	logging.SugaredLog.Infof("Searching for post by id %s", pathVars[idKey])
	for _, item := range s.posts {
		if item.ID == pathVars[idKey] {
			setJsonContentType(writer)
			err := json.NewEncoder(writer).Encode(item)
			if err != nil {
				logging.SugaredLog.Errorf(errorMessageFormat, requestStr, errorEncodeResponseMessage, err.Error())
				setStatusInternalServerError(writer)
				s.returnErrorResponse(writer, requestStr, err.Error())
				return
			}
			return
		}
	}

	logging.SugaredLog.Warnf("Post with id %s not found", pathVars[idKey])
	setStatusNotFound(writer)
	s.returnErrorResponse(writer, requestStr, errorPostNotFoundMessage)
}

func (s *Server) getPostByQuery(writer http.ResponseWriter, request *http.Request) {
	requestStr := "GET post by query request"
	logging.Log.Info(requestStr)

	id, queryErr := s.getIdFromQueryParams(request)
	if queryErr != nil {
		logging.SugaredLog.Errorf(errorMessageFormat, requestStr, "retrieve query param", queryErr.Error())
		setStatusBadRequest(writer)
		s.returnErrorResponse(writer, requestStr, queryErr.Error())
		return
	}

	logging.SugaredLog.Infof("Searching for post by id %s", id)
	for _, item := range s.posts {
		if item.ID == id {
			setJsonContentType(writer)
			encodeErr := json.NewEncoder(writer).Encode(item)
			if encodeErr != nil {
				logging.SugaredLog.Errorf(errorMessageFormat, requestStr, encodeErr.Error())
				setStatusInternalServerError(writer)
				s.returnErrorResponse(writer, requestStr, encodeErr.Error())
				return
			}
			return
		}
	}

	logging.SugaredLog.Warnf("Post with id %s not found", id)
	setStatusNotFound(writer)
	s.returnErrorResponse(writer, requestStr, errorPostNotFoundMessage)
}

func (s *Server) createPost(writer http.ResponseWriter, request *http.Request) {
	requestStr := "CREATE post request"
	logging.Log.Info(requestStr)

	var post post
	decErr := json.NewDecoder(request.Body).Decode(&post)
	if decErr != nil {
		logging.SugaredLog.Errorf(errorMessageFormat, requestStr, errorDecodeResponseMessage, decErr.Error())
		setStatusBadRequest(writer)
		s.returnErrorResponse(writer, requestStr, decErr.Error())
		return
	}

	post.ID = strconv.Itoa(rand.Intn(1000000))
	s.posts = append(s.posts, &post)

	setJsonContentType(writer)
	setStatusCreated(writer)
	encErr := json.NewEncoder(writer).Encode(&post)
	if encErr != nil {
		logging.SugaredLog.Errorf(errorMessageFormat, requestStr, errorEncodeResponseMessage, encErr.Error())
		setStatusInternalServerError(writer)
		s.returnErrorResponse(writer, requestStr, encErr.Error())
	}
}

func (s *Server) updatePost(writer http.ResponseWriter, request *http.Request) {
	requestStr := "UPDATE posts request"
	logging.Log.Info(requestStr)

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
				logging.SugaredLog.Errorf(errorMessageFormat, requestStr, errorDecodeResponseMessage, decErr.Error())
				setStatusBadRequest(writer)
				s.returnErrorResponse(writer, requestStr, decErr.Error())
				return
			}

			post.ID = params[idKey]
			s.posts = append(s.posts, &post)

			setJsonContentType(writer)
			setStatusAccepted(writer)
			encErr := json.NewEncoder(writer).Encode(&post)
			if encErr != nil {
				logging.SugaredLog.Errorf(errorMessageFormat, requestStr, errorEncodeResponseMessage, encErr.Error())
				setStatusInternalServerError(writer)
				s.returnErrorResponse(writer, requestStr, encErr.Error())
				return
			}
			return
		}
	}

	setStatusNotFound(writer)
}

func (s *Server) deletePost(writer http.ResponseWriter, request *http.Request) {
	requestStr := "DELETE posts request"
	logging.Log.Info(requestStr)

	params := mux.Vars(request)
	for index, item := range s.posts {
		if item.ID == params[idKey] {
			s.posts = append(s.posts[:index], s.posts[index+1:]...)
			break
		}
	}

	setJsonContentType(writer)
	setStatusAccepted(writer)
}

func (s *Server) getRoot(writer http.ResponseWriter, request *http.Request) {
	requestStr := "GET root request"
	logging.Log.Info(requestStr)

	_, err := writer.Write([]byte("Welcome to go-rest sample project"))
	if err != nil {
		logging.SugaredLog.Errorf(errorMessageFormat, requestStr, "write string response", err.Error())
		setStatusInternalServerError(writer)
		s.returnErrorResponse(writer, requestStr, err.Error())
	}
}

func (s *Server) getRoutes(writer http.ResponseWriter, request *http.Request) {
	requestStr := "GET routes request"
	logging.Log.Info(requestStr)

	if s.routes == nil || len(s.routes) < 1 {
		walkErr := s.router.Walk(s.walkRoute)
		if walkErr != nil {
			logging.SugaredLog.Errorf(errorMessageFormat, requestStr, "walk through routes", walkErr.Error())
			setStatusInternalServerError(writer)
			s.returnErrorResponse(writer, requestStr, walkErr.Error())
			return
		}
	}

	setJsonContentType(writer)
	err := json.NewEncoder(writer).Encode(s.routes)
	if err != nil {
		logging.SugaredLog.Errorf(errorMessageFormat, requestStr, errorEncodeResponseMessage, err.Error())
		setStatusInternalServerError(writer)
		s.returnErrorResponse(writer, requestStr, err.Error())
	}
}
