package set1

import (
	"errors"
	"fmt"
)

type Authenticator interface {
	Authenticate(username, password string) error
}

var (
	ErrUserName = errors.New("Invalid Username.")
	ErrPassword = errors.New("Wrong Password.")
)

type Username struct {
	Name string
}

func (u Username) Authenticate(username, password string) error {
	if u.Name == username {
		return nil
	}
	return ErrUserName
}

type Password struct {
	Key string
}

func (p Password) Authenticate(Username, password string) error {
	if p.Key == password {
		return nil
	}
	return ErrPassword
}

func AuthSystem() {
	user := Username{"Abhi"}
	pass := Password{"1234"}

	var username string
	fmt.Println("Enter Username:")
	fmt.Scan(&username)
	var password string

	if Uerr := user.Authenticate(username, password); Uerr != nil {
		fmt.Println(Uerr)
		return
	} else {
		fmt.Println("Username Verified.")
	}

	fmt.Println("Enter password:")
	fmt.Scan(&password)
	if Perr := pass.Authenticate(username, password); Perr != nil {
		fmt.Println(Perr)
		return
	} else {
		fmt.Println("Password verified.")
	}

}
