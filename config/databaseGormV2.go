package config

import (
	"fmt"
	mysql "gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
	logger "myapp/logger"
	"os"
)

//Generated By github.com/davidyap2002/GormCrudGenerator

//ConnectGorm Database Connection to Gorm V2
func ConnectGorm() *gorm.DB {
	databaseConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	db, err := gorm.Open(mysql.Open(databaseConfig), &gorm.Config{Logger: logger.InitLog()})

	if err != nil {
		fmt.Println(err)
		panic("Fail To Connect Database")
	}

	return db
}
