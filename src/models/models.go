package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	BookInfo
	BookRank
	BookPrice
	BookPoint
	BookSummary
}

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	gorm.Model
}

type BookInfo struct {
	ID          uint    `json:"id",gorm:"primarykey"`
	Title       string  `gorm:"type:varchar(100)",json:"title"`
	Author      string  `json:"Author"`
	Publisher   string  `json:"publisher"`
	ReleaseDate string  `json:"relesaseDate"`
	Category    string  `json:"category"`
	ReviewCnt   uint    `json:"review"`
	Avg         float32 `json:"avg"`
	Hash        string  `gorm:"index:unique"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BookRank struct {
	ID        uint `json:"id",gorm:"primarykey"`
	BookId    uint `gorm:"index",json:"bookId"`
	Rank      uint `json:"rank"`
	CreatedAt time.Time
}

type BookPrice struct {
	ID            uint   `json:"id",gorm:"primarykey"`
	BookId        uint   `gorm:"index",json:"bookId"`
	Price         uint   `json:"price"`
	DiscountPrice uint   `json:"discountPrice"`
	DiscountRate  string `json:"discountRate"`
	CreatedAt     time.Time
}

type BookPoint struct {
	ID        uint   `json:"id",gorm:"primarykey"`
	BookId    uint   `gorm:"index",json:"bookId"`
	Point     string `json:"point"`
	PointRate string `json:"pointRate"`
	CreatedAt time.Time
}

type BookSummary struct {
	ID        uint   `json:"id",gorm:"primarykey"`
	BookId    uint   `gorm:"index",json:"bookId"`
	Summary   string `json:"Summary"`
	Hash      string `gorm:"type:varchar(100)"`
	CreatedAt time.Time
}

func (User) TableName() string {
	return "users"
}

func (BookInfo) TableName() string {
	return "book_info"
}

func (BookRank) TableName() string {
	return "book_rank"
}

func (BookPrice) TableName() string {
	return "book_price"
}

func (BookPoint) TableName() string {
	return "book_point"
}

func (BookSummary) TableName() string {
	return "book_summary"
}
