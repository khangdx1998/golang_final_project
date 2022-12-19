package repository

import (
	"final_project/db"
	"final_project/models"
	"github.com/google/uuid"
)

func CreateUser(user models.User) (string, error){
	db, err := db.GetDB()
	if err != nil {
		return "Error", err
	}
	
	user.Id = uuid.New().String()
	e := db.Table("users").Create(&user).Error
	if e != nil {
		return "Error", e
	}
	return "Create successfully", nil
}

func UpdateUser(condition models.Condition, user models.User) (string, error){
	db, err := db.GetDB()
	if err != nil {
		return "Error", err
	}

	e := db.Table("users").Where(condition.Field + " = ?", condition.Value).Updates(user).Error
	if e != nil {
		return "Error", e
	}
	return "Update successfully", nil
}

func DeleteUser(condition models.Condition) (string, error){
	db, err := db.GetDB()
	if err != nil {
		return "Error", err
	}
	
	e := db.Where(condition.Field + " = ?", condition.Value).Delete(&models.User{}).Error
	if e != nil {
		return "Error", e
	}
	return "Delete successfully", nil
}

func ReadUser(condition models.Condition) (models.User, error){
	db, err := db.GetDB()
	if err != nil {
		return models.User{}, err
	}
	
	var user models.User
	e := db.Where(condition.Field + " = ?", condition.Value).First(&user).Error
	if e != nil {
		return models.User{}, e
	}
	return user, nil
}

func GetListRoles(condition models.Condition) ([]models.Role, error){
	db, err := db.GetDB()
	if err != nil {
		return nil, err
	}

	var user models.User
	e := db.Joins("JOIN user_role ON users.id = user_role.user_id").Preload("Roles").Find(&user).Where(condition.Field + " = ?", condition.Value).Error
	if e != nil {
		return nil, e
	}

	roles := user.Roles
	return roles, nil
}


