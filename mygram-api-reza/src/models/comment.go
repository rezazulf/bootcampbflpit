package models

import "final-project-h8-mygram-Reza/src/structs"

type Comment struct {
	GormModel
	UserID  uint
	PhotoID uint
	Message string `gorm:"not null" form:"message" json:"message" valid:"required~Message is required"`
	User    *User
	Photo   *Photo
}

func (c *Comment) ToGetCommentResponseStructs() *structs.GetCommentResponse {
	return &structs.GetCommentResponse{
		ID:        c.ID,
		Message:   c.Message,
		PhotoID:   c.PhotoID,
		UserID:    c.UserID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		User: structs.EmbeddedUserResponse{
			Username: c.User.Username,
			Email:    c.User.Email,
		},
		Photo: structs.EmbeddedPhotoResponse{
			ID:       c.Photo.ID,
			Title:    c.Photo.Title,
			Caption:  c.Photo.Caption,
			PhotoURL: c.Photo.PhotoURL,
			UserID:   c.Photo.UserID,
		},
	}
}
