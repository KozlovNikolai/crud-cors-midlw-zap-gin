package server

import (
	"net/http"
	"time"

	"github.com/KozlovNikolai/crud-cors-midlw-zap-gin/config"

	"github.com/KozlovNikolai/crud-cors-midlw-zap-gin/handlers"
	"github.com/KozlovNikolai/crud-cors-midlw-zap-gin/middlewares"
	"github.com/KozlovNikolai/crud-cors-midlw-zap-gin/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	router *gin.Engine
	logger *zap.Logger
	repo   repository.EmployerRepository
}

func NewServer(repoType string, connStr string) *Server {
	// Инициализация логгера Zap
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	var repo repository.EmployerRepository

	// Выбор репозитория
	switch repoType {
	case "memory":
		repo = repository.NewInMemoryEmployerRepository()
	case "postgres":
		config.InitDB(connStr)
		repo = repository.NewPostgresEmployerRepository(config.DB)
	default:
		logger.Fatal("Invalid repository type")
	}

	// Создание сервера
	server := &Server{
		router: gin.Default(),
		logger: logger,
		repo:   repo,
	}

	// Middleware
	server.router.Use(middlewares.LoggerMiddleware(logger))
	server.router.Use(middlewares.RequestIDMiddleware())

	// CORS
	server.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8085", "https://google.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "X-Request-ID"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Инициализация обработчиков
	employerHandler := handlers.NewEmployerHandler(logger, repo)

	// CRUD маршруты для Employers
	server.router.POST("/employers", employerHandler.CreateEmployer)
	server.router.GET("/employers/:id", employerHandler.GetEmployer)
	server.router.GET("/employers", employerHandler.GetAllEmployers)
	server.router.PUT("/employers/:id", employerHandler.UpdateEmployer)
	server.router.DELETE("/employers/:id", employerHandler.DeleteEmployer)

	return server
}

func (s *Server) Run() {
	defer s.logger.Sync() // flushes buffer, if any

	// Настройка сервера с таймаутами
	server := &http.Server{
		Addr:         ":8080",
		Handler:      s.router,
		ReadTimeout:  14 * time.Second,
		WriteTimeout: 14 * time.Second,
		IdleTimeout:  63 * time.Second,
	}

	// Запуск сервера
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Fatal("Could not listen on :8080", zap.Error(err))
	}
}
