package cmd

import (
	"bookwormia/admin/internal/data_handler"
	"bookwormia/pkg/logger"
	"bookwormia/pkg/utils"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"bookwormia/admin/internal/redis_repository"

	"github.com/spf13/cobra"
)

var (
	filePath string
	fileName string

	bookIds string
)

var newBooksCmd = &cobra.Command{
	Use:   "newBooks",
	Short: "queue new book information",
	Run: func(cmd *cobra.Command, args []string) {

		if filePath == "" || fileName == "" {
			logger.Logger.Warn().Msg("Please provide both file name and file path flag")
			cmd.Help()
			os.Exit(1)
		}

		file, err := os.Open(fmt.Sprintf("%s/%s", filePath, fileName))
		if err != nil {
			logger.Logger.Error().Err(err).Msg("error while opening csv file")
			os.Exit(1)
		}
		defer file.Close()

		reader := csv.NewReader(file)

		header, err := reader.Read()
		if err != nil {
			logger.Logger.Warn().Msg("error while reading csv header")
			os.Exit(1)
		}

		records, err := reader.ReadAll()
		if err != nil {
			logger.Logger.Error().Err(err).Msg("error while reading csv body")
			os.Exit(1)
		}

		if err := data_handler.NewDataHandler(*redis_repository.NewRedisRepository(rds)).CreateBulk(header, records); err != nil {
			logger.Logger.Error().Err(err).Msg("error while getting new book")
		}

		logger.Logger.Info().Str("fileName", fileName).Msg("New books added!")
	},
}

var editBooksCmd = &cobra.Command{
	Use:   "editBooks",
	Short: "queue new book information",
	Run: func(cmd *cobra.Command, args []string) {

		if filePath == "" || fileName == "" {
			logger.Logger.Warn().Msg("Please provide both file name and file path flag")
			cmd.Help()
			os.Exit(1)
		}

		file, err := os.Open(fmt.Sprintf("%s/%s", filePath, fileName))
		if err != nil {
			logger.Logger.Error().Err(err).Msg("error while opening csv file")
			os.Exit(1)
		}
		defer file.Close()

		reader := csv.NewReader(file)

		header, err := reader.Read()
		if err != nil {
			logger.Logger.Warn().Msg("error while reading csv header")
			os.Exit(1)
		}

		records, err := reader.ReadAll()
		if err != nil {
			logger.Logger.Error().Err(err).Msg("error while reading csv body")
			os.Exit(1)
		}

		if err := data_handler.NewDataHandler(*redis_repository.NewRedisRepository(rds)).UpdateBulk(header, records); err != nil {
			logger.Logger.Error().Err(err).Msg("error while getting new book")
		}

		logger.Logger.Info().Str("fileName", fileName).Msg("New books edited!")
	},
}

var deleteBooksCmd = &cobra.Command{
	Use:   "deleteBooks",
	Short: "queue book ids for deleting",
	Run: func(cmd *cobra.Command, args []string) {

		splitedIds := strings.Split(bookIds, ",")

		var ids []int64
		for _, strId := range splitedIds {
			id := utils.ParseInt64(strId)
			ids = append(ids, id)
		}

		if bookIds == "" {
			logger.Logger.Warn().Msg("Please add at list one book id")
			cmd.Help()
			os.Exit(1)
		}

		if err := data_handler.NewDataHandler(*redis_repository.NewRedisRepository(rds)).DeleteBulk(ids); err != nil {
			logger.Logger.Error().Err(err).Msg("error while deleting mutiple books")
		}

		logger.Logger.Info().Msg("Multiple books deleted!")
	},
}
