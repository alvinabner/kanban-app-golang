package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	user := entity.User{}
	err := r.db.WithContext(ctx).Where("id = ?", id).Limit(1).Find(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	user := entity.User{}
	err := r.db.WithContext(ctx).Where("email = ?", email).Limit(1).Find(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	err := r.db.WithContext(ctx).Updates(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	user := entity.User{}
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}
