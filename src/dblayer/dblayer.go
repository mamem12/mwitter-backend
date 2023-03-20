package dblayer

import (
	"fmt"
	"mwitter-backend/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserLayer interface {
	CreateUser(user *models.User) error
	SignInUser(email, password string) (*models.User, error)
	SelectUserByEmail(email string) (*models.User, error)
	UpdateProfile(id string, UpdateInfo *models.User) error
}

type BookLayer interface {
	GetAllBook(per int, pageIdx int, sort string) (*[]models.BookInfo, error)
	GetBookInfoById(id uint) (models.BookInfo, error)
	GetBookInfoWithRank(id uint) ([]models.BookRank, error)
	GetBookInfoWithPoint(id uint) ([]models.BookPoint, error)
	GetBookInfoWithPrice(id uint) ([]models.BookPrice, error)
	GetBookInfoWithSummary(id uint) (models.BookSummary, error)
}

type DBORM struct {
	// *gorm.DB
	*UserORM
	*BookORM
}

var DB *DBORM

func NewORM(dbname string, con gorm.Config) (*DBORM, error) {

	if DB == nil {
		dsn := fmt.Sprintf("root@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=true", dbname)
		dsn = dsn + "&loc=Asia%2FSeoul"
		db, err := gorm.Open(mysql.Open(dsn), &con)

		if err != nil {
			return nil, err
		}

		db.AutoMigrate(
			&models.User{},
			&models.BookInfo{},
			&models.BookPoint{},
			&models.BookPrice{},
			&models.BookRank{},
			&models.BookSummary{},
		)
		DB = &DBORM{
			// DB: db,
			&UserORM{DB: db},
			&BookORM{DB: db},
		}
	}

	return DB, nil
}
