package user

import (
	"shortlink/internal/user/app/user/command"
	"shortlink/internal/user/app/user/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	UserRegister command.UserRegisterHandler
	UserLogin    command.UserLoginHandler
	UserLogout   command.UserLogoutHandler
	UpdateUser   command.UpdateUserHandler
	DeleteUser   command.DeleteUserHandler
}

type Queries struct {
	GetUser        query.GetUserHandler
	CheckLogin     query.CheckLoginHandler
	CheckUserExist query.CheckUserExistHandler
}
