package app

import (
	"net/http"
	"tradething/config"

	"github.com/labstack/echo/v4"
)

func ReadLog(path_config string) (*config.AppConfig, error) {
	config, err := config.ReadConfig(path_config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func ReadAWSAppLog() (*config.AppConfig, error) {
	config, err := config.ReadAWSAppConfig()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func UpdateConfig(g *echo.Echo, config *config.AppConfig) {
	g.GET("/config/update", func(c echo.Context) error {
		_config, err := ReadAWSAppLog()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()})
		}
		config = _config
		return c.JSON(http.StatusOK, &echo.HTTPError{Code: http.StatusOK, Message: "config updated"})
	})
}
