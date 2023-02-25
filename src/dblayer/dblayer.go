package dblayer

import "mwitter-backend/src/models"

type DBLayer interface {
	GetAllMweeter() ([]models.Mweet, error)
	CreateMweet(*models.Mweet) error
	UpdateMweet()
	DeleteMweet()
	GetMweeterById()
	CreateUser(user *models.User) error
	SignInUser(email, password string) (*models.User, error)
	SelectUserByEmail(email string) (*models.User, error)
	SignOutUser()
	UpdateProfile(id string, UpdateInfo *models.User) error
}
