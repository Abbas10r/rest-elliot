package main

import "fmt"

// This func will be responsible for instaniation anmd startup of our app
func Run() error {
	fmt.Println("Starting out app")
	return nil
}

func main() {
	fmt.Println("Go REST course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
