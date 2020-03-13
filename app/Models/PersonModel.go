package Models

import (
	"github.com/jinzhu/gorm"
)

type PersonModel struct {
	gorm.Model
	Username string `gorm:"unique;not null;type:varchar(20)"`
	Full_name string `gorm:"size:40;not null;type:varchar(40)"`
	Password string `gorm:"not null"`
	Email string `gorm:"type:varchar(100);unique_index;not null"`
}

