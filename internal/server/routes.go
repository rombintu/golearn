package server

import (
	"net/http"
	"strconv"

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

// Create new user if not exists with new wallet
func (s *Server) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var u store.User
		Ok := store.Ping{
			Message: "user created",
		}

		if err := c.BindJSON(&u); err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, "user not created")
			return
		}

		account := u.Account
		password := u.Password
		role := u.Role

		if role == "" {
			u.Role = "user"
		}

		if account == "" || password == "" {
			s.Logger.Error("Some user fields is empty")
			respondWithError(c, 401, "Some user fields is empty")
			return
		}

		if err := s.Store.CreateUser(u); err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		c.JSON(http.StatusCreated, Ok)
	}
}

func (s *Server) GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {

		idStr := c.Query("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			s.Logger.Error(err)
			return
		}
		user, err := s.Store.GetUserByID(id)
		if err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		c.JSON(http.StatusCreated, user)
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
