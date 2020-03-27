// +build unit

package blogpost

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitBlogPost(t *testing.T) {
	posts := initBlogPosts()

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
		blogPosts: []*blogPost{
			{ID: "1", Title: "My first BlogPost", Body: "This is the content of my first BlogPost"},
			{ID: "2", Title: "My second BlogPost", Body: "This is the content of my second BlogPost"},
		},
	}
	server.setupRouter()

	assert.NotNil(t, server.router)
}

func TestSetupHttpServer(t *testing.T) {
	cfg := loadConfig()
	server := &Server{
		config: cfg,
		blogPosts: []*blogPost{
			{ID: "1", Title: "My first BlogPost", Body: "This is the content of my first BlogPost"},
			{ID: "2", Title: "My second BlogPost", Body: "This is the content of my second BlogPost"},
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
