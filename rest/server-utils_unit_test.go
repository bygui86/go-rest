// +build unit

package rest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitPost(t *testing.T) {
	posts := initPosts()

	assert.NotNil(t, posts)
	assert.Equal(t, 2, len(posts))

	expectedPost1 := posts[0]
	assert.Equal(t, expectedPost1.ID, posts[0].ID)
	assert.Equal(t, expectedPost1.Title, posts[0].Title)
	assert.Equal(t, expectedPost1.Body, posts[0].Body)
	expectedPost2 := posts[1]
	assert.Equal(t, expectedPost2.ID, posts[1].ID)
	assert.Equal(t, expectedPost2.Title, posts[1].Title)
	assert.Equal(t, expectedPost2.Body, posts[1].Body)
}

func TestSetupRouter(t *testing.T) {
	cfg := loadConfig()
	server := &Server{
		config: cfg,
		posts: []*post{
			{ID: "1", Title: "My first post", Body: "This is the content of my first post"},
			{ID: "2", Title: "My second post", Body: "This is the content of my second post"},
		},
	}
	server.setupRouter()

	assert.NotNil(t, server.router)
}

func TestSetupHttpServer(t *testing.T) {
	cfg := loadConfig()
	server := &Server{
		config: cfg,
		posts: []*post{
			{ID: "1", Title: "My first post", Body: "This is the content of my first post"},
			{ID: "2", Title: "My second post", Body: "This is the content of my second post"},
		},
	}
	server.setupRouter()
	server.setupHTTPServer()

	assert.NotNil(t, server.httpServer)
	assert.Equal(t,
		fmt.Sprintf(httpServerHostFormat, server.config.RestHost, server.config.RestPort), server.httpServer.Addr)
	assert.Equal(t, server.router, server.httpServer.Handler)
	assert.Equal(t, httpServerWriteTimeoutDefault, server.httpServer.WriteTimeout)
	assert.Equal(t, httpServerReadTimeoutDefault, server.httpServer.ReadTimeout)
	assert.Equal(t, httpServerIdelTimeoutDefault, server.httpServer.IdleTimeout)
}
