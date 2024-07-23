package app

import "tradething/config"

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
