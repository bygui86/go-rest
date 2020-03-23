// build unit

package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO missing createPost, getPost, updatePost, deletePost

func TestGetPosts_offline(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	request := buildRequestForTesting(t, http.MethodGet, routerPostsRootUrl, nil)
	requestRecorder := httptest.NewRecorder()
	server.getPosts(requestRecorder, request)

	response := requestRecorder.Result()
	assert.NotNil(t, response)
	assert.Equal(t, contentTypeHeaderValue, response.Header.Get(contentTypeHeaderKey))

	checkGetPostsResponseForTesting(t, response, nil, server)
}
