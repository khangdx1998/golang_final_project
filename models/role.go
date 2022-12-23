package models

type Role struct {
	BaseModel
	Permissions string `json:"permissions" gorm:"column:permissions"`
	Users []User `gorm:"many2many:user_role;"`
}