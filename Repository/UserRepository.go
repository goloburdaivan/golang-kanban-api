package Repository

import (
	"Golang/Models"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(username string, login string, password string) (*Models.User, error)
		Update(user *Models.User) (*Models.User, error)
		FindByEmail(email string) (*Models.User, error)
	}

	UserRepositoryImpl struct {
		db *gorm.DB
	}
)

func (u UserRepositoryImpl) FindByEmail(email string) (*Models.User, error) {
	user := &Models.User{}
	result := u.db.Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u UserRepositoryImpl) Create(
	username string,
	login string,
	password string,
) (*Models.User, error) {
	var user = &Models.User{
		Username: username,
		Email:    login,
		Role:     "user",
	}

	user.SetPassword(password)

	return user, u.db.Create(user).Error
}

func (u UserRepositoryImpl) Update(user *Models.User) (*Models.User, error) {
	return user, u.db.Save(user).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}
