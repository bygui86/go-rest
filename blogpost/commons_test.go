// +build unit

package blogpost

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func buildRequestForTesting(t *testing.T, method, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatal(err)
	}
	return request
}

func buildRequestForTestingWithQuery(t *testing.T, method, url string, queryMap map[string]string) *http.Request {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	if queryMap != nil && len(queryMap) > 0 {
		query := request.URL.Query()
		for queryParamKey, queryParamValue := range queryMap {
			query.Add(queryParamKey, queryParamValue)
		}
		request.URL.RawQuery = query.Encode()
	}
	return request
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

	var posts []*blogPost
	unmarshErr := json.Unmarshal(responseBody, &posts)
	assert.Nil(t, unmarshErr)
	assert.NotNil(t, posts)
	assert.Equal(t, len(server.blogPosts), len(posts))

	expectedPost1 := server.blogPosts[0]
	assert.Equal(t, expectedPost1.ID, posts[0].ID)
	assert.Equal(t, expectedPost1.Title, posts[0].Title)
	assert.Equal(t, expectedPost1.Body, posts[0].Body)
	expectedPost2 := server.blogPosts[1]
	assert.Equal(t, expectedPost2.ID, posts[1].ID)
	assert.Equal(t, expectedPost2.Title, posts[1].Title)
	assert.Equal(t, expectedPost2.Body, posts[1].Body)
}

func checkGetPostResponseForTesting(t *testing.T, response *http.Response, expectedPost *blogPost, verifyId bool) {
	defer closeResponseBodyForTesting(t, response.Body)
	responseBody, respBodyErr := ioutil.ReadAll(response.Body)
	assert.Nil(t, respBodyErr)
	assert.NotNil(t, responseBody)

	var post *blogPost
	unmarshErr := json.Unmarshal(responseBody, &post)
	assert.Nil(t, unmarshErr)
	assert.NotNil(t, post)
	if verifyId {
		assert.Equal(t, expectedPost.ID, post.ID)
	}
	assert.Equal(t, expectedPost.Title, post.Title)
	assert.Equal(t, expectedPost.Body, post.Body)
}

func checkErrorResponseForTesting(t *testing.T, response *http.Response, expectedErrorResponse *errorResponse) {
	defer closeResponseBodyForTesting(t, response.Body)
	responseBody, respBodyErr := ioutil.ReadAll(response.Body)
	assert.Nil(t, respBodyErr)
	assert.NotNil(t, responseBody)

	var errorResponse *errorResponse
	unmarshErr := json.Unmarshal(responseBody, &errorResponse)
	assert.Nil(t, unmarshErr)
	assert.NotNil(t, errorResponse)
	assert.Equal(t, expectedErrorResponse.Request, errorResponse.Request)
	assert.Equal(t, expectedErrorResponse.Message, errorResponse.Message)
}
