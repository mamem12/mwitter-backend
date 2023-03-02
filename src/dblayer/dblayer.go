package dblayer

import "mwitter-backend/src/models"

type DBLayer interface {
	CreateUser(user *models.User) error
	SignInUser(email, password string) (*models.User, error)
	SelectUserByEmail(email string) (*models.User, error)
	UpdateProfile(id string, UpdateInfo *models.User) error
}
