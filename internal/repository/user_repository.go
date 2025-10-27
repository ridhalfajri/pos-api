package repository

import (
	"github.com/ridhalfajri/pos-api/internal/repository/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user model.User) (model.User, error)
	FindByUsername(username string) (model.User, error)
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (repository *UserRepositoryImpl) FindByUsername(username string) (model.User, error) {
	user := model.User{}
	err := repository.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repository *UserRepositoryImpl) Register(user model.User) (model.User, error) {
	err := repository.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
