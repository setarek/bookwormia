package repository

import "context"

type IUser interface {
	CreateOrGetUser(ctx context.Context, email string, password string) (int64, error)
}
