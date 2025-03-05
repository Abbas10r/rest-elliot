package main

import (
	"context"
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
	if err := db.Ping(context.Background()); err != nil {
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
