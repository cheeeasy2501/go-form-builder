package app

import (
	"context"
	"net/http"

	"github.com/cheeeasy2501/go-form-builder/config"
	"github.com/cheeeasy2501/go-form-builder/pkg"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"go.uber.org/zap"
)

func Run(ctx context.Context, l *zap.SugaredLogger, c *config.Config, db *gorm.DB) {
	r := echo.New()

	r.GET("/forms", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	s := pkg.NewServer(l, r)
	s.StartServer()
}
