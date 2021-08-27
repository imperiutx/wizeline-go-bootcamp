package api

import (
	"net/http"
	db "wizeline-go-bootcamp/db/sqlc"
	"wizeline-go-bootcamp/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config util.Config
	store  db.Store
	router *echo.Echo
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter(config.ServerPort)
	return server, nil
}

func (server *Server) setupRouter(port string) {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.POST("/users", server.createUser)
	e.GET("/users/:id", server.getUser)

	e.Logger.Fatal(e.Start(port))
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}

func (server *Server) Start(port string) error {
	return server.router.Start(port)
}
