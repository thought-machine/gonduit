package server

import (
	"fmt"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// Server is a mock conduit server.
type Server struct {
	engine *gin.Engine
	server *httptest.Server
}

// New creates a new empty conduit server.
func New() *Server {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	ts := httptest.NewServer(r)

	return &Server{
		engine: r,
		server: ts,
	}
}

// RegisterCapabilities adds a default handler for the
// `conduit.getcapabilities` API endpoint.
func (s *Server) RegisterCapabilities() {
	s.RegisterMethod("conduit.getcapabilities", 200, gin.H{
		"result": map[string][]string{
			"authentication": []string{"token", "session"},
			"signatures":     []string{"consign"},
			"input":          []string{"json", "urlencoded"},
			"output":         []string{"json"},
		},
	})
}

// RegisterMethod adds a handler for a specific conduit API method.
func (s *Server) RegisterMethod(
	method string,
	httpCode int,
	response map[string]interface{},
) {
	s.engine.POST(fmt.Sprintf("/api/%s", method), func(c *gin.Context) {
		c.JSON(httpCode, response)
	})
}

// GetURL returns the URL of the root of the server.
func (s *Server) GetURL() string {
	return s.server.URL
}

// Close shuts down the server. This should be called at the end of every test
// or by using defer.
func (s *Server) Close() {
	s.server.Close()
}
