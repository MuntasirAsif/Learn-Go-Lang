package main

import "fmt"

func main() {
// ...existing code...
    fmt.Println("stuct in go")

    user1 := User{"John", "john@example.com", false, 25}

    fmt.Printf("User 1: %+v\n", user1)
}
// ...existing code...

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
