package main

type User struct {
	username string
	email    string
	age      int
}
type UserManager struct {
	Users []User
}

func (u *UserManager) AddUser(username, email string, age int) {
	u.Users = append(u.Users, User{username, email, age})
}
func (u UserManager) UsersOlderThan(age int) []User {
	x := make([]User, 0)
	for i, v := range u.Users {
		if v.age > age {
			x = append(x, u.Users[i])
		}
	}
	return x
}
