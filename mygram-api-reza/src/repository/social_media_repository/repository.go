package social_media_repository

import (
	"final-project-h8-mygram-Reza/src/models"
	"final-project-h8-mygram-Reza/src/pkg/errs"
)

type SocialMediaRepository interface {
	AddSocialMedia(socialMedia *models.SocialMedia) (*models.SocialMedia, errs.MessageErr)
	GetAllSocialMedias() ([]*models.SocialMedia, errs.MessageErr)
	GetSocialMediaByID(socialMediaID uint) (*models.SocialMedia, errs.MessageErr)
	EditSocialMediaData(socialMediaID uint, socialMedia *models.SocialMedia) (*models.SocialMedia, errs.MessageErr)
	DeleteSocialMedia(socialMediaID uint) errs.MessageErr
}
