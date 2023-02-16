package pkg

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Server struct {
	r *echo.Echo
	l *zap.SugaredLogger
}

func NewServer(l *zap.SugaredLogger, r *echo.Echo) *Server {
	return &Server{
		r: r,
		l: l,
	}
}

func (s *Server) StartServer() {
	go func() {
		s.r.Logger.Fatal(s.r.Start(":1323"))
	}()

	s.l.Infoln("Server is started")
}


func(s *Server) Routes() {
	s.r.GET("/forms", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}