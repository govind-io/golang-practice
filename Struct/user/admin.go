package user

type Admin struct {
	User
	email    string
	password string
}

func NewAdmin(email, password, fname, lname, dob string) *Admin {
	var admin = &Admin{
		User:     User{firstName: fname, lastName: lname, birthDate: dob},
		email:    email,
		password: password,
	}

	return admin
}
