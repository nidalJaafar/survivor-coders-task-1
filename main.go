package main

import (
	"github.com/labstack/echo"
	"student/config"
	"student/route"
)

var e = echo.New()

func main() {
	config.DbConfig()
	route.Routes(e)
	err := e.Start("localhost:8080")
	if err != nil {
		return
	}
}
