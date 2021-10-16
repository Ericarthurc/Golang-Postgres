package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DBPool *pgxpool.Pool

func DbConnect() {
	var err error
	DBPool, err = pgxpool.Connect(context.Background(), "postgres://postgres:root1234@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

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
}
