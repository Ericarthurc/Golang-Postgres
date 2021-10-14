package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

var DBPool *pgxpool.Pool

type Item struct {
	ID        string
	Product   string
	Serial    string
	Condition string
	Year      string
}

func main() {
	DBPool, err := pgxpool.Connect(context.Background(), "postgres://postgres:root1234@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer DBPool.Close()

	var item Item

	start := time.Now()
	for i := 0; i < 1489; i++ {

		err = pgxscan.Get(context.Background(), DBPool, &item, "SELECT * FROM items WHERE id = $1", "696d9aa6-8b85-47f6-b92b-af8d34244c6d")
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			os.Exit(1)
		}

	}
	duration := time.Since(start)
	fmt.Println(duration.Seconds())

	// fmt.Println(item.ID)
}
