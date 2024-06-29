package models

import (
	"errors"
	"time"

	"first-task-alterra/internal/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Password  string
	Email     string
	Phone     string
	BirthDate time.Time `gorm:"type:date"`
	Todos     []Todo    `gorm:"foreignKey:Owner"`
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(connection *gorm.DB) *UserModel {
	return &UserModel{
		db: connection,
	}
}

func (um *UserModel) Login(email, password string) (User, error) {
	var user User
	err := um.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return User{}, errors.New("user not found")
	}
	err = utils.VerifyPassword(user.Password, password)
	if err != nil {
		return User{}, errors.New("invalid credentials")
	}
	return user, nil
}

func (um *UserModel) Register(newUser User, hashPassword string) (bool, error) {
	newUser.Password = hashPassword
	err := um.db.Create(&newUser).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
