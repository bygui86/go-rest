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
	offline means that we are going to test every single handler method in isolation, without starting a router.
*/

func TestGetPostsOffline(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	request := buildRequestForTesting(t, http.MethodGet, routerPostsRootUrl, nil)
	requestRecorder := httptest.NewRecorder()
	server.getPosts(requestRecorder, request)
	response := requestRecorder.Result()

	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.NotNil(t, response)
	assert.Equal(t, applicationJsonValue, response.Header.Get(contentTypeHeaderKey))

	checkGetPostsResponseForTesting(t, response, nil, server)
}

// option 1: enrich url with vars and call directly the http handler method
// (see https://github.com/gorilla/mux/issues/373#issuecomment-388568971)
func TestGetPostByPathOffline_existing_option1(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	id := 1
	path := fmt.Sprintf("%s/%d", routerPostsRootUrl, id)
	request := buildRequestForTesting(t, http.MethodGet, path, nil)
	requestRecorder := httptest.NewRecorder()

	// option 1
	// To avoid the creation of a router (that we can pass the request through, so that the 'vars' will be added to the
	// context), we enrich the request with a URL var
	request = mux.SetURLVars(request, map[string]string{
		idKey: strconv.Itoa(id),
	})
	server.getPostByPath(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.NotNil(t, response)
	assert.Equal(t, applicationJsonValue, response.Header.Get(contentTypeHeaderKey))

	checkGetPostResponseForTesting(t, response, server.posts[0])
}

// option 2: enrich url with vars and use a http handler wrapper
// (see https://github.com/gorilla/mux/issues/373#issuecomment-388568971)
func TestGetPostByPathOffline_existing_option2(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	id := 1
	path := fmt.Sprintf("%s/%d", routerPostsRootUrl, id)
	request := buildRequestForTesting(t, http.MethodGet, path, nil)
	requestRecorder := httptest.NewRecorder()

	// option 2
	// To avoid the creation of a router (that we can pass the request through, so that the 'vars' will be added to the
	// context), we enrich the request with a URL var
	request = mux.SetURLVars(request, map[string]string{
		idKey: strconv.Itoa(id),
	})
	handler := http.HandlerFunc(server.getPostByPath)
	handler.ServeHTTP(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.NotNil(t, response)
	assert.Equal(t, applicationJsonValue, response.Header.Get(contentTypeHeaderKey))

	checkGetPostResponseForTesting(t, response, server.posts[0])
}

func TestGetPostByPathOffline_notFound(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	id := 3
	path := fmt.Sprintf("%s/%d", routerPostsRootUrl, id)
	request := buildRequestForTesting(t, http.MethodGet, path, nil)
	requestRecorder := httptest.NewRecorder()

	// To avoid the creation of a router (that we can pass the request through, so that the 'vars' will be added to the
	// context), we enrich the request with a URL var
	request = mux.SetURLVars(request, map[string]string{
		idKey: strconv.Itoa(id),
	})
	server.getPostByPath(requestRecorder, request)

	assert.Equal(t, http.StatusNotFound, requestRecorder.Code)
}

func TestGetPostByQueryOffline_existing(t *testing.T) {
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

func TestGetPostByQueryOffline_notFound(t *testing.T) {
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
