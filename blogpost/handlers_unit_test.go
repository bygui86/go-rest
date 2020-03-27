// +build unit

package blogpost

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

/*
	We are going to test every single handler method in isolation, without starting a router.
*/

const (
	whateverStr = "whatever"
)

func TestGetBlogPostsOffline(t *testing.T) {
	// We don't need to start the server, we just need it initialized to have access to its methods and fields
	server := NewRestServer()
	// We don't need to set method and path on the request as we are going to call directly the specific handler
	request := buildRequestForTesting(t, whateverStr, whateverStr, nil)
	requestRecorder := httptest.NewRecorder()

	server.getBlogPosts(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.NotNil(t, response)
	assert.Equal(t, applicationJsonValue, response.Header.Get(contentTypeHeaderKey))
	checkGetPostsResponseForTesting(t, response, nil, server)
}

// option 1: enrich url with vars and call directly the http handler method
// (see https://github.com/gorilla/mux/issues/373#issuecomment-388568971)
func TestGetBlogPostByPathOffline_existing_option1(t *testing.T) {
	// We don't need to start the server, we just need it initialized to have access to its methods and fields
	server := NewRestServer()
	// We don't need to set method and path on the request as we are going to call directly the specific handler
	request := buildRequestForTesting(t, whateverStr, whateverStr, nil)
	requestRecorder := httptest.NewRecorder()

	// option 1
	// To avoid the creation of a router (that we can pass the request through, so that the 'vars' will be added to the
	// context), we enrich the request with a URL var
	id := 1
	request = mux.SetURLVars(request, map[string]string{
		idKey: strconv.Itoa(id),
	})
	server.getBlogPostByPath(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.NotNil(t, response)
	assert.Equal(t, applicationJsonValue, response.Header.Get(contentTypeHeaderKey))
	checkGetPostResponseForTesting(t, response, server.blogPosts[0], true)
}

// option 2: enrich url with vars and use a http handler wrapper
// (see https://github.com/gorilla/mux/issues/373#issuecomment-388568971)
func TestGetBlogPostByPathOffline_existing_option2(t *testing.T) {
	// We don't need to start the server, we just need it initialized to have access to its methods and fields
	server := NewRestServer()
	// We don't need to set method and path on the request as we are going to call directly the specific handler
	request := buildRequestForTesting(t, whateverStr, whateverStr, nil)
	requestRecorder := httptest.NewRecorder()

	// option 2
	// To avoid the creation of a router (that we can pass the request through, so that the 'vars' will be added to the
	// context), we enrich the request with a URL var
	id := 1
	request = mux.SetURLVars(request, map[string]string{
		idKey: strconv.Itoa(id),
	})
	handler := http.HandlerFunc(server.getBlogPostByPath)
	handler.ServeHTTP(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.NotNil(t, response)
	assert.Equal(t, applicationJsonValue, response.Header.Get(contentTypeHeaderKey))

	checkGetPostResponseForTesting(t, response, server.blogPosts[0], true)
}

func TestGetBlogPostByPathOffline_notFound(t *testing.T) {
	// We don't need to start the server, we just need it initialized to have access to its methods and fields
	server := NewRestServer()
	// We don't need to set method and path on the request as we are going to call directly the specific handler
	request := buildRequestForTesting(t, whateverStr, whateverStr, nil)
	requestRecorder := httptest.NewRecorder()

	// To avoid the creation of a router (that we can pass the request through, so that the 'vars' will be added to the
	// context), we enrich the request with a URL var
	id := 3
	request = mux.SetURLVars(request, map[string]string{
		idKey: strconv.Itoa(id),
	})
	server.getBlogPostByPath(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusNotFound, requestRecorder.Code)

	checkErrorResponseForTesting(t, response, &errorResponse{
		Request: "GET BlogPost request by path-var",
		Message: errorPostNotFoundMessage,
	})
}

func TestGetBlogPostByQueryOffline_existing(t *testing.T) {
	// We don't need to start the server, we just need it initialized to have access to its methods and fields
	server := NewRestServer()
	// We don't need to set method and path on the request as we are going to call directly the specific handler
	id := 1
	request := buildRequestForTestingWithQuery(t, whateverStr, whateverStr,
		map[string]string{idKey: strconv.Itoa(id)})
	requestRecorder := httptest.NewRecorder()

	server.getBlogPostByQuery(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.NotNil(t, response)
	assert.Equal(t, applicationJsonValue, response.Header.Get(contentTypeHeaderKey))

	checkGetPostResponseForTesting(t, response, server.blogPosts[0], true)
}

func TestGetBlogPostByQueryOffline_notFound(t *testing.T) {
	// We don't need to start the server, we just need it initialized to have access to its methods and fields
	server := NewRestServer()

	id := 3
	request := buildRequestForTestingWithQuery(t, http.MethodGet, routerPostsRootUrl,
		map[string]string{idKey: strconv.Itoa(id)})
	requestRecorder := httptest.NewRecorder()

	server.getBlogPostByQuery(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusNotFound, requestRecorder.Code)
	checkErrorResponseForTesting(t, response, &errorResponse{
		Request: "GET BlogPost by query request",
		Message: errorPostNotFoundMessage,
	})
}

func TestCreateBlogPostOffline(t *testing.T) {
	// We don't need to start the server, we just need it initialized to have access to its methods and fields
	server := NewRestServer()
	post := &blogPost{
		Title: "test title",
		Body:  "test body",
	}
	postByte, jsonErr := json.Marshal(post)
	if jsonErr != nil {
		t.Error(jsonErr.Error())
	}
	// We don't need to set method and path on the request as we are going to call directly the specific handler
	request := buildRequestForTesting(t, whateverStr, whateverStr, bytes.NewBuffer(postByte))
	requestRecorder := httptest.NewRecorder()

	server.createBlogPost(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusCreated, requestRecorder.Code)
	assert.NotNil(t, response)
	assert.Equal(t, applicationJsonValue, response.Header.Get(contentTypeHeaderKey))
	checkGetPostResponseForTesting(t, response, post, false)
}

func TestUpdateBlogPostOffline_existing(t *testing.T) {
	// We don't need to start the server, we just need it initialized to have access to its methods and fields
	server := NewRestServer()
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
	// We don't need to set method and path on the request as we are going to call directly the specific handler
	request := buildRequestForTesting(t, whateverStr, whateverStr, bytes.NewBuffer(postByte))
	requestRecorder := httptest.NewRecorder()

	request = mux.SetURLVars(request, map[string]string{
		idKey: strconv.Itoa(id),
	})
	server.updateBlogPost(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusAccepted, requestRecorder.Code)
	assert.NotNil(t, response)
	assert.Equal(t, applicationJsonValue, response.Header.Get(contentTypeHeaderKey))
	checkGetPostResponseForTesting(t, response, post, false)
}

func TestUpdateBlogPostOffline_NotFound(t *testing.T) {
	// We don't need to start the server, we just need it initialized to have access to its methods and fields
	server := NewRestServer()
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
	// We don't need to set method and path on the request as we are going to call directly the specific handler
	request := buildRequestForTesting(t, whateverStr, whateverStr, bytes.NewBuffer(postByte))
	requestRecorder := httptest.NewRecorder()

	request = mux.SetURLVars(request, map[string]string{
		idKey: strconv.Itoa(id),
	})
	server.updateBlogPost(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusNotFound, requestRecorder.Code)
	checkErrorResponseForTesting(t, response, &errorResponse{
		Request: "UPDATE BlogPosts request",
		Message: errorPostNotFoundMessage,
	})
}

func TestDeleteBlogPostOffline_existing(t *testing.T) {
	// We don't need to start the server, we just need it initialized to have access to its methods and fields
	server := NewRestServer()
	// We don't need to set method and path on the request as we are going to call directly the specific handler
	id := 1
	request := buildRequestForTesting(t, whateverStr, whateverStr, nil)
	requestRecorder := httptest.NewRecorder()

	request = mux.SetURLVars(request, map[string]string{
		idKey: strconv.Itoa(id),
	})
	server.deleteBlogPost(requestRecorder, request)

	assert.Equal(t, http.StatusAccepted, requestRecorder.Code)
}

func TestDeleteBlogPostOffline_nonExisting(t *testing.T) {
	// We don't need to start the server, we just need it initialized to have access to its methods and fields
	server := NewRestServer()
	// We don't need to set method and path on the request as we are going to call directly the specific handler
	id := 3
	request := buildRequestForTesting(t, whateverStr, whateverStr, nil)
	requestRecorder := httptest.NewRecorder()

	request = mux.SetURLVars(request, map[string]string{
		idKey: strconv.Itoa(id),
	})
	server.deleteBlogPost(requestRecorder, request)

	assert.Equal(t, http.StatusAccepted, requestRecorder.Code)
}

func TestGetRootOffline(t *testing.T) {
	// We don't need to start the server, we just need it initialized to have access to its methods and fields
	server := NewRestServer()
	// We don't need to set method and path on the request as we are going to call directly the specific handler
	request := buildRequestForTesting(t, whateverStr, whateverStr, nil)
	requestRecorder := httptest.NewRecorder()

	server.getRoot(requestRecorder, request)

	response := requestRecorder.Result()
	assert.Equal(t, http.StatusOK, requestRecorder.Code)
	assert.NotNil(t, response)
	defer closeResponseBodyForTesting(t, response.Body)
	responseBody, respBodyErr := ioutil.ReadAll(response.Body)
	assert.Nil(t, respBodyErr)
	assert.NotNil(t, responseBody)
	assert.Equal(t, rootMessage, string(responseBody))
}

func TestGetRoutesOffline(t *testing.T) {
	// We don't need to start the server, we just need it initialized to have access to its methods and fields
	server := NewRestServer()
	// We don't need to set method and path on the request as we are going to call directly the specific handler
	request := buildRequestForTesting(t, whateverStr, whateverStr, nil)
	requestRecorder := httptest.NewRecorder()

	server.getRoutes(requestRecorder, request)

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
