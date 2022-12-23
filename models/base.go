package models

import "time"


type BaseModel struct {
	Id string `json:"id" gorm:"column:id;primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at;default:null"`
}