package user

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (appUser *User) ClearUserName() {
	appUser.firstName = "------"
	appUser.lastName = "------"
}

func (appUser *User) OutPutUserData() {
	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthDate, appUser.createdAt)
}

func New(fName, lName, dob string) *User {
	return &User{
		firstName: fName,
		lastName:  lName,
		birthDate: dob,
		createdAt: time.Now(),
	}
}
