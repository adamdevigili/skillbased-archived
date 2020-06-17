package models

import (
	"time"

	"github.com/rs/xid"
)

type Base struct {
	ID        xid.ID    `gorm:"type:varchar(20);primary_key;" json:"id"`
	Name      string    `gorm:"type:varchar(50)" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

//// BeforeCreate will set a UUID rather than numeric ID.
//func (base *Base) BeforeCreate(scope *gorm.Scope) error {
//	return scope.SetColumn("ID", xid.New())
//}
