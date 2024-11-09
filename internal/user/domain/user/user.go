package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	id         int
	name       string
	password   string
	realName   string
	email      string
	phone      string
	deleteTime time.Time
}

func NewUser(name, password, realName, email, phone string) User {
	return User{
		name:     name,
		password: password,
		realName: realName,
		email:    email,
		phone:    phone,
	}
}

func (u User) Name() string {
	return u.name
}

func (u User) Password() string {
	return u.password
}

func (u User) RealName() string {
	return u.realName
}

func (u User) Email() string {
	return u.email
}

func (u User) Phone() string {
	return u.phone
}

func (u User) DeleteTime() time.Time {
	return u.deleteTime
}

func (u User) Login() string {
	return uuid.NewString()
}

func GenToken() string {
	return uuid.NewString()
}
