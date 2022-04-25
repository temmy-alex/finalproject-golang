package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uint   `gorm:"primaryKey" json:"-"`
	Message   string `json:"title" form:"message" valid:"required~Message of your comment is required"`
	UserID    uint
	PhotoID   uint `json:"photo_id" form:"photo_id"`
	User      *User
	Photo     *Photo
	CreatedAt *time.Time `json:"-,omitempty"`
	UpdatedAt *time.Time `json:"-,omitempty"`
}

func (ct *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(ct)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (ct *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(ct)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
