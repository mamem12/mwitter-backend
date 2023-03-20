package handler

import (
	"encoding/json"
	"log"
	"mwitter-backend/src/common"
	"mwitter-backend/src/dblayer"
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

// func NewBookInsertHandler(books ...models.Book) error {

// 	db, err := dblayer.NewORM("test", gorm.Config{})

// 	if err != nil {
// 		return err
// 	}

// 	for _, book := range books {
// 		title := book.BookInfo.Title
// 		author := book.BookInfo.Author
// 		publisher := book.BookInfo.Publisher

// 		hash, _ := common.StrToHash(title + author + publisher)

// 		book.BookInfo.Hash = hash

// 		existBook, _ := db.GetBookInfoByHash(hash)
// 		var id uint
// 		if existBook.Hash == "" {

// 			err = db.InsertBookInfo(&book.BookInfo)
// 			id = book.BookInfo.ID
// 		} else {

// 			err = db.UpdateBookInfo(&book.BookInfo)
// 			if err != nil {
// 				continue
// 			}

// 			id = existBook.ID

// 		}

// 		if err != nil {
// 			return err
// 		}

// 		// id := book.BookInfo.ID

// 		book.BookRank.BookId = id
// 		book.BookPrice.BookId = id
// 		book.BookPoint.BookId = id
// 		book.BookSummary.BookId = id

// 		db.InsertBookRank(&book.BookRank)
// 		db.InsertBookPrice(&book.BookPrice)
// 		db.InsertBookPoint(&book.BookPoint)
// 		db.InsertBookSummary(&book.BookSummary)
// 	}

// 	return nil
// }

// per - 20개씩 (기본), 50개씩(옵션) √
// page - 하나씩 증가 가능 √
// sort - √
// 	정렬 ID 순(rank순),
//	리뷰순,
// 	평점순, => info, rank 조인하여 내부 처리
//	가격순 => info, price 조인하여 내부 처리

func (book *BookHandler) GetBookList(ctx *gin.Context) {

	sortParam := ctx.DefaultQuery("sort", "rank")
	perParam := ctx.DefaultQuery("per", "20")
	pageParam := ctx.DefaultQuery("page", "1")

	per, err := strconv.Atoi(perParam)

	if err != nil {
		per = 20
	} else if per > 50 {
		per = 50
	}

	page, err := strconv.Atoi(pageParam)

	if err != nil {
		page = 1
	}

	pageIdx := (per * page) - per

	log.Printf("sort : %s, per : %d, page : %d\n", sortParam, per, page)

	bookList, err := book.db.GetAllBook(per, pageIdx, sortParam)

	if err != nil {
		common.HandleErrorWithResponse(err.Error(), http.StatusInternalServerError, ctx)
		return
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
