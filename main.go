package main

import (
	"eric/database"
	"eric/models"
	"fmt"
	"os"
)

func main() {
	database.DbConnect()

	itemTwo, err := models.CreateItem(database.DBPool, models.Item{Product: "Macbook Pro", Serial: "cc324132", Condition: "Good", Year: "2021"})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Create failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(*itemTwo)

	itemThree, err := models.GetItem(database.DBPool, itemTwo.ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Get failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(*itemThree)

	itemsList, err := models.GetItems(database.DBPool)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gets failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(len(itemsList))

	updatedItem, err := models.UpdateItem(database.DBPool, itemThree.ID, models.Item{Product: "Asus Vivobook S15", Serial: "v2347v5278c3b4", Condition: "Bad", Year: "2017"})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Update failed!: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Updated:", *updatedItem)

	deleteStatus := models.DeleteItem(database.DBPool, itemTwo.ID)
	fmt.Println(deleteStatus)
}
