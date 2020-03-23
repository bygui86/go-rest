package rest

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func buildRequestForTesting(t *testing.T, method, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func closeResponseBodyForTesting(t *testing.T, body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func checkGetPostsResponseForTesting(t *testing.T, response *http.Response, respBodyErr error, server *Server) {
	defer closeResponseBodyForTesting(t, response.Body)
	responseBody, respBodyErr := ioutil.ReadAll(response.Body)
	assert.Nil(t, respBodyErr)
	assert.NotNil(t, responseBody)

	var posts []*post
	unmarshErr := json.Unmarshal(responseBody, &posts)
	assert.Nil(t, unmarshErr)
	assert.NotNil(t, posts)
	assert.Equal(t, len(server.posts), len(posts))
}
