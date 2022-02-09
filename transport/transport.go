package transport

import (
	"home/leonid/Git/Pract/telegram_bot/pkg/service"
	"net/http"

	"github.com/labstack/echo"
)

type api struct {
	echo *echo.Echo
	svc  service.Service
	port string
}

func NewServerConnect(port string, svc service.Service) *api {
	e := echo.New()

	return &api{
		echo: e,
		svc:  svc,
		port: port,
	}
}

func (a *api) InitEndpoints() {

	a.echo.GET("/getrequest", a.GetRequest)
}

func (a *api) StartServer() error {

	return a.echo.Start(":" + a.port)
}

func (a *api) GetRequest(c echo.Context) error {
	data, err := a.svc.GetRequest()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)

}
