package service

import (
	"final-project-h8-mygram-Reza/src/models"
	"final-project-h8-mygram-Reza/src/pkg/errs"
	"final-project-h8-mygram-Reza/src/pkg/helpers"
	"final-project-h8-mygram-Reza/src/repository/social_media_repository"
	"final-project-h8-mygram-Reza/src/structs"
)

type SocialMediaService interface {
	AddSocialMedia(userID uint, socialMediaPayload *structs.SocialMediaRequest) (*structs.SocialMediaResponse, errs.MessageErr)
	GetAllSocialMedias() ([]*structs.GetSocialMediaResponse, errs.MessageErr)
	EditSocialMediaData(socialMediaID uint, socialMediaPayload *structs.SocialMediaRequest) (*structs.UpdateSocialMediaResponse, errs.MessageErr)
	DeleteSocialMedia(socialMediaID uint) (*structs.DeleteSocialMediaResponse, errs.MessageErr)
}

type socialMediaService struct {
	socialMediaRepository social_media_repository.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepository social_media_repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{socialMediaRepository: socialMediaRepository}
}

func (s *socialMediaService) AddSocialMedia(userID uint, socialMediaPayload *structs.SocialMediaRequest) (*structs.SocialMediaResponse, errs.MessageErr) {
	if err := helpers.ValidateStruct(socialMediaPayload); err != nil {
		return nil, err
	}

	models := models.SocialMedia{
		Name:           socialMediaPayload.Name,
		SocialMediaURL: socialMediaPayload.SocialMediaURL,
		UserID:         userID,
	}

	socialMedia, err := s.socialMediaRepository.AddSocialMedia(&models)
	if err != nil {
		return nil, err
	}

	response := &structs.SocialMediaResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         socialMedia.UserID,
		CreatedAt:      socialMedia.CreatedAt,
	}

	return response, nil
}

func (s *socialMediaService) GetAllSocialMedias() ([]*structs.GetSocialMediaResponse, errs.MessageErr) {
	structs := make([]*structs.GetSocialMediaResponse, 0)
	socialmedias, err := s.socialMediaRepository.GetAllSocialMedias()
	if err != nil {
		return nil, err
	}

	for _, socialmedia := range socialmedias {
		structs = append(structs, socialmedia.ToGetSocialMediaResponseStructs())
	}

	return structs, nil
}

func (s *socialMediaService) EditSocialMediaData(socialMediaID uint, socialMediaPayload *structs.SocialMediaRequest) (*structs.UpdateSocialMediaResponse, errs.MessageErr) {
	if err := helpers.ValidateStruct(socialMediaPayload); err != nil {
		return nil, err
	}

	models := models.SocialMedia{
		Name:           socialMediaPayload.Name,
		SocialMediaURL: socialMediaPayload.SocialMediaURL,
	}

	socialMedia, err := s.socialMediaRepository.EditSocialMediaData(socialMediaID, &models)
	if err != nil {
		return nil, err
	}

	response := &structs.UpdateSocialMediaResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         socialMedia.UserID,
		UpdatedAt:      socialMedia.UpdatedAt,
	}

	return response, nil
}

func (s *socialMediaService) DeleteSocialMedia(socialMediaID uint) (*structs.DeleteSocialMediaResponse, errs.MessageErr) {
	if err := s.socialMediaRepository.DeleteSocialMedia(socialMediaID); err != nil {
		return nil, err
	}

	response := &structs.DeleteSocialMediaResponse{
		Message: "Your social media has been deleted",
	}

	return response, nil
}
