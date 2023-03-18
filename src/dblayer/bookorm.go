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
	result := db.Table("bookinfo").Select("id, hash").Where(&models.BookInfo{Hash: hash})
	return existBook, result.Find(&existBook).Error
}

func (db *BookORM) GetAllBook() (*[]models.BookInfo, error) {
	var bookInfoList *[]models.BookInfo
	result := db.Table("bookinfo")
	return bookInfoList, result.Limit(20).Find(&bookInfoList).Error
}

func (db *BookORM) GetBookInfoById(id uint) (models.BookInfo, error) {

	bookInfo := models.BookInfo{}

	result := db.Table("bookinfo").Select("title", "author", "publisher", "release_date", "category").Where("id = ?", id).Find(&bookInfo)

	return bookInfo, result.Error

}

func (db *BookORM) GetBookInfoWithRank(id uint) ([]models.BookRank, error) {

	book := []models.BookRank{}
	db.Table("bookinfo b").Select("br.rank, br.created_at").Joins("inner join bookRank br on b.id = br.book_id").Where("b.id = ?", id).Scan(&book)

	return book, nil
}

func (db *BookORM) GetBookInfoWithPoint(id uint) ([]models.BookPoint, error) {

	book := []models.BookPoint{}
	db.Table("bookinfo b").Select("bp.point, bp.point_rate, bp.created_at").Joins("inner join bookpoint bp on b.id = bp.book_id").Where("b.id = ?", id).Scan(&book)

	return book, nil
}

func (db *BookORM) GetBookInfoWithPrice(id uint) ([]models.BookPrice, error) {

	book := []models.BookPrice{}
	db.Table("bookinfo b").Select("bp.price, bp.discount_price, bp.discount_rate").Joins("inner join bookprice bp on b.id = bp.book_id").Where("b.id = ?", id).Scan(&book)

	return book, nil
}

func (db *BookORM) GetBookInfoWithSummary(id uint) (models.BookSummary, error) {

	book := models.BookSummary{}
	db.Limit(1).Table("bookinfo b").Select("bp.summary").Joins("inner join booksummary bp on b.id = bp.book_id").Where("b.id = ?", id).Order("bp.id desc").Scan(&book)

	return book, nil
}
