package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/golearn/store"
)

func (s *Server) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var u store.User

		if err := c.BindJSON(&u); err != nil {
			s.respondWithError(c, 401, "not created")
			return
		}

		account := u.Account
		password := u.Password
		role := u.Role

		if role == "" {
			u.Role = "user"
		}

		if account == "" || password == "" {
			s.respondWithError(c, 401, "Some fields is empty")
			return
		}

		userNew, err := s.Store.CreateUser(u)
		if err != nil {
			s.respondWithError(c, 401, err.Error())
			return
		}

		c.JSON(http.StatusCreated, userNew)
	}
}

func (s *Server) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var u store.User

		if err := c.BindJSON(&u); err != nil {
			s.respondWithError(c, 401, "not updated")
			return
		}

		userNew, err := s.Store.UpdateUser(u)
		if err != nil {
			s.respondWithError(c, 401, err.Error())
			return
		}

		c.JSON(http.StatusCreated, userNew)
	}
}

func (s *Server) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		idStr := c.Query("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			s.respondWithError(c, 401, err.Error())
			return
		}

		if err := s.Store.DeleteUser(id); err != nil {
			s.respondWithError(c, 401, err.Error())
			return
		}

		c.JSON(http.StatusCreated, store.Message{Content: fmt.Sprintf("user %s - deleted", idStr)})
	}
}

func (s *Server) CreateWorker() gin.HandlerFunc {
	return func(c *gin.Context) {

		var u store.Worker

		if err := c.BindJSON(&u); err != nil {
			s.respondWithError(c, 401, "not created")
			return
		}

		account := u.Account
		password := u.Password
		role := u.Role

		if role == "" {
			u.Role = "user"
		}

		if account == "" || password == "" {
			s.respondWithError(c, 401, "Some fields is empty")
			return
		}

		if err := s.Store.CreateWorker(u); err != nil {
			s.respondWithError(c, 401, err.Error())
			return
		}

		c.JSON(http.StatusCreated, u)
	}
}

func (s *Server) CreateTeacher() gin.HandlerFunc {
	return func(c *gin.Context) {

		var u store.Teacher

		if err := c.BindJSON(&u); err != nil {
			s.respondWithError(c, 401, "not created")
			return
		}

		account := u.Account
		password := u.Password
		role := u.Role

		if role == "" {
			u.Role = "user"
		}

		if account == "" || password == "" {
			s.respondWithError(c, 401, "Some fields is empty")
			return
		}

		if err := s.Store.CreateTeacher(u); err != nil {
			s.respondWithError(c, 401, err.Error())
			return
		}

		c.JSON(http.StatusCreated, u)
	}
}

func (s *Server) GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {

		idStr := c.Query("id")
		typeUser := c.Query("type")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			s.respondWithError(c, 401, err.Error())
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
			s.respondWithError(c, 401, err.Error())
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}
