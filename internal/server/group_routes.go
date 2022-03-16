package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/golearn/store"
)

func (s *Server) CreateGroup() gin.HandlerFunc {
	return func(c *gin.Context) {

		var g store.Group

		if err := c.BindJSON(&g); err != nil {
			s.respondWithError(c, 400, "not created")
			return
		}

		title := g.Title

		if title == "" {
			s.respondWithError(c, 400, "Some fields is empty")
			return
		}

		if err := s.Store.CreateGroup(g); err != nil {
			s.respondWithError(c, 400, err.Error())
			return
		}

		c.JSON(http.StatusCreated, g)
	}
}

func (s *Server) DeleteGroupByTitleOrID() gin.HandlerFunc {
	return func(c *gin.Context) {

		title := c.Query("title")
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			s.respondWithError(c, 400, err.Error())
			return
		}

		if title == "" || id == 0 {
			s.respondWithError(c, 400, "Some fields is empty")
			return
		}

		if title != "" && id != 0 {
			s.respondWithError(c, 400, "you have to choose one of the options")
			return
		}

		if title != "" && id == 0 {
			if err := s.Store.DeleteGroupByTitle(title); err != nil {
				s.respondWithError(c, 400, err.Error())
				return
			}
		} else if title == "" && id != 0 {
			if err := s.Store.DeleteGroupByID(id); err != nil {
				s.respondWithError(c, 400, err.Error())
				return
			}
		}

		c.JSON(http.StatusCreated, store.Message{Content: "Delete"})
	}
}

func (s *Server) JoinUserGroup() gin.HandlerFunc {
	return func(c *gin.Context) {

		userLogin := c.Query("user_login")
		groupName := c.Query("group_name")

		if userLogin == "" || groupName == "" {
			s.respondWithError(c, 400, "Some fields is empty")
			return
		}

		if err := s.Store.JoinUserGroup(userLogin, groupName); err != nil {
			s.respondWithError(c, 400, err.Error())
			return
		}

		c.JSON(http.StatusCreated, store.Message{Content: "User joined"})
	}
}
