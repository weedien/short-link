package adapter

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"math/rand"
	"shortlink/internal/link/constant"
	"shortlink/internal/user/adapter/po"
	"shortlink/internal/user/domain/group"
)

type GroupRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewGroupRepositoryImpl(db *gorm.DB, rdb *redis.Client) GroupRepository {
	return GroupRepository{db: db, rdb: rdb}
}

func (r GroupRepository) ListGroup(ctx context.Context, username string) ([]group.Group, error) {
	var groupPos []po.Group
	if err := r.db.WithContext(ctx).
		Where("username = ?", username).
		Order("sort_order").
		Find(&groupPos).Error; err != nil {
		return nil, err
	}
	groups := make([]group.Group, 0, len(groupPos))
	for _, groupPo := range groupPos {
		groups = append(groups, group.NewGroup(groupPo.Gid, groupPo.Username, groupPo.Name, groupPo.SortOrder))
	}
	return groups, nil
}

func (r GroupRepository) CreateGroup(ctx context.Context, g group.Group) error {
	groupPo := po.Group{
		Gid:       g.Gid(),
		Username:  g.Username(),
		Name:      g.Name(),
		SortOrder: g.SortOrder(),
	}
	if err := r.db.WithContext(ctx).Create(&groupPo).Error; err != nil {
		return err
	}
	return nil
}

func (r GroupRepository) GetGroupSize(ctx context.Context, username string) (int, error) {
	var count int64
	if err := r.db.Model(&po.Group{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r GroupRepository) UniqueReturnGid(ctx context.Context) (string, error) {
	gid := RandomString(6)
	exist, err := r.rdb.BFExists(ctx, constant.GidRegisterCacheBloomFilter, gid).Result()
	if err != nil {
		return "", err
	}
	if !exist {
		groupUniquePo := po.GroupUnique{Gid: gid}
		if err = r.db.Create(&groupUniquePo).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return "", err
			}
		}
		return gid, nil
	}
	return "", nil
}

func (r GroupRepository) UpdateGroupName(ctx context.Context, g group.Group) error {
	username := ctx.Value("username").(string)
	groupPo := po.Group{}
	if err := r.db.WithContext(ctx).Model(&po.Group{}).
		Where("gid = ?", g.Gid()).Where("username = ?", username).
		First(&groupPo).Error; err != nil {
		return err
	}
	groupPo.Name = g.Name()
	if err := r.db.WithContext(ctx).Save(&groupPo).Error; err != nil {
		return err
	}
	return nil
}

func (r GroupRepository) UpdateGroupSortOrder(ctx context.Context, g group.Group) error {
	username := ctx.Value("username").(string)
	groupPo := po.Group{}
	if err := r.db.WithContext(ctx).Model(&po.Group{}).
		Where("gid = ?", g.Gid()).Where("username = ?", username).
		First(&groupPo).Error; err != nil {
		return err
	}
	groupPo.SortOrder = g.SortOrder()
	if err := r.db.WithContext(ctx).Save(&groupPo).Error; err != nil {
		return err
	}
	return nil
}

func (r GroupRepository) DeleteGroup(ctx context.Context, gid string) error {
	username := ctx.Value("username").(string)
	if err := r.db.WithContext(ctx).
		Delete(&po.Group{}, "gid = ? and username = ?", gid, username).Error; err != nil {
		return err
	}
	return nil
}

func RandomString(n int) string {
	var letterBytes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
