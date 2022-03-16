package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/golearn/store"
)

func (s *Server) CreateDeclaration() gin.HandlerFunc {
	return func(c *gin.Context) {

		var d store.Declaration

		if err := c.BindJSON(&d); err != nil {
			s.respondWithError(c, 400, err.Error())
			return
		}

		title := d.Title
		userID := d.UserID

		if title == "" || userID == 0 {
			s.respondWithError(c, 400, "Some fields is empty")
			return
		}

		if err := s.Store.CreateDeclaration(d); err != nil {
			s.respondWithError(c, 400, err.Error())
			return
		}

		c.JSON(http.StatusCreated, d)
	}
}

func (s *Server) DeleteDeclarationByUserIDAndTitle() gin.HandlerFunc {
	return func(c *gin.Context) {

		title := c.Query("title")
		userID, err := strconv.Atoi(c.Query("user_id"))
		if err != nil {
			s.respondWithError(c, 400, err.Error())
			return
		}

		if userID == 0 {
			s.respondWithError(c, 400, "Some fields is empty")
			return
		}

		if err := s.Store.DeleteDeclaration(userID, title); err != nil {
			s.respondWithError(c, 400, err.Error())
			return
		}

		c.JSON(http.StatusOK, store.Message{Content: "Delete"})
	}
}

func (s *Server) GetDeclarationByUserID() gin.HandlerFunc {
	return func(c *gin.Context) {

		userID := c.Query("user_id")

		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			s.respondWithError(c, 400, err.Error())
			return
		}

		declaration, err := s.Store.GetDeclarationByUserID(userIDInt)
		if err != nil {
			s.respondWithError(c, 400, err.Error())
			return
		}

		c.JSON(http.StatusCreated, declaration)
	}
}
