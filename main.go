package main

import (
	"fmt"

	"github.com/chickenzord/terrigen/internal/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	for _, r := range handler.Routes {
		e.Match(r.Methods, r.Path, r.Func, r.Middlewares...)
	}

	for _, r := range e.Routes() {
		fmt.Println(r.Method, r.Path)
	}

	e.Logger.Fatal(e.Start(":8000"))
}
