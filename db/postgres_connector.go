package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func GetDB() (*gorm.DB, error){
	dsn := "host=127.0.0.1 user=khangdx password=abcd1234 dbname=final_project port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}