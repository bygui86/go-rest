// build unit

package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO missing createPost, getPost, updatePost, deletePost

func TestGetPosts_mock(t *testing.T) {
	server := NewRestServer()
	// INFO: We don't need to start the server, We just need it initialized to have access to its methods and fields

	testServer := httptest.NewServer(http.HandlerFunc(server.getPosts))
	defer testServer.Close()

	response, respBodyErr := http.Get(testServer.URL)
	assert.Nil(t, respBodyErr)
	assert.NotNil(t, response)

	checkGetPostsResponseForTesting(t, response, respBodyErr, server)
}
