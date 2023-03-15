package parsebook

import (
	"mwitter-backend/src/common"
	"mwitter-backend/src/dblayer"

	"gorm.io/gorm"
)

type HandlerInterface interface {
}

type HandlerParseInterface interface {
}

type BookHandler struct {
	db dblayer.BookLayer
}

func NewBookHandler() (HandlerInterface, error) {
	db, err := dblayer.NewORM("test", gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &BookHandler{
		db: db.BookORM,
	}, nil
}

func NewBookInsertHandler(books ...Book) error {

	db, err := dblayer.NewORM("test", gorm.Config{})

	if err != nil {
		return err
	}

	for _, book := range books {
		title := book.BookInfo.Title
		author := book.BookInfo.Author
		publisher := book.BookInfo.Publisher

		hash, _ := common.StrToHash(title + author + publisher)

		book.BookInfo.Hash = hash

		existBook, _ := db.GetBookInfoByHash(hash)

		if existBook.Hash == "" {

			err = db.InsertBookInfo(&book.BookInfo)
		} else {
			err = db.UpdateBookInfo(&book.BookInfo)
			if err != nil {
				continue
			}
		}

		if err != nil {
			return err
		}

		id := book.BookInfo.ID

		book.BookRank.BookId = id
		book.BookPrice.BookId = id
		book.BookPoint.BookId = id
		book.BookSummary.BookId = id

		db.InsertBookRank(&book.BookRank)
		db.InsertBookPrice(&book.BookPrice)
		db.InsertBookPoint(&book.BookPoint)
		db.InsertBookSummary(&book.BookSummary)
	}

	return nil
}
