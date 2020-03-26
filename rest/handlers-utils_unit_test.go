// +build unit

package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildPost(t *testing.T) {
	id := "42"
	title := "some title"
	body := "some body"
	post := buildPost(id, title, body)

	assert.Equal(t, id, post.ID)
	assert.Equal(t, title, post.Title)
	assert.Equal(t, body, post.Body)
}

func TestBuildErrorResponse(t *testing.T) {
	request := "GET posts"
	errorMsg := "marshal error"
	errResp := buildErrorResponse(request, errorMsg)

	assert.Equal(t, request, errResp.Request)
	assert.Equal(t, errorMsg, errResp.Message)
}
