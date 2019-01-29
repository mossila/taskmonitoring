package main

import (
	"github.com/labstack/echo"
	"github.com/mossila/taskmonitoring/server/api"
)

func main() {
	e := echo.New()
	e.GET("/", api.Greeting)
	e.Logger.Fatal(e.Start(":1323"))
}
