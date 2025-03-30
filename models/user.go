package models

import "fmt"

// User struct represents a user entity with personal and contact details.
type User struct {
	FullName    string
	ID      int
	Email   string
	Phone   string
	Address string
	IsActive  bool
}

func UserCreate() {
	user := User{
		FullName: "Nana Kwame",
		ID:   1,
		Email:  "nana.kwame@gmail.com",
        Phone:  "+2348123456789",
        Address: "123 Main St, Anytown, USA",
        IsActive: true,
    }
	
	fmt.Println(user)
	fmt.Println("User created")
}
