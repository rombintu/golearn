package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/golearn/store"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

// Test func, return 200 and {"message" : "pong"}
func (s *Server) Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		pong := store.Ping{Message: "pong"}
		c.JSON(http.StatusOK, pong)
	}
}

func (s *Server) CreateStudentGroup() gin.HandlerFunc {
	return func(c *gin.Context) {

		var g store.StudentGroup
		Ok := store.Ping{
			Message: "group created",
		}

		if err := c.BindJSON(&g); err != nil {
			s.Logger.Debug(err)
			respondWithError(c, 400, "group not created")
			return
		}

		title := g.Group.Title

		if title == "" {
			s.Logger.Debug("Some group fields is empty")
			respondWithError(c, 400, "Some group fields is empty")
			return
		}

		if err := s.Store.CreateStudentGroup(g); err != nil {
			s.Logger.Debug(err)
			respondWithError(c, 400, err.Error())
			return
		}

		c.JSON(http.StatusCreated, Ok)
	}
}
