package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title     string `gorm:"not null" json:"title" form:"title" valid:"required~The title is required"`
	Caption   string `gorm:"not null" json:"caption" form:"caption"`
	Photo_url string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~The photo url is required"`
	UserID    uint
	User      *User
	Comment   []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comment"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
