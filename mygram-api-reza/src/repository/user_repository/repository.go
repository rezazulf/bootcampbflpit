package user_repository

import (
	"final-project-h8-mygram-Reza/src/models"
	"final-project-h8-mygram-Reza/src/pkg/errs"
)

type UserRepository interface {
	Login(user *models.User) (*models.User, errs.MessageErr)
	Register(user *models.User) (*models.User, errs.MessageErr)
	GetUserByIDAndEmail(user *models.User) (*models.User, errs.MessageErr)
	UpdateUserData(userId uint, user *models.User) (*models.User, errs.MessageErr)
	DeleteUser(userId uint) errs.MessageErr
}
