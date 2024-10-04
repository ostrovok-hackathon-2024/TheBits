package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Config struct {
}

type ServerApi struct {
	config *Config
}

func NewServerAPI(cfg *Config) *ServerApi {
	return &ServerApi{
		config: cfg,
	}
}
func (s *ServerApi) Run(ctx context.Context) error {
	fmt.Println("Starting server api...")

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			answer string
		}{
			answer: "Hello, front!",
		})
	})
	e.POST("/hotels", paramsHandler)
	e.Logger.Fatal(e.Start(":8088"))
	return nil
}
