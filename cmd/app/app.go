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
<<<<<<< HEAD
		_ = db.Raw("SELECT * FROM forms")
	
=======
>>>>>>> 6e51639bc540e62dc4adfce5e5af7f57bb836381
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	s := pkg.NewServer(l, r)
	s.StartServer()
}
