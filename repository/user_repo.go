package main

import (
	"final_project/db"
	"final_project/models"
)

func CreateUser(user models.User) (string, error){
	db, err := db.GetDB()
	if err != nil {
		return "Error", err
	}

	e := db.Table("users").Create(&user).Error
	if e != nil {
		return "Error", e
	}
	return "Create successfully", nil
}

func UpdateUser(email string, user models.User) (string, error){
	db, err := db.GetDB()
	if err != nil {
		return "Error", err
	}

	e := db.Table("users").Where("email = ?", email).Updates(user).Error
	if e != nil {
		return "Error", e
	}
	return "Update successfully", nil
}

func DeleteUser(email string) (string, error){
	db, err := db.GetDB()
	if err != nil {
		return "Error", err
	}
	
	e := db.Where("email = ?", email).Delete(&models.User{}).Error
	if e != nil {
		return "Error", e
	}
	return "Delete successfully", nil
}

func ReadUser(email string) (interface{}, error){
	db, err := db.GetDB()
	if err != nil {
		return "Error", err
	}
	
	var user models.User
	e := db.Where("email = ?", email).First(&user).Error
	if e != nil {
		return "Error", e
	}
	return user, nil
}

func GetListRoles(email string) ([]models.Role, error){
	db, err := db.GetDB()
	if err != nil {
		return nil, err
	}

	var user models.User
	e := db.Joins("JOIN user_role ON users.id = user_role.user_id").Preload("Roles").Find(&user).Error
	if e != nil {
		return nil, e
	}

	roles := user.Roles
	return roles, nil
}


