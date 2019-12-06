package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/Hello", HelloController)

	e.Start(":8000")
}

//HelloController fungsi control hello
func HelloController(c echo.Context) error {
	//ganti hello world dengan print json
	return c.String(http.StatusOK, "Hello World")

}
