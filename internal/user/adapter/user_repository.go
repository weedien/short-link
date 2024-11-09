package adapter

import (
	"context"
	"errors"
	"github.com/bytedance/sonic"
	"gorm.io/gorm"
	"shortlink/internal/base/error_no"
	"shortlink/internal/link/constant"
	"shortlink/internal/user/adapter/po"
	"shortlink/internal/user/domain/user"
	"time"
)

type UserRepositoryImpl struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewUserRepositoryImpl(db *gorm.DB, rdb *redis.Client) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db, rdb: rdb}
}

func (r UserRepositoryImpl) GetUser(ctx context.Context, username string) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r UserRepositoryImpl) CheckUserExist(ctx context.Context, username string) (bool, error) {
	exist, err := r.rdb.BFExists(ctx, constant.UserRegisterBloomFilter, username).Result()
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (r UserRepositoryImpl) CreateUser(ctx context.Context, u *user.User) error {
	userPo := po.User{
		Username: u.Name(),
		Password: u.Password(),
		RealName: u.RealName(),
		Phone:    u.Phone(),
		Mail:     u.Email(),
	}
	if err := r.db.Create(&userPo).Error; err != nil {
		return err
	}
	if err := r.rdb.BFAdd(ctx, constant.UserRegisterBloomFilter, u.Name()).Err(); err != nil {
		return err
	}
	return nil
}

func (r UserRepositoryImpl) AddUserToBloomFilter(ctx context.Context, name string) error {
	if err := r.rdb.BFAdd(ctx, constant.UserRegisterBloomFilter, name).Err(); err != nil {
		return err
	}
	return nil
}

func (r UserRepositoryImpl) UpdateUser(ctx context.Context, u *user.User) error {
	//TODO implement me
	panic("implement me")
}

func (r UserRepositoryImpl) CheckLogin(ctx context.Context, username string, token string) (flag bool, err error) {
	var value string
	if value, err = r.rdb.HGet(ctx, constant.UserLoginKey+username, token).Result(); err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}
	}
	if value != "" {
		flag = true
	}
	return
}

func (r UserRepositoryImpl) InvalidateToken(ctx context.Context, username string, token string) error {
	err := r.rdb.HDel(ctx, constant.UserLoginKey+username, token).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r UserRepositoryImpl) Login(ctx context.Context, username string, password string) (token string, err error) {
	userPo := po.User{}
	if err = r.db.Where("username = ? AND password = ?", username, password).
		First(&userPo).Error; err != nil {
		return "", err
	}
	if userPo.ID == 0 {
		return "", error_no.UserNotExist
	}
	key := constant.UserLoginKey + username
	// 是否已经存在token
	tokens := make([]string, 0)
	if tokens, err = r.rdb.HKeys(ctx, key).Result(); err != nil {
		return "", err
	}
	if len(tokens) > 0 {
		// 重置过期时间
		if err = r.rdb.Expire(ctx, key, 30*24*time.Hour).Err(); err != nil {
			return "", err
		}
		return tokens[0], nil
	}
	token = user.GenToken()
	// hash key: login_username
	// 	field: token
	// 	value: user info
	var userInfo string
	if userInfo, err = sonic.MarshalString(&userPo); err != nil {
		return "", err
	}
	if err = r.rdb.HSet(ctx, key, token, userInfo).Err(); err != nil {
		return "", err
	}
	if err = r.rdb.Expire(ctx, key, 30*24*time.Hour).Err(); err != nil {
		return "", err
	}
	return token, nil
}

func (r UserRepositoryImpl) DeleteUser(id string) error {
	if err := r.db.Delete(&po.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
