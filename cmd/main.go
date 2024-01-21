package main

import (
	"html/template"

	"github.com/labstack/echo/v4"
)

type Data struct {
	PageTitle string
}

func main() {
	var data = Data{PageTitle: "Hello Echo 2"}
	var tmpl = template.Must(template.ParseFiles("views/hello.html"))

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		tmpl.Execute(c.Response().Writer, data)
		return nil
	})
	e.Logger.Fatal(e.Start(":8001"))
}
