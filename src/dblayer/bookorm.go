package dblayer

import (
	"mwitter-backend/src/models"

	"gorm.io/gorm"
)

type BookORM struct {
	*gorm.DB
}

func (db *BookORM) GetBookInfoByHash(hash string) (existBook *models.BookInfo, err error) {
	result := db.Table("book_info").Select("id, hash").Where(&models.BookInfo{Hash: hash})
	return existBook, result.Find(&existBook).Error
}

func (db *BookORM) GetAllBook(per int, pageIdx int, sort string) (*[]models.BookInfo, error) {
	var bookInfoList *[]models.BookInfo

	result := db.Table("book_info info").Select("info.*")
	result.Limit(per).Offset(pageIdx)

	if sort == "price" {

		result.Joins("inner join book_price bp on info.id = bp.book_id")
		result.Where("bp.created_at >= curdate()")
		result.Order("bp.price, id")

	} else {

		result.Joins("inner join book_rank br on info.id = br.book_id")
		result.Where("br.created_at >= curdate()")

		if sort == "rank" {

			result.Order("br.rank")

		} else if sort == "review" {

			result.Order("info.review_cnt desc")

		} else if sort == "avg" {

			result.Order("info.avg desc")

		}

	}

	return bookInfoList, result.Find(&bookInfoList).Error
}

func (db *BookORM) GetBookInfoById(id uint) (models.BookInfo, error) {

	bookInfo := models.BookInfo{}

	result := db.Table("book_info").Select("title", "author", "publisher", "release_date", "category, review_cnt, avg").Where("id = ?", id).Find(&bookInfo)

	return bookInfo, result.Error
}

func (db *BookORM) GetBookInfoWithRank(id uint) ([]models.BookRank, error) {

	book := []models.BookRank{}
	db.Table("book_info b").Select("br.rank, br.created_at").Joins("inner join book_rank br on b.id = br.book_id").Where("b.id = ?", id).Scan(&book)

	return book, nil
}

func (db *BookORM) GetBookInfoWithPoint(id uint) ([]models.BookPoint, error) {

	book := []models.BookPoint{}
	db.Table("book_info b").Select("bp.point, bp.point_rate, bp.created_at").Joins("inner join book_point bp on b.id = bp.book_id").Where("b.id = ?", id).Scan(&book)

	return book, nil
}

func (db *BookORM) GetBookInfoWithPrice(id uint) ([]models.BookPrice, error) {

	book := []models.BookPrice{}
	db.Table("book_info b").Select("bp.price, bp.discount_price, bp.discount_rate").Joins("inner join book_price bp on b.id = bp.book_id").Where("b.id = ?", id).Scan(&book)

	return book, nil
}

func (db *BookORM) GetBookInfoWithSummary(id uint) (models.BookSummary, error) {

	book := models.BookSummary{}
	db.Limit(1).Table("book_info b").Select("bp.summary").Joins("inner join book_summary bp on b.id = bp.book_id").Where("b.id = ?", id).Order("bp.id desc").Scan(&book)

	return book, nil
}
