package main

import (
	"context"
	"fmt"
	"rest-elliot/internal/comment"
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

	commentService := comment.NewService(db)

	commentService.PostComment(
		context.Background(),
		comment.Comment{
			ID:     "a0d9ff7a-fd8a-4baf-b072-dc495ae534e1",
			Slug:   "manuel-test",
			Author: "Abbas",
			Body:   "body",
		},
	)

	fmt.Println(commentService.GetComment(
		context.Background(),
		"a0d9ff7a-fd8a-4baf-b072-dc495ae534e1",
	))

	fmt.Println("Succesfully connected and pinged db")
	return nil
}

func main() {
	fmt.Println("Go REST course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
