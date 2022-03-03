package transport

import (
	"encoding/json"
	"home/leonid/Git/Pract/telegram_bot/pkg/models"
	"home/leonid/Git/Pract/telegram_bot/pkg/service"
	"io/ioutil"
	"net/http"
	"strconv"

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
	a.echo.GET("/getusers", a.GetUsers)
	a.echo.GET("/get_user_request", a.GetUserRequests)
	a.echo.PATCH("/delete_request", a.DeleteRequst)
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

func (a *api) GetUserRequests(c echo.Context) error {

	body := c.Request().Body

	buf, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	r := new(models.Request)

	err = json.Unmarshal(buf, r)
	if err != nil {
		return err
	}

	data, err := a.svc.GetUserRequests(r.UserName)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)

}

func (a *api) GetUsers(c echo.Context) error {
	data, err := a.svc.GetUsers()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)

}

func (a *api) DeleteRequst(c echo.Context) error {
	id := c.QueryParam("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	err = a.svc.DeleteRequst(ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Ok")
}
