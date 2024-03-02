package model

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base model
type Base struct {
	ID         *uuid.UUID       `json:"id,omitempty" gorm:"primaryKey;unique;type:varchar(36);not null" format:"uuid" swaggerignore:"true"` // model ID
	Sort       *int64           `json:"sort,omitempty" gorm:"default:0" swaggerignore:"true"`                                               // sort (increment)
	Status     *int             `json:"status,omitempty" gorm:"type:smallint;default:1" example:"1" swaggerignore:"true"`                   // status (0: deleted, 1: active, 2: draft. 3: blocked, 4: canceled)
	CreatedAt  *strfmt.DateTime `json:"created_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggerignore:"true"`               // created at automatically inserted on post
	UpdatedAt  *strfmt.DateTime `json:"updated_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggerignore:"true"`               // updated at automatically changed on put or add on post
	SyncedAt   *strfmt.DateTime `json:"synced_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggerignore:"true"`                // synced at automatically changed on syncronize post
	DeletedAt  gorm.DeletedAt   `json:"-" gorm:"index" swaggerignore:"true"`
	Additional *string          `json:"additional,omitempty"`
}
type BaseInt struct {
	ID         *int64           `json:"id,omitempty" gorm:"primaryKey;unique;not null"`                                       // model ID
	Sort       *int64           `json:"sort,omitempty" gorm:"default:0" swaggerignore:"true"`                                 // sort (increment)
	Status     *int             `json:"status,omitempty" gorm:"type:smallint;default:1" example:"1" swaggerignore:"true"`     // status (0: deleted, 1: active, 2: draft. 3: blocked, 4: canceled)
	CreatedAt  *strfmt.DateTime `json:"created_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggerignore:"true"` // created at automatically inserted on post
	UpdatedAt  *strfmt.DateTime `json:"updated_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggerignore:"true"` // updated at automatically changed on put or add on post
	SyncedAt   *strfmt.DateTime `json:"synced_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggerignore:"true"`  // synced at automatically changed on syncronize post
	DeletedAt  gorm.DeletedAt   `json:"-" gorm:"index" swaggerignore:"true"`
	Additional *string          `json:"additional,omitempty"`
}

// BeforeCreate Data
func (b *Base) BeforeCreate(tx *gorm.DB) (e error) {
	now := strfmt.DateTime(time.Now())

	if nil == b.ID {
		var id uuid.UUID
		id, e = uuid.NewRandom()
		b.ID = &id
	}

	if b.CreatedAt == nil {
		b.CreatedAt = &now
	}
	if b.UpdatedAt == nil {
		b.UpdatedAt = &now
	}
	if b.SyncedAt == nil {
		b.SyncedAt = &now
	}

	return e
}

// BeforeUpdate Data
func (b *Base) BeforeUpdate(tx *gorm.DB) error {
	now := strfmt.DateTime(time.Now())
	b.UpdatedAt = &now
	return nil
}
