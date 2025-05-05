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



// Create User function creates a new user entity.
func CreateUser(ID int, fullName string, email string, phone string, address string, isActive bool) User {
	user := User{
		FullName: fullName,
		ID:   ID,
		Email:  email,
		Phone:  phone,
		Address: address,
		IsActive: isActive,
	}

	return user
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
