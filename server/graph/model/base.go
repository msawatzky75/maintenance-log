package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BaseWithTimestamps struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *BaseWithTimestamps) BeforeCreate(tx *gorm.DB) error {
	base.ID = uuid.NewV4()
	return nil
}

type Base struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.ID = uuid.NewV4()
	return nil
}
