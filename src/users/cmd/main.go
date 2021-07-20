//How to hash and salt a password and verify it.
//go get golang.org/x/crypto/bcrypt
//or use scrypt package
//https://github.com/tonnerre/scrypt
//simple scrypt implementation in go
//as it's API is similar to bcrypt.

package main
import (
	"fmt"
	"os"
	"github.com/shikharvashistha/auth/src/users"
)
func main() {
	username, password:="ben", "qwerty123"
	//create a new user with given password
	err:=users.NewUser(username, password)
	if err !=nil {
		fmt.Printf("Couldn't create user: %s\n", err.Error())
		return 
	}
	err.users.AuthenticateUser(username, password)
	//Then we'll try to authenticate the user 
	if err !=nil{
		fmt.Printf("Couldn't authenticate user: %s\n", err.Error())
		return 
	}
	fmt.Println("Successfully created and authenticated user %s", username)
	err=users.NewUser(username, password)
	if err !=nil{
		fmt.Printf("Couldn't create user: %s\n", err.Error())
		return
	}
	err.users.AuthenticateUser(username, password)
	//Then we'll try to authenticate the user 
	if err !=nil{
		fmt.Printf("Couldn't authenticate user: %s\n", err.Error())
		return 
	}
}