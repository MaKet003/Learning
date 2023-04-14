package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func save(c echo.Context) error {
	//Get name
	name := c.FormValue("name")
	//Get avater
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	//Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	//Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	//Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}

func main() {
	e := echo.New()
	//e.GET("/show", show)
	e.POST("/save", save)
	e.Logger.Fatal(e.Start(":1323"))
}
