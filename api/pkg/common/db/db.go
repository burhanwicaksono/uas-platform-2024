package db

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()
	DB_USER := viper.Get("USER_DB")
	DB_PASSWORD := viper.Get("PASSWORD_DB")
	DB_HOST := viper.Get("HOST_DB")
	DB_NAME := viper.Get("NAME_DB")

	db, err := sql.Open("mysql", "DB_USER:DB_PASSWORD@/DB_NAME")

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}