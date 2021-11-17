package models

import (
	"time"

	"github.com/rs/xid"
)

// Base is the collection of fields all types in the Skillbased architecture have
type Base struct {
	ID        xid.ID     `gorm:"type:varchar(20);primary_key;" json:"id"`
	Name      string     `gorm:"type:varchar(50);unique" json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

var skillsList = []string{
	"handling",
	"power",
	"speed",
	"height",
	"stamina",
}

//// BeforeCreate will set a UUID rather than numeric ID.
//func (base *Base) BeforeCreate(scope *gorm.Scope) error {
//	return scope.SetColumn("ID", xid.New())
//}
