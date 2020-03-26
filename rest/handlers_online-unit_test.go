// build unit

package rest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

/*
	INFO
	online means that we are going to test every single handler isolated in respect to other handlers, starting
	a router handling just the specific request.
*/

func TestGetPostsOnline(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	request := buildRequestForTesting(t, http.MethodGet, routerPostsRootUrl, nil)
	requestRecorder := httptest.NewRecorder()

	// We need to create a router that we can pass the request through, so that the 'vars' will be added to the context
	router := mux.NewRouter()
	router.HandleFunc(routerPostsRootUrl, server.getPosts)
	router.ServeHTTP(requestRecorder, request)

	response := requestRecorder.Result()

	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.NotNil(t, response)
	assert.Equal(t, applicationJsonValue, response.Header.Get(contentTypeHeaderKey))

	checkGetPostsResponseForTesting(t, response, nil, server)
}

func TestGetPostByPathOnline_existing(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	id := 1
	path := fmt.Sprintf("%s/%d", routerPostsRootUrl, id)
	request := buildRequestForTesting(t, http.MethodGet, path, nil)
	requestRecorder := httptest.NewRecorder()

	// We need to create a router that we can pass the request through, so that the 'vars' will be added to the context
	router := mux.NewRouter()
	router.HandleFunc(routerPostsRootUrl+routerIdUrlPath, server.getPostByPath)
	router.ServeHTTP(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.NotNil(t, response)
	assert.Equal(t, applicationJsonValue, response.Header.Get(contentTypeHeaderKey))

	checkGetPostResponseForTesting(t, response, server.posts[0])
}

func TestGetPostByPathOnline_notFound(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	id := 3
	path := fmt.Sprintf("%s/%d", routerPostsRootUrl, id)
	request := buildRequestForTesting(t, http.MethodGet, path, nil)
	requestRecorder := httptest.NewRecorder()

	// method 3
	// We need to create a router that we can pass the request through, so that the 'vars' will be added to the context
	router := mux.NewRouter()
	router.HandleFunc(routerPostsRootUrl+routerIdUrlPath, server.getPostByPath)
	router.ServeHTTP(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusNotFound, requestRecorder.Code)
	assert.NotNil(t, response)
	assert.Equal(t, applicationJsonValue, response.Header.Get(contentTypeHeaderKey))
}

func TestGetPostByQueryOnline_existing(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	id := 1
	request := buildRequestForTestingWithQuery(t, http.MethodGet, routerPostsRootUrl,
		map[string]string{idKey: strconv.Itoa(id)})
	requestRecorder := httptest.NewRecorder()

	server.getPostByQuery(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.NotNil(t, response)
	assert.Equal(t, applicationJsonValue, response.Header.Get(contentTypeHeaderKey))

	checkGetPostResponseForTesting(t, response, server.posts[0])
}

func TestGetPostByQueryOnline_notFound(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	id := 3
	request := buildRequestForTestingWithQuery(t, http.MethodGet, routerPostsRootUrl,
		map[string]string{idKey: strconv.Itoa(id)})
	requestRecorder := httptest.NewRecorder()

	server.getPostByQuery(requestRecorder, request)

	assert.Equal(t, http.StatusNotFound, requestRecorder.Code)
}

// TODO missing createPost, updatePost, deletePost, getRoutes, getRoot
