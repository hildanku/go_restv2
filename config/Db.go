package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type ConfigDB struct {
	Driver string
	Port   int
	User   string
	Pass   string
	DbName string
}

func BuildConfig() *ConfigDB {
	ConfigDB := ConfigDB{
		Driver: "localhost",
		Port:   3306,
		User:   "root",
		Pass:   "",
		DbName: "go_rest",
	}
	return &ConfigDB
}

func GetConnection(configDB *ConfigDB) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		configDB.User,
		configDB.Pass,
		configDB.Driver,
		configDB.Port,
		configDB.DbName,
	)
}
