package models

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Item struct {
	ID        string `json:"id"`
	Product   string `json:"product"`
	Serial    string `json:"serial"`
	Condition string `json:"condition"`
	Year      string `json:"year"`
}

func GetItems(db *pgxpool.Pool) ([]Item, error) {
	var items []Item
	err := pgxscan.Select(context.Background(), db, &items, "SELECT * FROM items")
	if err != nil {
		return nil, err
	}
	return items, nil
}

func GetItem(db *pgxpool.Pool, id string) (*Item, error) {
	var item Item
	err := pgxscan.Get(context.Background(), db, &item, "SELECT * FROM items WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func CreateItem(db *pgxpool.Pool, i Item) (*Item, error) {
	var item Item
	err := pgxscan.Get(context.Background(), db, &item, "INSERT INTO items (product, serial, condition, year) VALUES ($1, $2, $3, $4) RETURNING *", &i.Product, &i.Serial, &i.Condition, &i.Year)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func UpdateItem(db *pgxpool.Pool, id string, i Item) (*Item, error) {
	var item Item
	err := pgxscan.Get(context.Background(), db, &item, "UPDATE items SET product = $2, serial = $3, condition = $4, year = $5 WHERE id = $1 RETURNING *", id, &i.Product, &i.Serial, &i.Condition, &i.Year)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func DeleteItem(db *pgxpool.Pool, id string) bool {
	res, err := db.Exec(context.Background(), "DELETE FROM items WHERE id = $1", id)
	if err != nil {
		return false
	}
	count := res.RowsAffected()
	return count != 0
}
