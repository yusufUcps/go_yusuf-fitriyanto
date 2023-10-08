package model

import (
	"github.com/jinzhu/gorm"

)

type Blog struct {
	gorm.Model
	UserID     uint `json:"userid" form:"userid"`
	Judul    string `json:"judul" form:"judul"`
	Konten string `json:"konten" form:"konten"`
	User    User   `gorm:"foreignkey:UserID"`
}