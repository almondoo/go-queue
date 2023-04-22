package main

import (
	"context"
	"go-queue/internal/route"
	"go-queue/internal/validator"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var ctx = context.Background()

func main() {
	// echo instance
	e := echo.New()
	// validator
	e.Validator = validator.NewValidator()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routing
	route.InitRouting(e, ctx)

	// setting server
	startServer(e)
}

func startServer(e *echo.Echo) {
	h2s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}
	s := http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(e, h2s),
	}
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
