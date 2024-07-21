package app

import "tradething/config"

func ReadLog() (*config.AppConfig, error) {
	config, err := config.ReadConfig()
	if err != nil {
		return nil, err
	}
	return config, nil
}
