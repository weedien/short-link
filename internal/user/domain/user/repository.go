package user

import "context"

type Repository interface {
	GetUser(ctx context.Context, username string) (*User, error)
	CheckUserExist(ctx context.Context, username string) (bool, error)
	CreateUser(ctx context.Context, u *User) error
	AddUserToBloomFilter(ctx context.Context, name string) error
	UpdateUser(ctx context.Context, u *User) error
	CheckLogin(ctx context.Context, username string, token string) (bool, error)
	InvalidateToken(ctx context.Context, username string, token string) error
	Login(ctx context.Context, username string, password string) (string, error)
	DeleteUser(id string) error
}
