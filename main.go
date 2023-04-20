package main

import (
	"fmt"
	"go_rest/config"
	"go_rest/models"
	"go_rest/routes"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"

)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.GetConnection(config.BuildConfig()))
	if err != nil {
		fmt.Println("Status : ", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Nasabah{})

	r := routes.SetupRoutes()
	r.Run()
}
