package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/bygui86/go-rest/logging"
)

const (
	idKey = "id"

	routerPostsRootUrl = "/posts"
	routerIdUrlPath    = "/{" + idKey + "}"

	httpServerHostFormat          = "%s:%d"
	httpServerWriteTimeoutDefault = time.Second * 15
	httpServerReadTimeoutDefault  = time.Second * 15
	httpServerIdelTimeoutDefault  = time.Second * 60
)

// NewRestServer - Create new REST server
func NewRestServer() *Server {
	logging.Log.Debug("Create new REST server")

	cfg := loadConfig()

	server := &Server{
		config: cfg,
		posts:  initPosts(),
	}

	server.setupRouter()
	server.setupHTTPServer()
	return server
}

// setupRouter - Create new Gorilla mux router
func (s *Server) setupRouter() {
	logging.Log.Debug("Create new router")

	s.router = mux.NewRouter().StrictSlash(true)

	s.router.HandleFunc(routerPostsRootUrl, s.getPosts).Methods(http.MethodGet)
	s.router.HandleFunc(routerPostsRootUrl, s.createPost).Methods(http.MethodPost)
	s.router.HandleFunc(routerPostsRootUrl+routerIdUrlPath, s.getPost).Methods(http.MethodGet)
	s.router.HandleFunc(routerPostsRootUrl+routerIdUrlPath, s.updatePost).Methods(http.MethodPut)
	s.router.HandleFunc(routerPostsRootUrl+routerIdUrlPath, s.deletePost).Methods(http.MethodDelete)
}

// setupHTTPServer - Create new HTTP server
func (s *Server) setupHTTPServer() {
	logging.SugaredLog.Debugf("Create new HTTP server on port %d", s.config.RestPort)

	if s.config != nil {
		s.httpServer = &http.Server{
			Addr:    fmt.Sprintf(httpServerHostFormat, s.config.RestHost, s.config.RestPort),
			Handler: s.router,
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: httpServerWriteTimeoutDefault,
			ReadTimeout:  httpServerReadTimeoutDefault,
			IdleTimeout:  httpServerIdelTimeoutDefault,
		}
		return
	}

	logging.Log.Error("HTTP server creation failed: REST server configurations not initialized")
}

// Start - Start REST server
func (s *Server) Start() {
	logging.Log.Info("Start REST server")

	if s.httpServer != nil && !s.running {
		go func() {
			err := s.httpServer.ListenAndServe()
			if err != nil {
				logging.SugaredLog.Errorf("Error starting REST server: %s", err.Error())
			}
		}()
		s.running = true
		logging.SugaredLog.Infof("REST server listening on port %d", s.config.RestPort)
		return
	}

	logging.Log.Error("REST server start failed: HTTP server not initialized or HTTP server already running")
}

// Shutdown - Shutdown REST server
func (s *Server) Shutdown(timeout int) {
	logging.SugaredLog.Warnf("Shutdown REST server, timeout %d", timeout)

	if s.httpServer != nil && s.running {
		// create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
		// does not block if no connections, otherwise wait until the timeout deadline
		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			logging.SugaredLog.Errorf("Error shutting down REST server: %s", err.Error())
		}
		s.running = false
		return
	}

	logging.Log.Error("REST server shutdown failed: HTTP server not initialized or HTTP server not running")
}
