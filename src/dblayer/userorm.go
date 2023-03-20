package dblayer

import (
	"mwitter-backend/src/models"

	"gorm.io/gorm"
)

type UserORM struct {
	*gorm.DB
}

func (db *UserORM) CreateUser(user *models.User) error {

	return db.Create(&user).Error
}

func (db *UserORM) SignInUser(email, password string) (userInfo *models.User, err error) {

	result := db.Table("users").Select("nickname", "email", "id").Where(&models.User{Email: email, Password: password})

	return userInfo, result.Find(&userInfo).Error
}

func (db *UserORM) UpdateProfile(id string, updateInfo *models.User) error {

	return db.Table("users").Where("id = ?", id).Update("nickname", updateInfo.Nickname).Error
}

func (db *UserORM) SelectUserByEmail(email string) (existUser *models.User, err error) {

	result := db.Table("users").Select("email").Where(&models.User{Email: email})

	return existUser, result.Find(&existUser).Error
}
