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


// Comment struct represents a comment entity with user and post details.
type Comment struct {
	ID      int
	Content string
	User    User
	Post    Post
}



// UserCreate function creates a new user entity.
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
