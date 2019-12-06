package main

import(
	"net/http"
	"github.com/labstack/echo"
)
func main(){
	e:= echo.New()

	e.Get("/", HelloController)

	e.Start(":8000")
}

//HelloController fungsi control hello
func HelloController (e echo.Context) error{
	return c.String(http.StatusOK, "Hello World")

}
