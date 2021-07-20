/*//How to hash and salt a password and verify it.
//go get golang.org/x/crypto/bcrypt
//or use scrypt package
//https://github.com/tonnerre/scrypt
//simple scrypt implementation in go
//as it's API is similar to bcrypt.

package main
import (
	"fmt"
	"os"
	"github.com/shikharvashistha/auth/users"
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
}*/
package main
import (
    "fmt"
    "log"    
	"golang.org/x/crypto/bcrypt"
)
func main() {    
	for { 
        // Enter a password and generate a salted hash
      //  pwd := getPwd()
     //   hash := hashAndSalt(pwd)
        
        // Enter the same password again and compare it with the
        // first password entered
      //  pwd2 := getPwd()
      //  pwdMatch := comparePasswords(hash, pwd2)
	//	fmt.Println("Passwords Match ?", pwd)    
	}
}


func getPwd() []byte {// Prompt the user to enter a password
    fmt.Println("Enter a password")// We will use this to store the users input
    var pwd string    // Read the users input
    _, err := fmt.Scan(&pwd)
    if err != nil {
        log.Println(err)
    }// Return the users input as a byte slice which will save us
    // from having to do this conversion later on
    return []byte(pwd)
}
func hashAndSalt(pwd []byte) string {
    
    // Use GenerateFromPassword to hash & salt pwd
    // MinCost is just an integer constant provided by the bcrypt
    // package along with DefaultCost & MaxCost. 
    // The cost can be any value you want provided it isn't lower
    // than the MinCost (4)
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }// GenerateFromPassword returns a byte slice so we need to
    // convert the bytes to a string and return it
    return string(hash)
}
func comparePasswords(hashedPwd string, plainPwd []byte) bool {    // Since we'll be getting the hashed password from the DB it
    // will be a string so we'll need to convert it to a byte slice
    byteHash := []byte(hashedPwd)    
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
    if err != nil {
        log.Println(err)
        return false
    }
    
    return true
}