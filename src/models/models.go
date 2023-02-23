package models

import "time"

type User struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	OAuth_Type string `json:"oAuth_Type"`
}

type Mweet struct {
	Image     string    `json:"img"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}
