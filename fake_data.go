package main

import (
	"final_project/models"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	cb "github.com/mxschmitt/golang-combinations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dsn = "host=127.0.0.1 user=khangdx password=abcd1234 dbname=final_project port=5432 sslmode=disable TimeZone=Asia/Shanghai"
)

func import_user_data(){
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	table := db.Table("users")

	count := 0
	tx := table.Begin()
	for count < 1000000 {
		email_fake := string(gofakeit.Number(97, 122)) + string(gofakeit.Number(97, 122)) + gofakeit.Email()
		user := models.User{
			Id: uuid.New().String(),
			Email: email_fake,
			Name: gofakeit.Name(),
			Address: gofakeit.Address().Address,
			Password: gofakeit.Password(true, true, true, true, false, 10),
		}

		if err := tx.Create(&user).Error; err == nil {
			count += 1
		} else {
			fmt.Print(err)
			tx.Rollback()
			return
		}
	}
	tx.Commit()
}

func import_role_data() {
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	table := db.Table("roles")

	for i:=0; i < 1000; i++ {
		var roles []models.Role
		for i:=0; i < 1000; i++ {
			roles = append(roles, models.Role{
				Id: uuid.New().String(),
				Permissions: get_random_permission(),
			})
		}
		table.Create((&roles))
	}
	
}

func get_random_permission() string {
	random_number := gofakeit.Number(1, 4)
	result := cb.Combinations([]string{"create", "read", "update", "delete"}, random_number)
	action_string := ""
	for i, action := range result[0]{
		if i == 0 {
			action_string = action 
		} else {
			action_string = action_string + "," + action
		}
	}

	return action_string
}

func main() {
	import_user_data()
	import_role_data()
}