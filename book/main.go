package main

import (
	"bookwormia/book/internal/cache"
	"bookwormia/book/internal/handler"
	"bookwormia/book/internal/repository"
	"bookwormia/book/internal/router"
	"bookwormia/book/internal/worker"
	"bookwormia/pkg/config"
	"bookwormia/pkg/logger"
	"bookwormia/pkg/postgres"
	"bookwormia/pkg/redis"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/spf13/viper"
)

const (
	APP_NAME = "book"
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

	redisClient := redis.GetRedisClient(conf)

	if err := postgres.Migrate(context.Background(), fmt.Sprintf("%s/%s", currentPath, viper.GetString("db_path")), db); err != nil {
		logger.Logger.Error().Err(err).Msg("error while migrating db")
		os.Exit(1)
	}

	bookRepository := repository.NewBookRepository(db)
	c := cache.NewCache(redisClient)

	var wg sync.WaitGroup

	createOneBook := worker.NewWorker(worker.CreateOneBookWorker, c, bookRepository)
	createBulkBook := worker.NewWorker(worker.CreateBulkBooksWorker, c, bookRepository)
	editOneBook := worker.NewWorker(worker.EditOneBookWorker, c, bookRepository)
	editBulkBook := worker.NewWorker(worker.EditBulkBooksWorker, c, bookRepository)
	deleteOneBook := worker.NewWorker(worker.DeleteOneBookWorker, c, bookRepository)
	deleteBulkBook := worker.NewWorker(worker.DeleteBulkBooksWorker, c, bookRepository)

	wg.Add(1)
	go createOneBook.Start(&wg)
	go createBulkBook.Start(&wg)
	go editOneBook.Start(&wg)
	go editBulkBook.Start(&wg)
	go deleteOneBook.Start(&wg)
	go deleteBulkBook.Start(&wg)

	router := router.New()
	v1 := router.Group("/api/v1")
	handler := handler.NewHandler(bookRepository, c)
	handler.Register(v1)

	router.Start(fmt.Sprintf("%s:%s", conf.GetString("hostname"), conf.GetString("port")))

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	createOneBook.Stop()
	createBulkBook.Stop()
	editOneBook.Stop()
	editBulkBook.Stop()
	deleteOneBook.Stop()
	deleteBulkBook.Stop()

	wg.Wait()
}
