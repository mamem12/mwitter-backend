package parsebook

import (
	"fmt"
	"mwitter-backend/src/models"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/robfig/cron"
	"github.com/tidwall/gjson"
)

const (
	CronSpec   = "0 0,30 * * * *"
	apiURL     = "https://product.kyobobook.co.kr/api/gw/pdt/best-seller/online?page=%d&per=20&saleCmdtDvsnCode=KOR&saleCmdtClstCode=01&saleCmdtDsplDvsnCode=KOR&period=002&dsplDvsnCode=001&dsplTrgtDvsnCode=004"
	user_agent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
	MAX_PAGE   = 5
)

type Book struct {
	models.BookInfo
	models.BookRank
	models.BookPrice
	models.BookPoint
	models.BookSummary
}

func ParseRun() {
	cron := cron.New()
	var wg sync.WaitGroup

	wg.Add(1)
	cron.AddFunc(CronSpec, getPage)
	cron.Start()

	wg.Wait()
}

func getPage() {
	var wg sync.WaitGroup
	wg.Add(1)
	pageRecieveChannel := make(chan []gjson.Result, MAX_PAGE)
	contentChannel := make(chan []Book)

	var content []gjson.Result
	for i := 1; i <= 2; i++ {
		go getBookContent(i, pageRecieveChannel)
	}
	for i := 1; i <= 2; i++ {
		contents := <-pageRecieveChannel
		content = append(content, contents...)
	}

	go parsePage(content, contentChannel, &wg)

	close(pageRecieveChannel)

	result := <-contentChannel

	wg.Wait()
	close(contentChannel)
	fmt.Println("wait end", len(result))
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

func parsePage(jsonContents []gjson.Result, contentChannel chan []Book, wg *sync.WaitGroup) {
	var BookList []Book

	for _, jsonBookcontent := range jsonContents {
		bookInfoCh := make(chan *Book)
		bookRankCh := make(chan *Book)
		bookPriceCh := make(chan *Book)
		bookPointCh := make(chan *Book)
		bookSummaryCh := make(chan *Book)

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

func parseBookBasicInfo(jsonBook gjson.Result, bookInfoCh chan *Book) {
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

	book := &Book{
		BookInfo: *bookInfo,
	}

	bookInfoCh <- book

	close(bookInfoCh)
}

func parseBookRank(jsonBook gjson.Result, bookInfoCh chan *Book, bookRankCh chan *Book) {

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

func parseBookPrice(jsonBook gjson.Result, bookRankCh chan *Book, bookPriceCh chan *Book) {
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

func parseBookPoint(jsonBook gjson.Result, bookPriceCh chan *Book, bookPointCh chan *Book) {
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

func parseBookSummary(jsonBook gjson.Result, bookPointCh chan *Book, bookSummaryCh chan *Book) {
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
