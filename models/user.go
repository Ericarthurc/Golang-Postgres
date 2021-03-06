package models

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Role     string   `json:"role"`
	Tokens   []string `json:"tokens"`
}

func GetUsers() {}

func GetUserByID(db *pgxpool.Pool, id string) (*User, error) {
	var user User
	err := pgxscan.Get(context.Background(), db, &user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(db *pgxpool.Pool, username string) (*User, error) {
	var user User
	err := pgxscan.Get(context.Background(), db, &user, "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CheckForToken(db *pgxpool.Pool, id string, token string) bool {
	var tokens []string
	err := pgxscan.Get(context.Background(), db, &tokens, "SELECT tokens FROM users WHERE id = $1", &id)
	if err != nil {
		return false
	}

	for _, v := range tokens {
		if v == token {
			return true
		}
	}
	return false
}

func LoginUser() {}

func LogoutUser(db *pgxpool.Pool, id string, token string) error {
	_, err := db.Exec(context.Background(), "UPDATE users SET tokens = array_remove(tokens, $2) WHERE id = $1", &id, &token)
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(db *pgxpool.Pool, u User) (*User, error) {
	var user User
	err := pgxscan.Get(context.Background(), db, &user, "INSERT INTO users (id, username, password, role, tokens) VALUES ($1, $2, $3, $4, $5) RETURNING *", &u.ID, &u.Username, &u.Password, &u.Role, &u.Tokens)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUserTokens(db *pgxpool.Pool, u User) (*User, error) {
	var user User
	err := pgxscan.Get(context.Background(), db, &user, "UPDATE users SET tokens = $2 WHERE id = $1 RETURNING *", &u.ID, &u.Tokens)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser() {}

func DeleteUser() {}
