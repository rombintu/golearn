package server

import "github.com/gin-gonic/gin"

func (s *Server) VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			respondWithError(c, 401, "API token required")
			return
		}

		maker, err := NewJWTMaker(s.Config.Server.Secret)
		if err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		if _, err := maker.VerifyToken(token); err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		c.Next()
	}
}

func (s *Server) VerifyTokenAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			respondWithError(c, 401, "API token required")
			return
		}

		maker, err := NewJWTMaker(s.Config.Server.Secret)
		if err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		payload, err := maker.VerifyToken(token)
		if err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		if payload.Username != "admin" {
			respondWithError(c, 401, "you are not an administrator")
			return
		}
		c.Next()
	}
}
