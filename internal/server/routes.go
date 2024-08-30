package server

import (
	"Fanduel_Gen/cmd/web"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/assets/*", echo.WrapHandler(fileServer))

	e.GET("/web", echo.WrapHandler(templ.Handler(web.HelloForm())))
	e.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler)))

	apiGroup := e.Group("/api")
	firstVersion := apiGroup.Group("/v1")

	// Next Gen Stats
	statsGroup := firstVersion.Group("/stats")
	statsGroup.GET("/next_gen_passing", s.NextGenPassingHandler)
	statsGroup.GET("/next_gen_rushing", s.NextGenRushingHandler)

	// procections
	projections := firstVersion.Group("/projections")
	projections.GET("/rotowire", s.RotoWireHandler)

	return e
}
