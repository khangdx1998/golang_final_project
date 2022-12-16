package models

type UserRole struct {
	Id string `json:"id" gorm:"column:id;primaryKey"`
	UserID string `json:"user_id" gorm:"column:user_id"`
	RoleID string `json:"role_id" gorm:"column:role_id"`
	Status string `json:"status" gorm:"column:status"`
}