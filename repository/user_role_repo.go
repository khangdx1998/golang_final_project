package repository

import (
	"final_project/db"
	"final_project/models"
	"github.com/google/uuid"
)


func CreateUserRole(user_role models.UserRole) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}

	user_role.Id = uuid.New().String()
	user_role.Status = "true"
	e := db.Table("user_role").Create(&user_role).Error
	if e != nil {
		return err
	}
	return nil
}