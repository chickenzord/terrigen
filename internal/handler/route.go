package handler

import "github.com/labstack/echo/v4"

type Route struct {
	Methods     []string
	Path        string
	Func        echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc
}

var Routes = []Route{}
