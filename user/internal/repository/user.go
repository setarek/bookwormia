package repository

import (
	"bookwormia/user/internal/model"
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) IUser {
	return UserRepository{
		DB: db,
	}
}

func (u UserRepository) CreateOrGetUser(ctx context.Context, email string, password string) (int64, error) {

	query := `
		select id, password from users where email = $1;
	`

	var user model.User
	err := u.DB.QueryRowContext(ctx, query, email).Scan(&user.Id, &user.Password)
	if err == sql.ErrNoRows {
		// todo: check password is string
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)
		if err != nil {
			return 0, err
		}
		insertQuery := `
			insert into users (email, password) values ($1, $2) returning id;
		`

		if err := u.DB.QueryRowContext(ctx, insertQuery, email, hashedPassword).Scan(&user.Id); err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return 0, err
	}

	return user.Id, nil
}
