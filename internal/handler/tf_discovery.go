package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var RemoteServiceDiscovery = Route{
	Methods: []string{"GET"},
	Path:    "/.well-known/terraform.json",
	Func: func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"modules.v1":   ModulesV1Path + "/",
			"providers.v1": ProvidersV1Path + "/",
		})
	},
}

func init() {
	Routes = append(Routes, RemoteServiceDiscovery)
}
