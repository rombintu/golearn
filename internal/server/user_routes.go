package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/golearn/store"
)

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

func (s *Server) CreateWorker() gin.HandlerFunc {
	return func(c *gin.Context) {

		var u store.Worker
		Ok := store.Ping{
			Message: "worker created",
		}

		if err := c.BindJSON(&u); err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, "worker not created")
			return
		}

		account := u.Account
		password := u.Password
		role := u.Role

		if role == "" {
			u.Role = "user"
		}

		if account == "" || password == "" {
			s.Logger.Error("Some worker fields is empty")
			respondWithError(c, 401, "Some worker fields is empty")
			return
		}

		if err := s.Store.CreateWorker(u); err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		c.JSON(http.StatusCreated, Ok)
	}
}

func (s *Server) CreateTeacher() gin.HandlerFunc {
	return func(c *gin.Context) {

		var u store.Teacher
		Ok := store.Ping{
			Message: "teacher created",
		}

		if err := c.BindJSON(&u); err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, "teacher not created")
			return
		}

		account := u.Account
		password := u.Password
		role := u.Role

		if role == "" {
			u.Role = "user"
		}

		if account == "" || password == "" {
			s.Logger.Error("Some teacher fields is empty")
			respondWithError(c, 401, "Some teacher fields is empty")
			return
		}

		if err := s.Store.CreateTeacher(u); err != nil {
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
		typeUser := c.Query("type")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			s.Logger.Error(err)
			return
		}

		var user interface{}

		switch typeUser {
		case "worker":
			user, err = s.Store.GetWorkerByID(id)
		case "teacher":
			user, err = s.Store.GetTeacherByID(id)
		default:
			user, err = s.Store.GetUserByID(id)
		}

		if err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}
