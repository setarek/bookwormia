package main

import (
	"bookwormia/pkg/config"
	"bookwormia/pkg/logger"
	"bookwormia/pkg/postgres"
	"bookwormia/user/internal/handler"
	"bookwormia/user/internal/repository"
	"bookwormia/user/internal/router"
	"context"
	"fmt"
	"os"
)

const (
	APP_NAME = "user"
)

func main() {

	currentPath, err := os.Getwd()
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while getting current path")
		os.Exit(1)
	}

	conf, err := config.InitConfig(APP_NAME, currentPath, "config")
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while initializing config")
		os.Exit(1)
	}

	db, err := postgres.InitDB(conf)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while initializing DB")
		os.Exit(1)
	}

	if err := postgres.Migrate(context.Background(), conf.GetString("db_path"), db); err != nil {
		logger.Logger.Error().Err(err).Msg("error while migrating db")
		os.Exit(1)
	}

	userRepository := repository.NewUserRepository(db)

	router := router.New()
	v1 := router.Group("/api/v1")
	handler := handler.NewHandler(conf, userRepository)
	handler.Register(v1)

	router.Start(fmt.Sprintf("%s:%s", conf.GetString("hostname"), conf.GetString("port")))

}
