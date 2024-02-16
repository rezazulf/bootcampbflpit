package service

import (
	"final-project-h8-mygram-Reza/src/models"
	"final-project-h8-mygram-Reza/src/pkg/errs"
	"final-project-h8-mygram-Reza/src/pkg/helpers"
	"final-project-h8-mygram-Reza/src/repository/photo_repository"
	"final-project-h8-mygram-Reza/src/structs"
	"fmt"
)

type PhotoService interface {
	PostPhoto(userID uint, photoPayload *structs.PhotoRequest) (*structs.PhotoResponse, errs.MessageErr)
	GetAllPhotos() ([]*structs.GetPhotoResponse, errs.MessageErr)
	EditPhotoData(photoID uint, photoPayload *structs.PhotoRequest) (*structs.UpdatePhotoResponse, errs.MessageErr)
	DeletePhoto(photoID uint) (*structs.DeletePhotoResponse, errs.MessageErr)
}

type photoService struct {
	photoRepository photo_repository.PhotoRepository
}

func NewPhotoService(photoRepository photo_repository.PhotoRepository) PhotoService {
	return &photoService{photoRepository: photoRepository}
}

func (p *photoService) PostPhoto(userID uint, photoPayload *structs.PhotoRequest) (*structs.PhotoResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(photoPayload)
	if err != nil {
		return nil, err
	}

	payload := &models.Photo{
		Title:    photoPayload.Title,
		Caption:  photoPayload.Caption,
		PhotoURL: photoPayload.PhotoURL,
		UserID:   userID,
	}

	photo, err := p.photoRepository.PostPhoto(payload)
	if err != nil {
		return nil, err
	}

	response := &structs.PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
	}

	return response, nil
}

func (p *photoService) GetAllPhotos() ([]*structs.GetPhotoResponse, errs.MessageErr) {
	structs := make([]*structs.GetPhotoResponse, 0)
	photos, err := p.photoRepository.GetAllPhotos()
	if err != nil {
		return nil, err
	}

	for _, photo := range photos {
		structs = append(structs, photo.ToGetPhotoResponseStructs())
	}

	return structs, nil
}

func (p *photoService) EditPhotoData(photoID uint, photoPayload *structs.PhotoRequest) (*structs.UpdatePhotoResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(photoPayload)
	if err != nil {
		return nil, err
	}

	payload := &models.Photo{
		Title:    photoPayload.Title,
		Caption:  photoPayload.Caption,
		PhotoURL: photoPayload.PhotoURL,
	}

	photo, err := p.photoRepository.EditPhotoData(photoID, payload)
	if err != nil {
		return nil, err
	}

	response := &structs.UpdatePhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}

	return response, nil
}

func (p *photoService) DeletePhoto(photoID uint) (*structs.DeletePhotoResponse, errs.MessageErr) {
	_, err := p.photoRepository.GetPhotoByID(photoID)
	if err != nil {
		return nil, err
	}

	err = p.photoRepository.DeletePhoto(photoID)
	if err != nil {
		return nil, err
	}

	response := &structs.DeletePhotoResponse{
		Message: "Your photo has been deleted",
	}

	fmt.Println("Melihat response di service: ", response)
	return response, nil
}
