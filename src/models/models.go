package models

import (
	"gorm.io/gorm"
)

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	gorm.Model
}

type BookInfo struct {
	Title       string `json:"title"`
	Author      string `json:"Author"`
	Publisher   string `json:"publisher"`
	ReleaseDate string `json:"relesaseDate"`
	Category    string `json:"category"`
}

type BookRank struct {
	ID     uint `json:"id"`
	BookId uint `json:"bookId"`
	Rank   uint `json:"rank"`
}

type BookPrice struct {
	ID            uint `json:"id"`
	BookId        uint `json:"bookId"`
	Price         int  `json:"price"`
	DiscountPrice int  `json:"discountPrice"`
	DiscountRate  int  `json:"discountRate"`
}

type BookPoint struct {
	ID        uint `json:"id"`
	Point     int  `json:"point"`
	PointRate int  `json:"pointRate"`
}

type BookData struct {
	BestSeller string `json:"bestSeller"`
}

type BookSummary struct {
	ID      uint   `json:"id"`
	Summary string `json:"Summary"`
}

func (User) TableName() string {
	return "users"
}
