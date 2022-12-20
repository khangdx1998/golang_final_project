package repository

import  (
	"final_project/db"
	"final_project/models"
	"github.com/google/uuid"
)

func CreateRole(role models.Role) error{
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	
	role.Id = uuid.New().String()
	e := db.Table("roles").Create(&role).Error
	if e != nil {
		return e
	}
	return nil
}

func UpdateRole(condition models.Condition, role models.Role) error{
	db, err := db.GetDB()
	if err != nil {
		return err
	}

	e := db.Table("roles").Where(condition.Field + " = ?", condition.Value).Updates(role).Error
	if e != nil {
		return e
	}
	return nil
}

func DeleteRole(condition models.Condition) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	
	e := db.Where(condition.Field + " = ?", condition.Value).Delete(&models.Role{}).Error
	if e != nil {
		return e
	}
	return nil
}

func ReadRole(condition models.Condition) (models.Role, error) {
	db, err := db.GetDB()
	if err != nil {
		return models.Role{}, err
	}
	
	var role models.Role
	e := db.Where(condition.Field + " = ?", condition.Value).First(&role).Error
	if e != nil {
		return models.Role{}, e
	}
	return role, nil
}