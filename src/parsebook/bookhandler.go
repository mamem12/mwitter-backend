package parsebook

import (
	"encoding/json"
	"mwitter-backend/src/common"
	"mwitter-backend/src/dblayer"
	"mwitter-backend/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HandlerInterface interface {
	GetBookList(ctx *gin.Context)
	GetBookInfoAll(ctx *gin.Context)
	GetBookRank(ctx *gin.Context)
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

func NewBookInsertHandler(books ...models.Book) error {

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
		var id uint
		if existBook.Hash == "" {

			err = db.InsertBookInfo(&book.BookInfo)
			id = book.BookInfo.ID
		} else {

			err = db.UpdateBookInfo(&book.BookInfo)
			if err != nil {
				continue
			}

			id = existBook.ID

		}

		if err != nil {
			return err
		}

		// id := book.BookInfo.ID

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

// get all, limit 20
func (book *BookHandler) GetBookList(ctx *gin.Context) {
	bookList, err := book.db.GetAllBook()

	if err != nil {
		common.HandleErrorWithResponse(err.Error(), http.StatusInternalServerError, ctx)
	}

	ctx.JSON(http.StatusOK, bookList)
}

func (book *BookHandler) GetBookInfoAll(ctx *gin.Context) {

	idParm := ctx.Param("id")
	id, err := strconv.Atoi(idParm)

	if err != nil {
		common.HandleErrorWithResponse("id not valid", http.StatusBadGateway, ctx)
		return
	}

	bookInfoResult := make(map[string]any)

	bookInfo, _ := book.db.GetBookInfoById(uint(id))

	bookInfoResult["bookinfo"] = bookInfo

	bookInfoRankList, _ := book.db.GetBookInfoWithRank(uint(id))

	bookInfoResult["rank"] = bookInfoRankList

	bookPointList, _ := book.db.GetBookInfoWithPoint(uint(id))

	bookInfoResult["point"] = bookPointList

	bookPriceList, _ := book.db.GetBookInfoWithPrice(uint(id))

	bookInfoResult["price"] = bookPriceList

	bookSummary, _ := book.db.GetBookInfoWithSummary(uint(id))

	bookInfoResult["summary"] = bookSummary.Summary

	data, _ := json.Marshal(bookInfoResult)

	ctx.JSON(http.StatusOK, string(data))
}

// get one with rank list
func (book *BookHandler) GetBookRank(ctx *gin.Context) {
	idParm := ctx.Param("id")
	id, err := strconv.Atoi(idParm)

	if err != nil {
		common.HandleErrorWithResponse("id not valid", http.StatusBadGateway, ctx)
		return
	}

	bookRank := make(map[string]any)

	bookInfo, _ := book.db.GetBookInfoById(uint(id))

	bookRank["bookinfo"] = bookInfo

	bookInfoRankList, _ := book.db.GetBookInfoWithRank(uint(id))

	bookRank["rank"] = bookInfoRankList

	data, _ := json.Marshal(bookRank)

	ctx.JSON(http.StatusOK, string(data))
}

// get one by price & point list
