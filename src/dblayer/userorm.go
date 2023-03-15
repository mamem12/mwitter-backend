package dblayer

import (
	"mwitter-backend/src/models"

	"gorm.io/gorm"
)

// type DBORM struct {
// 	*gorm.DB
// }

type UserORM struct {
	*gorm.DB
}

// func NewORM(dbname string, con gorm.Config) (*DBORM, error) {

// 	dsn := fmt.Sprintf("root@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=true", dbname)
// 	dsn = dsn + "&loc=Asia%2FSeoul"
// 	db, err := gorm.Open(mysql.Open(dsn), &con)

// 	db.AutoMigrate(
// 		&models.User{},
// 		&models.BookInfo{},
// 		&models.BookPoint{},
// 		&models.BookPrice{},
// 		&models.BookRank{},
// 		&models.BookSummary{},
// 	)

// 	return &DBORM{
// 		DB: db,
// 	}, err
// }

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
