// +build integration

package blogpost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

/*
	We are going to test every single handler isolated in respect to other handlers, starting a router handling
	just the specific request.
*/

func TestGetBlogPostsOnline(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	request := buildRequestForTesting(t, http.MethodGet, routerPostsRootUrl, nil)
	requestRecorder := httptest.NewRecorder()

	// We need to create a router that we can pass the request through, so that the 'vars' will be added to the context
	router := mux.NewRouter()
	router.HandleFunc(routerPostsRootUrl, server.getBlogPosts)
	router.ServeHTTP(requestRecorder, request)

	response := requestRecorder.Result()

	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.NotNil(t, response)
	assert.Equal(t, applicationJsonValue, response.Header.Get(contentTypeHeaderKey))

	checkGetPostsResponseForTesting(t, response, nil, server)
}

func TestGetBlogPostByPathOnline_existing(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	id := 1
	path := fmt.Sprintf("%s/%d", routerPostsRootUrl, id)
	request := buildRequestForTesting(t, http.MethodGet, path, nil)
	requestRecorder := httptest.NewRecorder()

	// We need to create a router that we can pass the request through, so that the 'vars' will be added to the context
	router := mux.NewRouter()
	router.HandleFunc(routerPostsRootUrl+routerIdUrlPath, server.getBlogPostByPath)
	router.ServeHTTP(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.NotNil(t, response)
	assert.Equal(t, applicationJsonValue, response.Header.Get(contentTypeHeaderKey))

	checkGetPostResponseForTesting(t, response, server.blogPosts[0], true)
}

func TestGetBlogPostByPathOnline_notFound(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	id := 3
	path := fmt.Sprintf("%s/%d", routerPostsRootUrl, id)
	request := buildRequestForTesting(t, http.MethodGet, path, nil)
	requestRecorder := httptest.NewRecorder()

	// We need to create a router that we can pass the request through, so that the 'vars' will be added to the context
	router := mux.NewRouter()
	router.HandleFunc(routerPostsRootUrl+routerIdUrlPath, server.getBlogPostByPath)
	router.ServeHTTP(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusNotFound, requestRecorder.Code)

	checkErrorResponseForTesting(t, response, &errorResponse{
		Request: "GET BlogPost request by path-var",
		Message: errorPostNotFoundMessage,
	})
}

func TestGetBlogPostByQueryOnline_existing(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	id := 1
	request := buildRequestForTestingWithQuery(t, http.MethodGet, routerPostsRootUrl,
		map[string]string{idKey: strconv.Itoa(id)})
	requestRecorder := httptest.NewRecorder()

	// We need to create a router that we can pass the request through, so that the 'vars' will be added to the context
	router := mux.NewRouter()
	router.HandleFunc(routerPostsRootUrl, server.getBlogPostByQuery)
	router.ServeHTTP(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.NotNil(t, response)
	assert.Equal(t, applicationJsonValue, response.Header.Get(contentTypeHeaderKey))

	checkGetPostResponseForTesting(t, response, server.blogPosts[0], true)
}

func TestGetBlogPostByQueryOnline_notFound(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	id := 3
	request := buildRequestForTestingWithQuery(t, http.MethodGet, routerPostsRootUrl,
		map[string]string{idKey: strconv.Itoa(id)})
	requestRecorder := httptest.NewRecorder()

	// We need to create a router that we can pass the request through, so that the 'vars' will be added to the context
	router := mux.NewRouter()
	router.HandleFunc(routerPostsRootUrl, server.getBlogPostByQuery)
	router.ServeHTTP(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusNotFound, requestRecorder.Code)

	checkErrorResponseForTesting(t, response, &errorResponse{
		Request: "GET BlogPost by query request",
		Message: errorPostNotFoundMessage,
	})
}

func TestCreateBlogPostOnline(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	post := &blogPost{
		Title: "test title",
		Body:  "test body",
	}
	postByte, jsonErr := json.Marshal(post)
	if jsonErr != nil {
		t.Error(jsonErr.Error())
	}
	request := buildRequestForTesting(t, http.MethodPost, routerPostsRootUrl, bytes.NewBuffer(postByte))
	requestRecorder := httptest.NewRecorder()

	// We need to create a router that we can pass the request through, so that the 'vars' will be added to the context
	router := mux.NewRouter()
	router.HandleFunc(routerPostsRootUrl, server.createBlogPost)
	router.ServeHTTP(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusCreated, requestRecorder.Code)
	assert.NotNil(t, response)
	assert.Equal(t, applicationJsonValue, response.Header.Get(contentTypeHeaderKey))

	checkGetPostResponseForTesting(t, response, post, false)
}

func TestUpdateBlogPostOnline_existing(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	id := 1
	post := &blogPost{
		ID:    strconv.Itoa(id),
		Title: "updated title",
		Body:  "updated body",
	}
	postByte, jsonErr := json.Marshal(post)
	if jsonErr != nil {
		t.Error(jsonErr.Error())
	}
	path := fmt.Sprintf("%s/%d", routerPostsRootUrl, id)
	request := buildRequestForTesting(t, http.MethodPut, path, bytes.NewBuffer(postByte))
	requestRecorder := httptest.NewRecorder()

	// We need to create a router that we can pass the request through, so that the 'vars' will be added to the context
	router := mux.NewRouter()
	router.HandleFunc(routerPostsRootUrl+routerIdUrlPath, server.updateBlogPost)
	router.ServeHTTP(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusAccepted, requestRecorder.Code)
	assert.NotNil(t, response)
	assert.Equal(t, applicationJsonValue, response.Header.Get(contentTypeHeaderKey))

	checkGetPostResponseForTesting(t, response, post, false)
}

func TestUpdateBlogPostOnline_NotFound(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	id := 3
	post := &blogPost{
		ID:    strconv.Itoa(id),
		Title: "updated title",
		Body:  "updated body",
	}
	postByte, jsonErr := json.Marshal(post)
	if jsonErr != nil {
		t.Error(jsonErr.Error())
	}
	path := fmt.Sprintf("%s/%d", routerPostsRootUrl, id)
	request := buildRequestForTesting(t, http.MethodPut, path, bytes.NewBuffer(postByte))
	requestRecorder := httptest.NewRecorder()

	// We need to create a router that we can pass the request through, so that the 'vars' will be added to the context
	router := mux.NewRouter()
	router.HandleFunc(routerPostsRootUrl+routerIdUrlPath, server.updateBlogPost)
	router.ServeHTTP(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusNotFound, requestRecorder.Code)

	checkErrorResponseForTesting(t, response, &errorResponse{
		Request: "UPDATE BlogPosts request",
		Message: errorPostNotFoundMessage,
	})
}

func TestDeleteBlogPostOnline_existing(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	id := 1
	path := fmt.Sprintf("%s/%d", routerPostsRootUrl, id)
	request := buildRequestForTesting(t, http.MethodDelete, path, nil)
	requestRecorder := httptest.NewRecorder()

	// We need to create a router that we can pass the request through, so that the 'vars' will be added to the context
	router := mux.NewRouter()
	router.HandleFunc(routerPostsRootUrl+routerIdUrlPath, server.deleteBlogPost)
	router.ServeHTTP(requestRecorder, request)

	assert.Equal(t, http.StatusAccepted, requestRecorder.Code)
}

func TestDeleteBlogPostOnline_nonExisting(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	id := 3
	path := fmt.Sprintf("%s/%d", routerPostsRootUrl, id)
	request := buildRequestForTesting(t, http.MethodDelete, path, nil)
	requestRecorder := httptest.NewRecorder()

	// We need to create a router that we can pass the request through, so that the 'vars' will be added to the context
	router := mux.NewRouter()
	router.HandleFunc(routerPostsRootUrl+routerIdUrlPath, server.deleteBlogPost)
	router.ServeHTTP(requestRecorder, request)

	assert.Equal(t, http.StatusAccepted, requestRecorder.Code)
}

func TestGetRootOnline(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	request := buildRequestForTesting(t, http.MethodGet, routerRootUrl, nil)
	requestRecorder := httptest.NewRecorder()

	// We need to create a router that we can pass the request through, so that the 'vars' will be added to the context
	router := mux.NewRouter()
	router.HandleFunc(routerRootUrl, server.getRoot)
	router.ServeHTTP(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.NotNil(t, response)
	defer closeResponseBodyForTesting(t, response.Body)
	responseBody, respBodyErr := ioutil.ReadAll(response.Body)
	assert.Nil(t, respBodyErr)
	assert.NotNil(t, responseBody)
	assert.Equal(t, rootMessage, string(responseBody))
}

func TestGetRoutesOnline(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	request := buildRequestForTesting(t, http.MethodGet, routerRoutesRootUrl, nil)
	requestRecorder := httptest.NewRecorder()

	// We need to create a router that we can pass the request through, so that the 'vars' will be added to the context
	router := mux.NewRouter()
	router.HandleFunc(routerRoutesRootUrl, server.getRoutes)
	router.ServeHTTP(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.NotNil(t, response)
	defer closeResponseBodyForTesting(t, response.Body)
	responseBody, respBodyErr := ioutil.ReadAll(response.Body)
	assert.Nil(t, respBodyErr)
	assert.NotNil(t, responseBody)

	var routes []*serverRoute
	unmarshErr := json.Unmarshal(responseBody, &routes)
	assert.Nil(t, unmarshErr)
	assert.NotNil(t, routes)
	assert.Equal(t, len(server.routes), len(routes))

	expectedRoute := server.routes[0]
	actualRoute := routes[0]
	assert.Equal(t, expectedRoute.PathTemplate, actualRoute.PathTemplate)
	assert.Equal(t, expectedRoute.PathRegexp, actualRoute.PathRegexp)
	assert.Equal(t, len(expectedRoute.QueriesTemplate), len(actualRoute.QueriesTemplate))
	assert.Equal(t, expectedRoute.QueriesTemplate[0], actualRoute.QueriesTemplate[0])
	assert.Equal(t, len(expectedRoute.QueriesRegexp), len(actualRoute.QueriesRegexp))
	assert.Equal(t, expectedRoute.QueriesRegexp[0], actualRoute.QueriesRegexp[0])
	assert.Equal(t, len(expectedRoute.Methods), len(actualRoute.Methods))
	assert.Equal(t, expectedRoute.Methods[0], actualRoute.Methods[0])
}
