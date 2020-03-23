// +build unit

package rest

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewRestServer(t *testing.T) {
	server := NewRestServer()

	assert.NotNil(t, server)
	assert.NotNil(t, server.config)
	assert.NotNil(t, server.router)
	assert.NotNil(t, server.httpServer)
	assert.NotNil(t, server.posts)
	assert.Equal(t, false, server.running)
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

func TestStart(t *testing.T) {
	server := NewRestServer()
	server.Start()

	assert.Equal(t, true, server.running)

	server.Shutdown(1)
}

func TestShutdown(t *testing.T) {
	server := NewRestServer()
	server.Start()
	time.Sleep(200 * time.Millisecond)
	server.Shutdown(1)

	assert.Equal(t, false, server.running)
}
