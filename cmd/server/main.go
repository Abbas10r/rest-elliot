package main

import (
	"fmt"
	"rest-elliot/internal/db"
)

// This func will be responsible for instaniation anmd startup of our app
func Run() error {
	fmt.Println("Starting out app")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to db")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	fmt.Println("Succesfully connected and pinged db")
	return nil
}

func main() {
	fmt.Println("Go REST course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
