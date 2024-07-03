package infra

import (
	"context"
	"iroly/app/adapter"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	slogecho "github.com/samber/slog-echo"
)

type Handler struct {
	echoLambda *echoadapter.EchoLambdaALB
}

func NewHandler() *Handler {
	e := echo.New()

	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
	e.Use(slogecho.New(slog.New(slog.NewJSONHandler(os.Stdout, nil)))) // access log
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	basepath := "/" + strings.Trim(os.Getenv("BASE_PATH"), "/")
	e.Pre(middleware.Rewrite(map[string]string{
		basepath + "/*": "/$1",
	}))

	ctr := adapter.NewController()

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{})
	})
	e.GET("/hello", func(c echo.Context) error {
		return ctr.Hello(c)
	})

	return &Handler{
		echoLambda: echoadapter.NewALB(e),
	}
}

func (h *Handler) HandleRequest(ctx context.Context, req events.ALBTargetGroupRequest) (events.ALBTargetGroupResponse, error) {
	r, err := h.echoLambda.ProxyWithContext(ctx, req)
	return r, err
}

func (h *Handler) DebugRunEcho() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	e := h.echoLambda.Echo
	e.Logger.Fatal(e.Start(":" + port))
}
