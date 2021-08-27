package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func SetupModels() *gorm.DB {
	viper.AutomaticEnv()
	viper_user := viper.Get("POSTGRES_USER")
	viper_password := viper.Get("POSTGRES_PASSWORD")
	viper_db := viper.Get("POSTGRES_DB")
	viper_host := viper.Get("POSTGRES_HOST")
	viper_port := viper.Get("POSTGRES_PORT")

	postgres_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		viper_host, viper_port, viper_user, viper_db, viper_password)

	fmt.Println("conname is\t\t", postgres_conname)
	db, err := gorm.Open("postgres", postgres_conname)
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&Book{})

	m := Book{Author: "Betsy Beyer", Title: "Site Reliability Engineering at google"}
	db.Create(&m)

	return db
}
