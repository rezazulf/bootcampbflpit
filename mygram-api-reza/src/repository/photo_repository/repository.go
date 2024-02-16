package photo_repository

import (
	"final-project-h8-mygram-Reza/src/models"
	"final-project-h8-mygram-Reza/src/pkg/errs"
)

type PhotoRepository interface {
	PostPhoto(photo *models.Photo) (*models.Photo, errs.MessageErr)
	GetAllPhotos() ([]*models.Photo, errs.MessageErr)
	GetPhotoByID(photoID uint) (*models.Photo, errs.MessageErr)
	EditPhotoData(photoID uint, photo *models.Photo) (*models.Photo, errs.MessageErr)
	DeletePhoto(photoID uint) errs.MessageErr
}
