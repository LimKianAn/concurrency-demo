package main

import "fmt"

// User is the struct which persists in DB
type User struct {
	ID   int
	Name string
}

func (u *User) String() string {
	return fmt.Sprintf("Name: %q\n", u.Name)
}
