package server

import (
	"github.com/DiscoreMe/magic-world/world"
	"github.com/labstack/echo"
	"net/http"
)

func (s *Server) ZoneCells(c echo.Context) error {
	data, err := s.w.ExportToJSON()
	if err != nil {
		return err
	}
	return c.HTMLBlob(http.StatusOK, data)
}

func (s *Server) ZoneTypes(c echo.Context) error {
	return c.JSON(http.StatusOK, world.ZoneTypeNames)
}
