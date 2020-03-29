package server

import (
	"github.com/DiscoreMe/magic-world/world"
	"github.com/labstack/echo"
)

type Server struct {
	w *world.World
}

func NewServer(w *world.World) *Server {
	return &Server{w: w}
}

func (s *Server) Listen(endpoint string) error {
	e := echo.New()

	zone := e.Group("zone")
	zone.GET("/cells", s.ZoneCells)

	return e.Start(endpoint)
}
