package models

import (
	"time"
)

// Base is the collection of fields all types in the Skillbased architecture have
type Base struct {
	ID        string    `gorm:"type:varchar(27);primaryKey;" json:"id"`
	Name      string    `gorm:"type:varchar(50);unique" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	IsSeed bool `json:"-"`
}

var SkillsList = []string{
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
