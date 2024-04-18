package server

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/gin-gonic/gin"
)

type mockResponse struct {
	HTTPCode int
	Response map[string]interface{}
}

// Server is a mock conduit server.
type Server struct {
	engine    *gin.Engine
	server    *httptest.Server
	mu        sync.Mutex
	responses map[string][]mockResponse
}

// New creates a new empty conduit server.
func New() *Server {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	ts := httptest.NewServer(r)

	return &Server{
		engine:    r,
		server:    ts,
		responses: make(map[string][]mockResponse),
	}
}

// RegisterCapabilities adds a default handler for the
// `conduit.getcapabilities` API endpoint.
func (s *Server) RegisterCapabilities() {
	s.RegisterMethod("conduit.getcapabilities", 200, gin.H{
		"result": map[string][]string{
			"authentication": {"token", "session"},
			"signatures":     {"consign"},
			"input":          {"json", "urlencoded"},
			"output":         {"json"},
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

// RegisterMethodResponse adds a response to return from the conduit API method.
// Call this multiple times to register multiple responses from the API server.
func (s *Server) RegisterMethodResponse(
	method string,
	httpCode int,
	response map[string]interface{},
) {
	s.mu.Lock()
	defer s.mu.Unlock()

	path := fmt.Sprintf("/api/%s", method)
	resp := mockResponse{
		HTTPCode: httpCode,
		Response: response,
	}
	if responses, exists := s.responses[path]; exists {
		s.responses[path] = append(responses, resp)
	} else {
		var responses []mockResponse
		responses = append(responses, resp)
		s.responses[path] = responses

		// create handler
		s.engine.POST(path, func(c *gin.Context) {
			s.mu.Lock()
			defer s.mu.Unlock()
			if responses, exists := s.responses[path]; exists {
				if len(responses) == 0 {
					log.Printf("no mock responses left for path %s", path)
					c.JSON(http.StatusInternalServerError, nil)
					return
				}

				mockResponse, rest := responses[0], responses[1:]
				c.JSON(mockResponse.HTTPCode, mockResponse.Response)
				s.responses[path] = rest
				return
			}

			log.Printf("no mock responses defined for path %s", path)
			c.JSON(http.StatusInternalServerError, nil)
			return
		})
	}
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
