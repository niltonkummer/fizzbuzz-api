package http

import (
	"context"
	"net"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	app     *echo.Echo
	handler *Handler
}

func NewRouter(ctx context.Context) *Router {

	app := echo.New()
	app.Validator = NewValidator()
	app.Server.BaseContext = func(_ net.Listener) context.Context {
		return ctx
	}

	return &Router{
		app: app,
	}
}

// RegisterRoutes registers the HTTP routes for the application
func (r *Router) RegisterRoutes(handler *Handler) {

	r.handler = handler

	r.app.Use(middleware.Logger())

	r.app.POST("/fizzbuzz", handler.HandleFizzBuzzRequest)
	r.app.GET("/stats", handler.HandleGetStats)
}

func (r *Router) GetHandler() *Handler {
	return r.handler
}

// Start starts the HTTP server
func (r *Router) Start(addr string) error {

	return r.app.Start(addr)
}

// Shutdown gracefully shuts down the HTTP server
func (r *Router) Shutdown(ctx context.Context) error {
	return r.app.Shutdown(ctx)
}
