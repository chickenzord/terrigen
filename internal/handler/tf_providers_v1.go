package handler

import (
	"net/http"

	"github.com/chickenzord/terrigen/internal/terraform"
	"github.com/labstack/echo/v4"
)

var ProvidersV1Path = "/v1/providers"

func ProviderList(d terraform.ProviderDiscovery) Route {
	return Route{
		Methods: []string{"GET"},
		Path:    ProvidersV1Path,
		Func: func(c echo.Context) error {

			providers, err := d.Discover()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, H{
					"status": "error",
					"error":  err.Error(),
				})
			}

			return c.JSON(http.StatusOK, H{
				"status":    "success",
				"providers": providers,
			})
		},
	}
}

func ProviderVersions(d terraform.ProviderDiscovery) Route {
	return Route{
		Methods: []string{"GET"},
		Path:    ProvidersV1Path + "/:namespace/:type/versions",
		Func: func(c echo.Context) error {
			providerNamespace := c.Param("namespace")
			providerType := c.Param("type")

			providers, err := d.Discover()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, H{
					"status": "error",
					"error":  err.Error(),
				})
			}

			versions, err := terraform.FindVersions(providers, providerNamespace, providerType)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, H{
					"status": "error",
					"error":  err.Error(),
				})
			}

			return c.JSON(http.StatusOK, map[string]interface{}{
				"versions": versions,
			})
		},
	}
}

func init() {
	d := terraform.NewLocalProviderDiscovery(terraform.LocalProviderDiscoveryOpts{
		PluginsPath:    terraform.DefaultPluginsPath(),
		RegistryFilter: terraform.StringFilterMatchAll,
	})

	Routes = append(Routes,
		ProviderList(d),
		ProviderVersions(d),
	)
}
