package parsebook

import (
	"fmt"
	"log"
	"mwitter-backend/src/models"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/robfig/cron"
	"github.com/tidwall/gjson"
)

const (
	CronSpec = "0 0,52 * * * *"
	apiURL   = "https://product.kyobobook.co.kr/api/gw/pdt/best-seller/online?page=%d&per=20&saleCmdtDvsnCode=KOR&saleCmdtClstCode=01&saleCmdtDsplDvsnCode=KOR&period=002&dsplDvsnCode=001&dsplTrgtDvsnCode=004"
	MAX_PAGE = 2
)

func ParseRun() {
	cron := cron.New()
	var wg sync.WaitGroup

	wg.Add(1)
	cron.AddFunc(CronSpec, getPage)
	cron.Start()

	wg.Wait()
}

func getPage() {
	log.Println("cron start")
	var wg sync.WaitGroup
	wg.Add(1)
	pageRecieveChannel := make(chan []gjson.Result, MAX_PAGE)
	contentChannel := make(chan []models.Book)

	var content []gjson.Result
	for i := 1; i <= MAX_PAGE; i++ {
		go getBookContent(i, pageRecieveChannel)
	}

	for i := 1; i <= MAX_PAGE; i++ {
		contents := <-pageRecieveChannel
		content = append(content, contents...)
	}

	go parseBook(content, contentChannel, &wg)

	close(pageRecieveChannel)

	result := <-contentChannel

	wg.Wait()
	close(contentChannel)

	fmt.Println("wait end", len(result))

	NewBookInsertHandler(result...)
}

func getBookContent(pageNum int, pageRecieveChannel chan<- []gjson.Result) {

	content := []gjson.Result{}

	res, err := http.Get(fmt.Sprintf(apiURL, pageNum))

	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		fmt.Println(err.Error())
	}

	result := gjson.Get(doc.Text(), "data")

	result = result.Get("bestSeller")

	content = append(content, result.Array()...)

	pageRecieveChannel <- content
}

func parseBook(jsonContents []gjson.Result, contentChannel chan []models.Book, wg *sync.WaitGroup) {
	var BookList []models.Book

	for _, jsonBookcontent := range jsonContents {
		bookInfoCh := make(chan *models.Book)
		bookRankCh := make(chan *models.Book)
		bookPriceCh := make(chan *models.Book)
		bookPointCh := make(chan *models.Book)
		bookSummaryCh := make(chan *models.Book)

		go parseBookBasicInfo(jsonBookcontent, bookInfoCh)
		go parseBookRank(jsonBookcontent, bookInfoCh, bookRankCh)
		go parseBookPrice(jsonBookcontent, bookRankCh, bookPriceCh)
		go parseBookPoint(jsonBookcontent, bookPriceCh, bookPointCh)
		go parseBookSummary(jsonBookcontent, bookPointCh, bookSummaryCh)
		book := <-bookSummaryCh

		BookList = append(BookList, *book)
	}

	contentChannel <- BookList
	wg.Done()
}

func parseBookBasicInfo(jsonBook gjson.Result, bookInfoCh chan *models.Book) {
	title := jsonBook.Get("cmdtName")
	author := jsonBook.Get("chrcName")
	publisher := jsonBook.Get("pbcmName")
	releaseDate := jsonBook.Get("rlseDate")
	category := jsonBook.Get("saleCmdtClstName")

	bookInfo := &models.BookInfo{
		Title:       title.String(),
		Author:      author.String(),
		Publisher:   publisher.String(),
		ReleaseDate: releaseDate.String(),
		Category:    category.String(),
	}

	book := &models.Book{
		BookInfo: *bookInfo,
	}

	bookInfoCh <- book

	close(bookInfoCh)
}

func parseBookRank(jsonBook gjson.Result, bookInfoCh chan *models.Book, bookRankCh chan *models.Book) {

	for bookInfo := range bookInfoCh {
		rank := jsonBook.Get("prstRnkn")

		bookRank := &models.BookRank{
			Rank: rank.String(),
		}
		bookInfo.BookRank = *bookRank
		bookRankCh <- bookInfo
	}
	close(bookRankCh)
}

func parseBookPrice(jsonBook gjson.Result, bookRankCh chan *models.Book, bookPriceCh chan *models.Book) {
	for bookInfo := range bookRankCh {
		price := jsonBook.Get("price")
		discountPrice := jsonBook.Get("sapr")
		discountRate := jsonBook.Get("dscnRate")

		bookPrice := &models.BookPrice{
			Price:         price.String(),
			DiscountPrice: discountPrice.String(),
			DiscountRate:  discountRate.String(),
		}

		bookInfo.BookPrice = *bookPrice

		bookPriceCh <- bookInfo
	}
	close(bookPriceCh)
}

func parseBookPoint(jsonBook gjson.Result, bookPriceCh chan *models.Book, bookPointCh chan *models.Book) {
	for bookInfo := range bookPriceCh {
		point := jsonBook.Get("upntAcmlAmnt")
		pointRate := jsonBook.Get("upntAcmlRate")

		bookPoint := &models.BookPoint{
			Point:     point.String(),
			PointRate: pointRate.String(),
		}

		bookInfo.BookPoint = *bookPoint
		bookPointCh <- bookInfo
	}
	close(bookPointCh)
}

func parseBookSummary(jsonBook gjson.Result, bookPointCh chan *models.Book, bookSummaryCh chan *models.Book) {
	for bookInfo := range bookPointCh {
		summary := jsonBook.Get("inbukCntt")

		bookSummary := &models.BookSummary{
			Summary: summary.String(),
		}

		bookInfo.BookSummary = *bookSummary
		bookSummaryCh <- bookInfo
	}
	close(bookSummaryCh)
}
