package model

import (
	"fmt"
)

type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

var users []User

func GetUsers() []User {
	return users
}

func AddUser(user User) User {

	users = append(users, user)
	return user
}

func SetupModel() {
	println("Dummy - do nothing")
	// err := storage.DB_.AutoMigrate(&User{})
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "unable to connect to execute the query: %v\n", err)
	// }
}

func SetStatus(user User, status string) User {
	user.Status = status
	return user
}

// works with fmt.Println(user)
func (user User) String() string {
	//return println("User: ",user.ID, " Name: ", user.Name, " Email: ",user.Email)
	return fmt.Sprintf("User: %s Name: %s Email: %s Status: %s", user.ID, user.Name, user.Email, user.Status)
}
