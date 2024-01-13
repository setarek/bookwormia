package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"bookwormia/admin/internal/data_handler"
	"bookwormia/admin/internal/redis_repository"
	"bookwormia/pkg/logger"
)

var (
	newBookName    string
	newBookDetails string

	bookId      int64
	bookName    string
	bookDetails string
)

var addNewBookCmd = &cobra.Command{
	Use:   "addNewBook",
	Short: "queue new book information",
	Run: func(cmd *cobra.Command, args []string) {
		if newBookName == "" || newBookDetails == "" {
			logger.Logger.Warn().Msg("Please provide both name and details flag")
			cmd.Help()
			os.Exit(1)
		}
		book := data_handler.Book{
			Name:        newBookName,
			Description: newBookDetails,
		}

		fmt.Println("================== ", book)

		err := data_handler.NewDataHandler(*redis_repository.NewRedisRepository(rds)).CreateOne(book)
		if err != nil {
			logger.Logger.Error().Err(err).Msg("error while getting new book")
		}

		logger.Logger.Info().Str("book_name", newBookName).Msg("New book added!")
	},
}

var editBookCmd = &cobra.Command{
	Use:   "editBook",
	Short: "queue edited book information",
	Run: func(cmd *cobra.Command, args []string) {
		if bookId == 0 && (newBookName == "" || bookDetails == "") {
			logger.Logger.Warn().Msg("Please provide both name and details flag")
			cmd.Help()
			os.Exit(1)
		}
		book := data_handler.Book{
			Id:          bookId,
			Name:        newBookName,
			Description: bookDetails,
		}
		err := data_handler.NewDataHandler(*redis_repository.NewRedisRepository(rds)).UpdateOne(book)
		if err != nil {
			logger.Logger.Error().Err(err).Msg("error while getting new book")
		}

		logger.Logger.Info().Str("bookName", bookName).Msg("a book edited!")
	},
}

var deleteBookCmd = &cobra.Command{
	Use:   "deleteBook",
	Short: "queue delete book information",
	Run: func(cmd *cobra.Command, args []string) {
		if bookId == 0 {
			logger.Logger.Warn().Msg("Please provide bookId flag")
			cmd.Help()
			os.Exit(1)
		}

		err := data_handler.NewDataHandler(*redis_repository.NewRedisRepository(rds)).DeleteOne(bookId)
		if err != nil {
			logger.Logger.Error().Err(err).Msg("error while deleting a book")
		}

		logger.Logger.Info().Int64("bookId", bookId).Msg("a book deleted!")
	},
}
