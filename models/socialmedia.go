package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey" json:"-"`
	Name           string `json:"name" form:"name" valid:"required~Name of your social media is required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" valid:"required~Social Media URL is required"`
	UserID         uint
	User           *User
	CreatedAt      *time.Time `json:"-,omitempty"`
	UpdatedAt      *time.Time `json:"-,omitempty"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
