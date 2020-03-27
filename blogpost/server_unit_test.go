// +build unit

package blogpost

import (
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
	assert.NotNil(t, server.blogPosts)
	assert.Equal(t, false, server.running)
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
