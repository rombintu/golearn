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

		crsS, err := s.Store.CreateCourse(crs)
		if err != nil {
			s.respondWithError(c, 400, err.Error())
			return
		}

		c.JSON(http.StatusCreated, crsS)
	}
}

func (s *Server) UploadCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			s.respondWithError(c, 400, err.Error())
			return
		}
		if err := c.SaveUploadedFile(file, "store/upload/"+file.Filename); err != nil {
			s.respondWithError(c, 400, err.Error())
			return
		}
		c.JSON(http.StatusOK, store.Ping{Message: "Success"})
	}
}

func (s *Server) DownloadCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileName, err := c.GetQuery("filename")
		if !err {
			s.respondWithError(c, 400, "empty filename")
			return
		}
		c.File("store/upload/" + fileName)
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
