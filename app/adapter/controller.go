package adapter

import (
	"iroly/app/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (ctr *Controller) Hello(c echo.Context) error {
	hello := domain.NewGreeter().Hello()
	return c.JSON(http.StatusOK, hello)
}
