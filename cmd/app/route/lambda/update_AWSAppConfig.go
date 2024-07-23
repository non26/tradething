package route

import (
	"tradething/config"

	"github.com/labstack/echo/v4"
)

func UpdateAWSAppConfig(app *echo.Echo, _config *config.AppConfig) {
	app.GET("/update-aws-config", func(c echo.Context) error {
		var err error
		_config, err = config.ReadAWSAppConfig()
		if err != nil {
			return err
		}
		return nil
	})
}
