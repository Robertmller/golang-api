package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var (
	db *gorm.DB
)

func Connect() {
	user_db := goDotEnvVariable("USER_DB")
	password_db := goDotEnvVariable("PASSWORD_DB")
	host_db := goDotEnvVariable("HOST_DB")
	port_db := goDotEnvVariable("PORT_DB")
	table_db := goDotEnvVariable("TABLE_DB")

	d, err := gorm.Open("mysql", user_db+":"+password_db+"@tcp("+host_db+":"+port_db+")/"+table_db+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db

}
