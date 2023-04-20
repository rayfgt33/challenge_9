package model

import (
	"moddleware/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Usergo struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your full name is required"` // Valid buat kasih valdasi dan pesannya
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required"`
	Level    string    `gorm:"" json:"level" form:"level"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products" form:"products"`
}

// Hooks
func (u *Usergo) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	u.Password = helpers.HashPass(u.Password)
	return
}
