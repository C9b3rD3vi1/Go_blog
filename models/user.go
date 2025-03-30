package models

import "fmt"

// User struct
type User struct {
	Name    string
	Id      int
	Email   string
	Phone   string
	Address string
	Status  bool
}

func UserCreate() {
	fmt.Println("User created")
}
