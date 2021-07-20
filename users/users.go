package users

import (
	"sync"
	"golang.org/x/crypto/bcrypt"
)
//Stire us vert sunoke ub nenirt databasem that we'll use to store our users.
//It is protected by a read-write mutex, so that tow goroutines can;t moify it
//The underlying map at the same time(since maps are not safe for concurrent use in GO)
type Store struct{
	rwm *sync.RWMutex
	m map[string]string//map protected by a read-write mutex
	//They are not safe for concurrent use so we protect them
}

//newDB is a convenience method to initalize our in memory DB when our program starts.
func newDB() *Store {
	return &Store{
		rwm: &sync.RWMutex{},
		m: make(map[string]string),
	}
}
//New USer accepts a username and password and creates a new user in our DB

func NewUser(username string, password string) error{
	err :=exists(username)
	if err!=nil{
		return err
	}
	DB.rwm.Lock()//We will use sync package lock to protect our map it locks our maps for writing
	defer DB.rwm.Unlock()

	hashedPassword, err:=bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	//To hash and salt our password we'll use bcrypt.GenerateFromPassword
	//This will create both salt and hash for us.
	//It accepts byte slices not string so we'll convert our password to a byte slice
	//The second allice is for the amount of rounds we want to use with bcrypt function
	if err!=nil{
		return err
	}
	DB.m[username]=string(hashedPassword)//After securing our password we'll store it in our map
	return nil
}
//AuthenticateUser accepts a username and password, and checks that the given 
//password matches the hashed password. It returns nil on sucess, and an error on failure.
func AuthenticateUser(username string, password string) error{
	DB.rwm.RLock()//We will use sync package lock to protect our map it locks our maps for reading
	defer DB.rwm.RUnlock()

	hashedPassword:= DB.m[username]//return hashedpasword from our map as username as the key
	err:=bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err
}


//exists is an internal utility function for ensuring the usernames are unique
func exists(username string) error{
	DB.rwm.RLock()//locks our map while we read from it.
	defer DB.rwn.RUnlock()//make sure to unlock it after our function returns

	if DB.m[username]!=""{
		return ErrUserAlreadyExists
	}
	return nil
}