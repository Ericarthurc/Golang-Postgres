package models

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"-"`
	Tokens   []string `json:"-"`
}

func GetUsers() {}

func GetUser() {}

func GetUserByUsername(db *pgxpool.Pool, username string) (*User, error) {
	var user User
	err := pgxscan.Get(context.Background(), db, &user, "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func LoginUser() {}

func LogoutUser() {}

func CreateUser(db *pgxpool.Pool, u User) (*User, error) {
	var user User
	err := pgxscan.Get(context.Background(), db, &user, "INSERT INTO users (id, username, password, tokens) VALUES ($1, $2, $3, $4) RETURNING *", &u.ID, &u.Username, &u.Password, &u.Tokens)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser() {}

func DeleteUser() {}
