package main

import (
	"context"
	"fmt"
	"os"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

var DBPool *pgxpool.Pool

type Item struct {
	ID        string `json:"id"`
	Product   string `json:"product"`
	Serial    string `json:"serial"`
	Condition string `json:"condition"`
	Year      string `json:"year"`
}

func getItems(db *pgxpool.Pool) ([]Item, error) {
	var items []Item
	err := pgxscan.Select(context.Background(), db, &items, "SELECT * FROM items")
	if err != nil {
		return nil, err
	}
	return items, nil
}

func getItem(db *pgxpool.Pool, id string) (*Item, error) {
	var item Item
	err := pgxscan.Get(context.Background(), db, &item, "SELECT * FROM items WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func createItem(db *pgxpool.Pool, i Item) (*Item, error) {
	var item Item
	err := pgxscan.Get(context.Background(), db, &item, "INSERT INTO items (product, serial, condition, year) VALUES ($1, $2, $3, $4) returning *", &i.Product, &i.Serial, &i.Condition, &i.Year)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// func updateItem(db *pgxpool.Pool, id string) {}

func deleteItem(db *pgxpool.Pool, id string) bool {
	res, err := db.Exec(context.Background(), "DELETE FROM items WHERE id = $1", id)
	if err != nil {
		return false
	}
	count := res.RowsAffected()
	return count != 0
}

func main() {
	DBPool, err := pgxpool.Connect(context.Background(), "postgres://postgres:root1234@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer DBPool.Close()

	_, err = DBPool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS ITEMS (
		id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
		product VARCHAR(150) NOT NULL,
		serial VARCHAR(150) NOT NULL,
		condition VARCHAR(150) NOT NULL,
		year VARCHAR(150) NOT NULL
	)`)
	if err != nil {
		fmt.Println(err.Error())
	}

	itemTwo, err := createItem(DBPool, Item{Product: "Macbook Pro", Serial: "cc324132", Condition: "Good", Year: "2021"})
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(*itemTwo)

	itemThree, err := getItem(DBPool, itemTwo.ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(*itemThree)

	itemsList, err := getItems(DBPool)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(len(itemsList))

	deleteStatus := deleteItem(DBPool, itemTwo.ID)
	fmt.Println(deleteStatus)
}
