package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/help", help)
	e.GET("/event", getEvent)
	e.POST("/event", updateEvent)
	e.Logger.Fatal(e.Start(":1323"))
}

type Event struct {
	Id string `json:"id"`
	Hits int `json:"hits"`
	Current int `json:"current"`
	Landing string `json:"landing"`
}

func getEvent(c echo.Context) error {
	id := c.QueryParam("id")
	e := new(Event)
	e.Id = id
	return c.JSON(http.StatusOK, e)
}

func updateEvent(c echo.Context) error {
	id := c.QueryParam("id")
	e := new(Event)
	e.Id = id
	return c.JSON(http.StatusOK, e)
}

func help(c echo.Context) error {
	return c.HTML(http.StatusOK, "help")
}
