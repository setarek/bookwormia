package main

import (
	"bookwormia/admin/cmd"
	"os"
)

const (
	APP_NAME = "admin"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
