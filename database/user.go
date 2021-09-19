package database

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(64);uniqueIndex"`
	Password string `json:"password"`
	Token    string `json:"token" gorm:"index"`
}

func InsertUser(u *User) error {
	return db.Create(u).Error
}

func UpdateUser(u *User) error {
	return db.Model(u).Updates(u).Error
}

func UpdateUserToken(userID uint, token string) error {
	u := &User{
		Model: gorm.Model{ID: userID},
	}
	return db.Model(u).UpdateColumns(map[string]interface{}{
		"token": token,
	}).Error
}

func FindUserByID(id uint) (*User, error) {
	u := &User{}
	err := db.First(u, "id = ?", id).Error
	return u, err
}

func FindUserByUserName(name string) (*User, error) {
	u := &User{}
	err := db.First(u, "username = ?", name).Error
	return u, err
}

func FindUserByToken(token string) (*User, error) {
	if token == "" {
		return nil, errors.New("empty user token")
	}
	u := &User{}
	err := db.First(u, "token = ?", token).Error
	return u, err
}
