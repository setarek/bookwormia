package config

import (
	"fmt"

	"bookwormia/pkg/logger"

	"github.com/spf13/viper"
)

// todo: improve config to validate and check variables
type Config struct {
	*viper.Viper
}

func InitConfig(appName string, currentPath string, defaultConfigName string) (*Config, error) {

	config := &Config{viper.New()}
	config.AutomaticEnv()
	if err := viper.BindEnv("env"); err != nil {
		logger.Logger.Error().Err(err).Msg("error while binding env in config file")
		return nil, err
	}

	config.SetConfigName(defaultConfigName)

	config.AddConfigPath(fmt.Sprintf("%s/%s/config/", currentPath, appName))
	config.AddConfigPath(fmt.Sprintf("/go/src/%s/config/", appName))
	if err := config.ReadInConfig(); err != nil {
		logger.Logger.Error().Err(err).Msg("Failed to read config file")
	}

	config.WatchConfig()
	return config, nil
}
