package cmd

import (
	"bookwormia/pkg/config"
	"bookwormia/pkg/logger"
	"bookwormia/pkg/redis"
	"os"

	"github.com/spf13/cobra"

	redisPkg "github.com/go-redis/redis"
)

var (
	rds *redisPkg.Client
)

var RootCmd = &cobra.Command{
	Use:   "book",
	Short: "A CLI to create, update or delete book's information",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

func init() {
	// currentPath, _ := os.Getwd()
	currentPath := "/home/setareh/projects/go/src/github.com/setarek/bookwormia"
	config, err := config.InitConfig("admin", currentPath, "config")
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while initializing config")
	}

	rds = redis.GetRedisClient(config)
	addNewBookCmd.Flags().StringVarP(&newBookName, "name", "n", "", "Name of the book")
	addNewBookCmd.Flags().StringVarP(&newBookDetails, "details", "d", "", "Details of the book")
	addNewBookCmd.MarkFlagRequired("name")
	addNewBookCmd.MarkFlagRequired("details")
	RootCmd.AddCommand(addNewBookCmd)

	editBookCmd.Flags().Int64Var(&bookId, "id", 0, "Id of the book")
	editBookCmd.Flags().StringVarP(&newBookName, "name", "n", "", "Name of the book")
	editBookCmd.Flags().StringVarP(&bookDetails, "details", "d", "", "Details of the book")
	editBookCmd.MarkFlagRequired("id")
	RootCmd.AddCommand(editBookCmd)

	deleteBookCmd.Flags().Int64Var(&bookId, "id", 0, "Id of the book")
	deleteBookCmd.MarkFlagRequired("id")
	RootCmd.AddCommand(deleteBookCmd)

	newBooksCmd.Flags().StringVarP(&fileName, "name", "n", "", "Name of the book")
	newBooksCmd.Flags().StringVarP(&filePath, "path", "p", "", "Details of the book")
	newBooksCmd.MarkFlagRequired("name")
	newBooksCmd.MarkFlagRequired("path")
	RootCmd.AddCommand(newBooksCmd)

	editBooksCmd.Flags().StringVarP(&fileName, "name", "n", "", "Name of the book")
	editBooksCmd.Flags().StringVarP(&filePath, "path", "p", "", "Details of the book")
	editBooksCmd.MarkFlagRequired("name")
	editBooksCmd.MarkFlagRequired("path")
	RootCmd.AddCommand(editBooksCmd)

	deleteBooksCmd.Flags().StringVarP(&bookIds, "ids", "i", "", "Ids of books")
	// #todo: get ids as slice of int64
	deleteBooksCmd.MarkFlagRequired("ids")
	RootCmd.AddCommand(deleteBooksCmd)
}
