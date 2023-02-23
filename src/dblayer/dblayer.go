package dblayer

import "mwitter-backend/src/models"

type DBLayer interface {
	GetAllMweeter() ([]models.Mweet, error)
	CreateMweet(*models.User) error
	UpdateMweet()
	DeleteMweet()
	GetMweeterById()
	CreateUser()
	SignInUser()
	SignOutUser()
	UpdateProfile()
}
