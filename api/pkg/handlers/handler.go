package handlers

import (
	"github.com/jinzhu/gorm"
)

type (
	Handler struct {
		DB *gorm.DB
	}
)
