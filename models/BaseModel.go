package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type Base struct {
	Uuid      uuid.UUID `json:"uuid" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

func GenerateTimeStamps() string {
	return time.Now().Format("2006-01-02T15:04:05.999Z07:00")
}

// BeforeCreate - sets Base struct before every insert
func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	//Create new random UUIDS
	base.Uuid = uuid.New()
	//TIMESTAMPS
	t := GenerateTimeStamps()
	base.CreatedAt, base.UpdatedAt = t, t
	return
}

// AfterUpdate will update the Base struct after every update
func (base *Base) AfterUpdate(tx *gorm.DB) (err error) {
	// update timestamps
	base.UpdatedAt = GenerateTimeStamps()
	return
}
