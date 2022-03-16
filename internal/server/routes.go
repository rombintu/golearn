package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/golearn/store"
)

func (s *Server) respondWithError(c *gin.Context, code int, message interface{}) {
	s.Logger.Error(message)
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

// Test func, return 200 and {"message" : "pong"}
func (s *Server) Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		pong := store.Ping{Message: "pong"}
		c.JSON(http.StatusOK, pong)
	}
}
