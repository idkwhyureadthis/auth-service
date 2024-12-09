package handler

import (
	"github.com/idkwhyureadthis/auth-service/pkg/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	e *echo.Echo
	s *service.Service
}

func New(connUrl string, secret []byte) *Handler {
	h := &Handler{
		e: echo.New(),
		s: service.New(connUrl, secret),
	}
	return h
}

func (h *Handler) Start(port string) {
	h.e.Logger.Fatal(h.e.Start(":" + port))
}
