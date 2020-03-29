package server

import (
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
