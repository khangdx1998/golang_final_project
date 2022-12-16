package models

import "time"

type User struct {
	Id string `json:"id" gorm:"column:id;primaryKey"`
	Email string `json:"email" gorm:"column:email;unique"`
	Name string `json:"name" gorm:"column:name"`
	Address string `json:"address" gorm:"column:address"`
	Password string `json:"password" gorm:"column:password"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at;default:null"`
	Roles []Role `gorm:"many2many:user_role;"`
}