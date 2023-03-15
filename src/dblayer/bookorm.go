package dblayer

import (
	"mwitter-backend/src/models"
	"time"

	"gorm.io/gorm"
)

type BookORM struct {
	*gorm.DB
}

func (db *BookORM) InsertBookInfo(bookInfo *models.BookInfo) error {
	return db.Create(&bookInfo).Error
}

func (db *BookORM) UpdateBookInfo(bookInfo *models.BookInfo) error {
	return db.Table("bookinfo").Where("hash = ?", bookInfo.Hash).Update("updated_at", time.Now().UTC()).Error
}

func (db *BookORM) InsertBookRank(bookRank *models.BookRank) error {
	return db.Create(&bookRank).Error
}

func (db *BookORM) InsertBookPrice(bookPrice *models.BookPrice) error {
	return db.Create(&bookPrice).Error
}

func (db *BookORM) InsertBookPoint(bookPoint *models.BookPoint) error {
	return db.Create(&bookPoint).Error
}

func (db *BookORM) InsertBookSummary(bookSummary *models.BookSummary) error {
	return db.Create(&bookSummary).Error
}

func (db *BookORM) GetBookInfoByHash(hash string) (existBook *models.BookInfo, err error) {
	result := db.Table("bookinfo").Select("hash").Where(&models.BookInfo{Hash: hash})
	return existBook, result.Find(&existBook).Error
}
