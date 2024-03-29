package photo_pg

import (
	"final-project-h8-mygram-Reza/src/models"
	"final-project-h8-mygram-Reza/src/pkg/errs"
	"final-project-h8-mygram-Reza/src/repository/photo_repository"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type photoPG struct {
	db *gorm.DB
}

func NewPhotoPG(db *gorm.DB) photo_repository.PhotoRepository {
	return &photoPG{db: db}
}

func (p *photoPG) PostPhoto(photoPayload *models.Photo) (*models.Photo, errs.MessageErr) {
	photo := models.Photo{}

	err := p.db.Model(photo).Create(&photoPayload).Error
	if err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	err = p.db.Last(&photo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("Photo not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &photo, nil
}

func (p *photoPG) GetAllPhotos() ([]*models.Photo, errs.MessageErr) {
	photos := []*models.Photo{}

	if err := p.db.Preload("User").Find(&photos).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return photos, nil
}

func (p *photoPG) GetPhotoByID(photoID uint) (*models.Photo, errs.MessageErr) {
	photo := models.Photo{}

	err := p.db.Model(photo).Where("id = ?", photoID).First(&photo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("Photo not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &photo, nil
}

func (p *photoPG) EditPhotoData(photoID uint, photoPayload *models.Photo) (*models.Photo, errs.MessageErr) {
	photo := models.Photo{}

	err := p.db.Raw("Update photos SET title = ?, caption = ?, photo_url = ?, updated_at = ? WHERE id = ? RETURNING id, title, caption, user_id, photo_url, updated_at", photoPayload.Title, photoPayload.Caption, photoPayload.PhotoURL, time.Now(), photoID).Scan(&photo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("Photo not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}
	fmt.Println("Melihat photo", &photo)

	return &photo, nil
}

func (p *photoPG) DeletePhoto(photoID uint) errs.MessageErr {
	photo := models.Photo{}

	err := p.db.Debug().Model(photo).Where("id = ?", photoID).Delete(&photo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errs.NewNotFoundError("Photo not found")
		}
		return errs.NewInternalServerErrorr("Something went wrong")
	}

	return nil
}
