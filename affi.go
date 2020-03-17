package main

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func init() {
	log.Println("init")
	e.GET("/affi", affi)
}

func affi(c echo.Context) error {
	referer := c.Request().Header.Get("referer")
	return c.String(http.StatusOK, "referer: "+referer)
}
