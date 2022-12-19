package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"final_project/config"
	"fmt"
)


func GetDB() (*gorm.DB, error){
	env := config.Get_config()
	dsn := "host=" + env.Host + " " + "user=" + env.Username + " " + "password=" + env.Password + " " + "dbname=" + env.DB_name + " " + "port=" + env.Port + " sslmode=disable TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}