package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/golearn/store"
)

func (s *Server) CreateCourse() gin.HandlerFunc {
	return func(c *gin.Context) {

		var crs store.Course

		if err := c.BindJSON(&crs); err != nil {
			s.respondWithError(c, 400, "not created")
			return
		}

		if crs.Title == "" {
			s.respondWithError(c, 400, "Some fields is empty")
			return
		}

		crs, err := s.Store.CreateCourse(crs)
		if err != nil {
			s.respondWithError(c, 400, err.Error())
			return
		}

		c.JSON(http.StatusCreated, crs)
	}
}

func (s *Server) GetAllCourses() gin.HandlerFunc {
	return func(c *gin.Context) {

		crsAll, err := s.Store.GetAllCourses()
		if err != nil {
			s.respondWithError(c, 400, err.Error())
			return
		}

		c.JSON(http.StatusCreated, crsAll)
	}
}
