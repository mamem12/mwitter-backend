package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	gorm.Model
}

type BookInfo struct {
	ID          uint   `json:"id",gorm:"primarykey"`
	Title       string `json:"title"`
	Author      string `json:"Author"`
	Publisher   string `json:"publisher"`
	ReleaseDate string `json:"relesaseDate"`
	Category    string `json:"category"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BookRank struct {
	ID        uint   `json:"id",gorm:"primarykey"`
	BookId    uint   `gorm:"index",json:"bookId"`
	Rank      string `json:"rank"`
	CreatedAt time.Time
}

type BookPrice struct {
	ID            uint   `json:"id",gorm:"primarykey"`
	BookId        uint   `gorm:"index",json:"bookId"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discountPrice"`
	DiscountRate  string `json:"discountRate"`
	CreatedAt     time.Time
}

type BookPoint struct {
	ID        uint   `json:"id",gorm:"primarykey"`
	BookId    uint   `json:"bookId",gorm:"index"`
	Point     string `json:"point"`
	PointRate string `json:"pointRate"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BookSummary struct {
	ID        uint   `json:"id",gorm:"primarykey"`
	BookId    uint   `json:"bookId",gorm:"index"`
	Summary   string `json:"Summary"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "users"
}

func (BookInfo) TableName() string {
	return "bookinfo"
}

func (BookRank) TableName() string {
	return "bookRank"
}

func (BookPrice) TableName() string {
	return "bookprice"
}

func (BookPoint) TableName() string {
	return "bookpoint"
}

func (BookSummary) TableName() string {
	return "booksummary"
}
