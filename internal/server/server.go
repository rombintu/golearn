package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/golearn/config"
	"github.com/rombintu/golearn/store"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Logger  *logrus.Logger
	LogFile *os.File
	Config  *config.Config
	Router  *gin.Engine
	Store   *store.Store
}

// Return Server
func NewApp(config *config.Config) *Server {
	return &Server{
		Config: config,
		Logger: logrus.New(),
		Router: gin.Default(),
	}
}

// Init
func (s *Server) Start() error {
	s.Logger.Debug("Configure Logger")
	if err := s.OpenLogFile(); err != nil {
		return err
	}
	if err := s.ConfigureLogger(); err != nil {
		return err
	}
	s.Logger.Debug("Configure Router")
	s.ConfigureRouter()

	s.Logger.Debug("Configure store")
	if err := s.ConfigureStore(); err != nil {
		return err
	}

	s.Logger.Info(fmt.Sprintf(
		"Starting API server on http://%s%s",
		s.Config.Server.Host,
		s.Config.Server.Port,
	),
	)

	return http.ListenAndServe(
		s.Config.Server.Port,
		s.Router,
	)
}

func (s *Server) OpenLogFile() error {
	f, err := os.OpenFile(s.Config.Default.LogFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	s.LogFile = f
	return nil
}

func (s *Server) CloseLogFile() error {
	return s.LogFile.Close()
}

// Configure Logger, set value from Config file
func (s *Server) ConfigureLogger() error {
	level, err := logrus.ParseLevel(s.Config.Default.LogLevel)
	if err != nil {
		return err
	}
	s.Logger.SetOutput(s.LogFile)
	s.Logger.SetLevel(level)

	return nil
}

// Add routes
func (s *Server) ConfigureRouter() {
	// Test connect
	s.Router.GET("/ping", s.Ping())

	// Create user (registration)
	s.Router.POST("/user", s.CreateUser())

	// Get token
	s.Router.POST("/auth", s.Auth())

	// Middleware: req token (user)
	s.Router.Use(s.VerifyToken())

	// Middleware: req token (admin)
	s.Router.Use(s.VerifyTokenAdmin())

	// require: id INT, [type STR]
	s.Router.GET("/user", s.GetUserByID())

	// Create workers && teachers
	s.Router.POST("/worker", s.CreateWorker())
	s.Router.POST("/teacher", s.CreateTeacher())
	s.Router.POST("/group", s.CreateStudentGroup())

}

// Configure db, from Config file
func (s *Server) ConfigureStore() error {
	s.Store = &store.Store{
		Config: &s.Config.Postgres,
	}
	s.Store.Open()
	s.Store.Database.AutoMigrate(
		&store.User{},
		&store.Worker{},
		&store.Teacher{},
		&store.Declaration{},
		&store.Course{},
		&store.Journal{},
		&store.Contract{},
		&store.Payment{},
		&store.Services{},
		&store.Refferal{},
		&store.Group{},
		&store.StudentGroup{},
	)
	return nil
}
