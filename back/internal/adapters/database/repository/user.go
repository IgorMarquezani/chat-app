package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"app/internal/core/models/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) (*UserRepository, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}

	return &UserRepository{
		db: db,
	}, nil
}

func (r *UserRepository) Insert(ctx context.Context, u *user.User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *UserRepository) SelectByEmail(ctx context.Context, email string) (user.User, error) {
	var u user.User

	return u, r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
}

func (r *UserRepository) SelectByID(ctx context.Context, userID uint32) (user.User, error) {
	var u user.User

	return u, r.db.WithContext(ctx).Model(&u).Where("id = ?", userID).First(&u).Error
}

func (r *UserRepository) SelectByName(ctx context.Context, userName string) ([]user.User, error) {
	var (
		regex string
		users = make([]user.User, 0)
	)

	strs := strings.SplitSeq(userName, " ")

	for str := range strs {
		regex += fmt.Sprintf("(%s).*", str)
	}

	return users, r.db.WithContext(ctx).Model(&user.User{}).Where("name ~* ?", regex).Scan(&users).Error
}

func (r *UserRepository) DeleteByID(ctx context.Context, id uint32) error {
	return r.db.WithContext(ctx).Delete(&user.User{ID: id}).Error
}
