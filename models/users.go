package models


type User struct {
	BaseModel
	Email string `json:"email" gorm:"column:email;unique"`
	Name string `json:"name" gorm:"column:name"`
	Address string `json:"address" gorm:"column:address"`
	Password string `json:"password" gorm:"-"`
	Roles []Role `json:"roles" gorm:"many2many:user_role;"`
}