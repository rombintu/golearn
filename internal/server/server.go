package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/golearn/config"
	"github.com/rombintu/golearn/store"
	"github.com/rombintu/golearn/tools"
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

	if s.Config.Default.LogFile != "" {
		if err := s.OpenLogFile(); err != nil {
			return err
		}
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
	if s.Config.Default.LogFile != "" {
		s.Logger.SetOutput(s.LogFile)
	}
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

	// Get all courses
	s.Router.GET("/course/all", s.GetAllCourses())
	// s.Router.POST("/course/delete", s.DeleteCourse()) // TODO

	// =================== required USER auth =================== //
	// Middleware: req token (user)
	s.Router.Use(s.VerifyToken())

	// Create new declaration: {title, user_id}
	s.Router.POST("/paper/declaration/create", s.CreateDeclaration())
	// Get all declarations by user id: user_id?
	s.Router.GET("/paper/declaration/list", s.GetDeclarationByUserID())
	// Delete declaration by user id and title; if title == nil => title = *
	s.Router.POST("/paper/declaration/delete", s.DeleteDeclarationByUserIDAndTitle())

	s.Router.GET("/user", s.GetUserByID())
	s.Router.GET("/user/delete", s.DeleteUser())
	s.Router.POST("/user/update", s.UpdateUser())
	// =================== required ADMIN auth =================== //
	// Middleware: req token (admin)
	s.Router.Use(s.VerifyTokenAdmin())

	s.Router.POST("/course", s.CreateCourse())
	// require: id? [type]?

	// Create workers && teachers
	s.Router.POST("/user/worker", s.CreateWorker())
	s.Router.POST("/user/teacher", s.CreateTeacher())
	s.Router.POST("/group", s.CreateGroup())

	// Join User on Group: user_login? group_name?
	s.Router.POST("/group/join", s.JoinUserGroup())

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
		// &store.UserGroup{},
	)
	if err := s.Store.Database.SetupJoinTable(
		&store.Group{},
		"Users",
		&store.GroupUsers{},
	); err != nil {
		return err
	}
	// FOR DEV

	hashPass, err := tools.HashPassword("admin")
	if err != nil {
		return err
	}

	admin := store.User{
		Account:  "admin",
		Password: hashPass,
		Role:     "admin",
	}

	if err := s.Store.Database.FirstOrCreate(&admin).Error; err != nil {
		return err
	}
	return nil
}
