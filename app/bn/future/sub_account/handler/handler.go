package handler

import (
	"github.com/labstack/echo/v4"
)

type IHandler[Req any] interface {
	Handler(c echo.Context) error
	GetReqBody(c echo.Context) (*Req, error)
}
