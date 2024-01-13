package handler

import (
	"bookwormia/pkg/config"
	"bookwormia/user/internal/repository"
)

type UserHandler struct {
	Config     *config.Config
	Repository repository.IUser
}

func NewHandler(config *config.Config, userRepository repository.IUser) *UserHandler {
	return &UserHandler{
		Config:     config,
		Repository: userRepository,
	}
}
